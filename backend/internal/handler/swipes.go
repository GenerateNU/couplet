package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/event_id"
	"couplet/internal/database/swipe"
	"couplet/internal/database/user_id"
	"errors"
)

// CreateEventSwipe implements api.Handler.
func (h Handler) CreateEventSwipe(ctx context.Context, req *api.EventSwipe) (api.CreateEventSwipeRes, error) {
	var eventSwipeToCreate swipe.EventSwipe
	eventSwipeToCreate.UserId = user_id.UserID(req.UserId)
	eventSwipeToCreate.EventId = event_id.EventID(req.EventId)
	eventSwipeToCreate.Liked = req.Liked

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

func (h Handler) CreateUserSwipe(ctx context.Context, req *api.UserSwipe) (api.CreateUserSwipeRes, error) {
	var userSwipeToCreate swipe.UserSwipe
	userSwipeToCreate.UserId = user_id.UserID(req.UserId)
	userSwipeToCreate.UserSwipeId = user_id.UserID(req.OtherUserId)
	userSwipeToCreate.Liked = req.Liked

	us, err := h.controller.CreateUserSwipe(userSwipeToCreate)
	if err != nil {
		return nil, errors.New("failed to create user swipe")
	}

	res := api.UserSwipe{
		UserId:      us.UserId.Unwrap(),
		OtherUserId: us.UserSwipeId.Unwrap(),
		Liked:       us.Liked,
	}

	return &res, nil
}
