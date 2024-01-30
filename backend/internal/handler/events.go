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
	err := h.controller.DeleteEventById(ctx, params.EventId)
	if err != nil {
		return nil, err
	}
	return &api.DeleteEventByIdNoContent{}, nil
}

