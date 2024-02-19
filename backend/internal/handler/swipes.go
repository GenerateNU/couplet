package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/event_id"
	"couplet/internal/database/swipe"
	"couplet/internal/database/user_id"
	"errors"
	"time"
)

// CreateEventSwipe implements api.Handler.
func (h Handler) CreateEventSwipe(ctx context.Context, req *api.EventSwipe) (api.CreateEventSwipeRes, error) {
	var eventSwipeToCreate swipe.EventSwipe
	eventSwipeToCreate.UserId = user_id.UserID(req.UserId)
	eventSwipeToCreate.EventId = event_id.EventID(req.EventId)
	eventSwipeToCreate.Liked = req.Liked
	eventSwipeToCreate.CreatedAt = time.Now()
	eventSwipeToCreate.UpdatedAt = time.Now()

	es, err := h.controller.CreateEventSwipe(eventSwipeToCreate)
	if err != nil {
		return nil, errors.New("failed to create event swipe")
	}

	res := api.CreateEventSwipeCreated{
		ID:      es.ID.Unwrap(),
		UserId:  es.UserId.Unwrap(),
		EventId: es.EventId.Unwrap(),
		Liked:   es.Liked,
	}

	return &res, nil
}

// CreateUserSwipe implements api.Handler.
func (Handler) CreateUserSwipe(ctx context.Context, req *api.UserSwipe) (api.CreateUserSwipeRes, error) {
	panic("unimplemented")
}
