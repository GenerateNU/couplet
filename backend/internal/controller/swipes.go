package controller

import (
	"couplet/internal/database/swipe"
	"fmt"
)

// Creates a new event swipe in the database
func (c Controller) CreateEventSwipe(params swipe.EventSwipe) (e *swipe.EventSwipe, err error) {
	newEventSwipe := swipe.EventSwipe{
		UserId:  params.UserId,
		EventId: params.EventId,
		Liked:   params.Liked,
	}

	res := c.database.Create(&newEventSwipe)
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("no event swipe created")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return &newEventSwipe, nil
}

// creates a new user swipe in the database
func (c Controller) CreateUserSwipe(params swipe.UserSwipe) (e *swipe.UserSwipe, err error) {
	newUserSwipe := swipe.UserSwipe{
		UserId:      params.UserId,
		UserSwipeId: params.UserSwipeId,
		Liked:       params.Liked,
	}

	res := c.database.Create(&newUserSwipe)
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("no user swipe created")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return &newUserSwipe, nil
}
