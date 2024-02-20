package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"errors"
	"fmt"
)

// Creates a new event.
// POST /events
func (h Handler) CreateEvent(ctx context.Context, req *api.CreateEventReq) (api.CreateEventRes, error) {
	// TODO: Write tests
	h.logger.Info("POST /events")

	var eventToCreate event.Event
	eventToCreate.Name = req.Name
	eventToCreate.Bio = req.Bio

	e, err := h.controller.CreateEvent(eventToCreate)
	// TODO: check for validation error from the controller and return 400
	if err != nil {
		return nil, errors.New("failed to create event")
	}

	res := api.CreateEventCreated{
		ID:   e.ID.Unwrap(),
		Name: e.Name,
		Bio:  e.Bio,
	}

	return &res, nil
}

// Deletes an event by its ID.
// DELETE /events/{id}
func (h Handler) DeleteEvent(ctx context.Context, params api.DeleteEventParams) (api.DeleteEventRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("DELETE /events/%s", params.ID))
	o, err := h.controller.DeleteEvent(event_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.DeleteEventOK{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
	}
	return &res, nil
}

// GET (/events/:id) a single event by their id
func (h Handler) GetEvent(ctx context.Context, params api.GetEventParams) (api.GetEventRes, error) {
	h.logger.Info(fmt.Sprintf("GET /events/%s", params.ID))
	e, err := h.controller.GetEvent(event_id.Wrap(params.ID))
	if err != nil {
		return nil, errors.New("failed to get event")
	}

	res := api.GetEventOK{
		ID:   e.ID.Unwrap(),
		Name: e.Name,
		Bio:  e.Bio,
	}

	return &res, nil
}

// GET (/events) all events with pagination
func (h Handler) GetEvents(ctx context.Context, params api.GetEventsParams) ([]api.GetEventsOKItem, error) {
	h.logger.Info("GET /events")

	// this syntax is not correct

	// events, err := h.controller.GetEvents(params.Limit, params.Offset)
	// if err != nil {
	// 	return nil, errors.New("failed to get events")
	// }

	// var res api.GetEventsOK
	// for _, e := range events {
	// 	res = append(res, &api.Event{
	// 		ID:   e.ID.Unwrap(),
	// 		Name: e.Name,
	// 		Bio:  e.Bio,
	// 	})
	// }

	return nil, nil
}

// PUT (/events/:id) to completely update an existing event, returning the created object if successful
func (h Handler) PutEvent(ctx context.Context, updatedEvent *api.PutEventReq, params api.PutEventParams) (api.PutEventRes, error) {
	h.logger.Info(fmt.Sprintf("PUT /events/%s", params.ID))
	e, err := h.controller.PutEvent(event_id.Wrap(params.ID), event.Event{
		Name: updatedEvent.Name,
		Bio:  updatedEvent.Bio,
	})
	if err != nil {
		return nil, errors.New("failed to update event")
	}

	res := api.PutEventOK{
		ID:   e.ID.Unwrap(),
		Name: e.Name,
		Bio:  e.Bio,
	}

	return &res, nil
}

// PATCH (/events/:id) to partially update one or many fields of an existing event, returning the created object if successful

func (h Handler) PatchEvent(ctx context.Context, req *api.Event, params api.PatchEventParams) (api.PatchEventRes, error) {
	h.logger.Info(fmt.Sprintf("PATCH /events/%s", params.ID))
	// todo
	return nil, nil
}
