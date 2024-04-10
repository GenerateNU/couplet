package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"errors"
	"fmt"

	ht "github.com/ogen-go/ogen/http"
)

// Creates a new event.
// POST /events
func (h Handler) EventsPost(ctx context.Context, req *api.EventsPostReq) (api.EventsPostRes, error) {
	// TODO: Write tests
	h.logger.Info("POST /events")

	var eventToCreate event.Event
	eventToCreate.Name = req.Name
	eventToCreate.Bio = req.Bio
	eventToCreate.Images = []event.EventImage{}
	for _, v := range req.Images {
		eventToCreate.Images = append(eventToCreate.Images, event.EventImage{Url: v.String()})
	}
	eventToCreate.OrgID = org_id.Wrap(req.OrgId)

	e, err := h.controller.CreateEvent(eventToCreate)
	// TODO: check for validation error from the controller and return 400
	if err != nil {
		return nil, errors.New("failed to create event")
	}

	res := api.EventsPostCreated{
		ID:     e.ID.Unwrap(),
		Name:   e.Name,
		Bio:    e.Bio,
		Images: req.Images,
		OrgId:  api.NewOptUUID(e.OrgID.Unwrap()),
	}

	return &res, nil
}

// Deletes an event by its ID.
// DELETE /events/{id}
func (h Handler) EventsIDDelete(ctx context.Context, params api.EventsIDDeleteParams) (api.EventsIDDeleteRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("DELETE /events/%s", params.ID))
	o, err := h.controller.DeleteEvent(event_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.EventsIDDeleteOK{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
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
		res = append(res, api.RecommendationEventsGetOKItem{
			ID:   e.ID.Unwrap(),
			Name: e.Name,
			Bio:  e.Bio,
		})
	}
	return res, nil
}

// Gets an event by its ID.
// GET /events/{id}
func (h Handler) EventsIDGet(ctx context.Context, params api.EventsIDGetParams) (api.EventsIDGetRes, error) {
	h.logger.Info(fmt.Sprintf("GET /events/%s", params.ID))
	e, err := h.controller.GetEvent(event_id.Wrap(params.ID))
	if err != nil {
		return nil, errors.New("failed to get event")
	}

	res := api.EventsIDGetOK{
		ID:   e.ID.Unwrap(),
		Name: e.Name,
		Bio:  e.Bio,
	}

	return &res, nil
}

// Gets all events with pagination.
// GET /events
func (h Handler) EventsGet(ctx context.Context, params api.EventsGetParams) ([]api.EventsGetOKItem, error) {
	h.logger.Info("GET /events")

	events, err := h.controller.GetEvents(params.Limit, params.Offset)
	if err != nil {
		return nil, errors.New("failed to get events")
	}

	var res []api.EventsGetOKItem
	for _, e := range events {
		res = append(res, api.EventsGetOKItem{
			ID:   e.ID.Unwrap(),
			Name: e.Name,
			Bio:  e.Bio,
		})
	}

	return res, nil
}

// Completely updates an existing event
// PUT /events/{id}
func (h Handler) EventsIDPut(ctx context.Context, updatedEvent *api.EventsIDPutReq, params api.EventsIDPutParams) (api.EventsIDPutRes, error) {
	h.logger.Info(fmt.Sprintf("PUT /events/%s", params.ID))
	e, err := h.controller.PutEvent(event_id.Wrap(params.ID), event.Event{
		Name: updatedEvent.Name,
		Bio:  updatedEvent.Bio,
	})
	if err != nil {
		return nil, errors.New("failed to update event")
	}

	res := api.EventsIDPutOK{
		ID:   e.ID.Unwrap(),
		Name: e.Name,
		Bio:  e.Bio,
	}

	return &res, nil
}

// Partially update one or many fields of an existing event
// PATCH /events/{id}
func (h Handler) EventsIDPatch(ctx context.Context, req *api.Event, params api.EventsIDPatchParams) (api.EventsIDPatchRes, error) {
	h.logger.Info(fmt.Sprintf("PATCH /events/%s", params.ID))
	return nil, ht.ErrNotImplemented
}
