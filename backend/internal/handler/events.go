package handler

import (
	"context"
	"couplet/internal/api"
)

func (h Handler) CreateEvent(ctx context.Context, event *api.Event) (api.CreateEventRes, error) {
	event, err := h.controller.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (h Handler) DeleteEventById(ctx context.Context, params api.DeleteEventByIdParams) (api.DeleteEventByIdRes, error) {
	deletedEvent, err := h.controller.DeleteEventById(ctx, params.EventId)
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}
	return deletedEvent, nil
}
