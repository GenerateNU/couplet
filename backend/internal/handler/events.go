package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"errors"
	"fmt"
)

// Creates a new event.
// POST /events
func (h Handler) EventsPost(ctx context.Context, req *api.EventsPostReq) (api.EventsPostRes, error) {
	if h.logger != nil {
		h.logger.Info("POST /events")
	}

	eventTags := []event.EventTag{}
	for _, v := range req.Tags {
		eventTags = append(eventTags, event.EventTag{ID: v})
	}

	e, valErr, txErr := h.controller.CreateEvent(event.Event{
		Name:      req.Name,
		Bio:       req.Bio,
		Images:    url_slice.Wrap(req.Images),
		EventTags: eventTags,
		OrgID:     org_id.Wrap(req.OrgId),
	})
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: "failed to validate event",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to create event")
	}

	tags := []string{}
	for _, eventTag := range e.EventTags {
		tags = append(tags, eventTag.ID)
	}
	res := api.EventsPostCreated{
		ID:     e.ID.Unwrap(),
		Name:   e.Name,
		Bio:    e.Bio,
		Images: e.Images.Unwrap(),
		Tags:   tags,
		OrgId:  e.OrgID.Unwrap(),
	}
	return &res, nil
}

// Deletes an event by its ID.
// DELETE /events/{id}
func (h Handler) EventsIDDelete(ctx context.Context, params api.EventsIDDeleteParams) (api.EventsIDDeleteRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("DELETE /events/%s", params.ID))
	}

	e, txErr := h.controller.DeleteEvent(event_id.Wrap(params.ID))
	if txErr != nil {
		return &api.Error{
			Code:    404,
			Message: "event not found",
		}, nil
	}

	tags := []string{}
	for _, eventTag := range e.EventTags {
		tags = append(tags, eventTag.ID)
	}
	res := api.EventsIDDeleteOK{
		ID:     e.ID.Unwrap(),
		Name:   e.Name,
		Bio:    e.Bio,
		Images: e.Images.Unwrap(),
		Tags:   tags,
		OrgId:  e.OrgID.Unwrap(),
	}
	return &res, nil
}

// RecommendationEventsGet implements api.Handler.
func (h Handler) RecommendationEventsGet(ctx context.Context, params api.RecommendationEventsGetParams) ([]api.RecommendationEventsGetOKItem, error) {
	events, err := h.controller.GetRandomEvents(params)
	if err != nil {
		return nil, err
	}
	var res []api.RecommendationEventsGetOKItem
	for _, e := range events {

		// Get Event Tags
		eventTags := []string{}
		if e.EventTags != nil {
			for i := range e.EventTags {
				eventTags = append(eventTags, e.EventTags[i].ID)
			}
		}

		res = append(res, api.RecommendationEventsGetOKItem{
			ID:     e.ID.Unwrap(),
			Name:   e.Name,
			Bio:    e.Bio,
			Images: e.Images,
			Tags:   eventTags,
			OrgId:  api.NewOptUUID(e.OrgID.Unwrap()),
		})
	}
	return res, nil
}

// Gets an event by its ID.
// GET /events/{id}
func (h Handler) EventsIDGet(ctx context.Context, params api.EventsIDGetParams) (api.EventsIDGetRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("GET /events/%s", params.ID))
	}

	e, txErr := h.controller.GetEvent(event_id.Wrap(params.ID))
	if txErr != nil {
		return &api.Error{
			Code:    404,
			Message: "event not found",
		}, nil
	}

	tags := []string{}
	for _, eventTag := range e.EventTags {
		tags = append(tags, eventTag.ID)
	}
	res := api.EventsIDGetOK{
		ID:     e.ID.Unwrap(),
		Name:   e.Name,
		Bio:    e.Bio,
		Images: e.Images.Unwrap(),
		Tags:   tags,
		OrgId:  e.OrgID.Unwrap(),
	}
	return &res, nil
}

// Gets all events with pagination.
// GET /events
func (h Handler) EventsGet(ctx context.Context, params api.EventsGetParams) ([]api.EventsGetOKItem, error) {
	if h.logger != nil {
		h.logger.Info("GET /events")
	}

	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	events, txErr := h.controller.GetEvents(limit, offset)
	if txErr != nil {
		return nil, errors.New("failed to get events")
	}
	res := []api.EventsGetOKItem{}
	for _, e := range events {
		tags := []string{}
		for _, eventTag := range e.EventTags {
			tags = append(tags, eventTag.ID)
		}
		item := api.EventsGetOKItem{
			ID:     e.ID.Unwrap(),
			Name:   e.Name,
			Bio:    e.Bio,
			Images: e.Images.Unwrap(),
			Tags:   tags,
			OrgId:  e.OrgID.Unwrap(),
		}
		res = append(res, item)
	}
	return res, nil
}

// Partially updates an organization by its ID.
// PATCH /events/{id}
func (h Handler) EventsIDPatch(ctx context.Context, req *api.EventsIDPatchReq, params api.EventsIDPatchParams) (api.EventsIDPatchRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PATCH /events/%s", params.ID))
	}

	var reqEvent event.Event
	reqEvent.ID = event_id.Wrap(params.ID)
	if req.Name.Set {
		reqEvent.Name = req.Name.Value
	}
	if req.Bio.Set {
		reqEvent.Bio = req.Bio.Value
	}
	reqEvent.Images = url_slice.Wrap(req.Images)
	reqEvent.EventTags = []event.EventTag{}
	for _, v := range req.Tags {
		reqEvent.EventTags = append(reqEvent.EventTags, event.EventTag{ID: v})
	}

	e, valErr, txErr := h.controller.UpdateEvent(reqEvent)
	if valErr != nil {
		fmt.Println(valErr)
		return &api.EventsIDPatchBadRequest{
			Code:    400,
			Message: "failed to validate event",
		}, nil
	}
	if txErr != nil {
		return &api.EventsIDPatchNotFound{
			Code:    404,
			Message: "event not found",
		}, nil
	}

	tags := []string{}
	for _, eventTag := range e.EventTags {
		tags = append(tags, eventTag.ID)
	}
	res := api.EventsIDPatchOK{
		ID:     e.ID.Unwrap(),
		Name:   e.Name,
		Bio:    e.Bio,
		Images: e.Images.Unwrap(),
		Tags:   tags,
		OrgId:  e.OrgID.Unwrap(),
	}
	return &res, nil
}

// Creates a new event or updates an existing event.
// PUT /events/{id}
func (h Handler) EventsIDPut(ctx context.Context, req *api.EventsIDPutReq, params api.EventsIDPutParams) (api.EventsIDPutRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PUT /events/%s", params.ID))
	}

	eventTags := []event.EventTag{}
	for _, v := range req.Tags {
		eventTags = append(eventTags, event.EventTag{ID: v})
	}

	e, valErr, txErr := h.controller.SaveEvent(event.Event{
		Name:      req.Name,
		Bio:       req.Bio,
		Images:    url_slice.Wrap(req.Images),
		EventTags: eventTags,
		OrgID:     org_id.Wrap(req.OrgId),
	})
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: "failed to validate event",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to save event")
	}

	tags := []string{}
	for _, eventTag := range e.EventTags {
		tags = append(tags, eventTag.ID)
	}
	res := api.EventsIDPutOK{
		ID:     e.ID.Unwrap(),
		Name:   e.Name,
		Bio:    e.Bio,
		Images: e.Images.Unwrap(),
		Tags:   tags,
		OrgId:  e.OrgID.Unwrap(),
	}
	return &res, nil
}
