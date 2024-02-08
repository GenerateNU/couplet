package handler

import (
	"context"
	"couplet/internal/api"
)

func (h Handler) GetEventById(ctx context.Context, params api.GetEventByIdParams) (api.GetEventByIdRes, error) {
	event, err := h.controller.GetEventById(ctx, params)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (h Handler) GetAllEvents(ctx context.Context, params api.GetAllEventsParams) (api.GetAllEventsRes, error) {
	events, err := h.controller.GetAllEvents(ctx, params)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (h Handler) UpdateEventById(ctx context.Context, req *api.Event, params api.UpdateEventByIdParams) (api.UpdateEventByIdRes, error) {
	event, err := h.controller.UpdateEventById(ctx, req, params)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (h Handler) PatchEventById(ctx context.Context, req *api.Event, params api.PatchEventByIdParams) (api.PatchEventByIdRes, error) {
	event, err := h.controller.PatchEventById(ctx, req, params)
	if err != nil {
		return nil, err
	}

	return event, nil
}
