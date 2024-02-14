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
