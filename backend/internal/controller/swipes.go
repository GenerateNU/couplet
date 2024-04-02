package controller

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/event_swipe"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"couplet/internal/database/user_swipe"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Creates a new event swipe in the database
func (c Controller) CreateEventSwipe(params event_swipe.EventSwipe) (es event_swipe.EventSwipe, valErr error, txErr error) {
	es = params
	var timestampErr error
	if es.UpdatedAt.Before(es.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var userIdErr error
	if (es.UserID == user_id.UserID{}) {
		userIdErr = fmt.Errorf("invalid user ID")
	}
	var eventIdErr error
	if (es.EventID == event_id.EventID{}) {
		eventIdErr = fmt.Errorf("invalid event ID")
	}
	valErr = errors.Join(timestampErr, userIdErr, eventIdErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Omit(clause.Associations).Create(&es).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Creates a new user swipe in the database and checks for a match
func (c Controller) CreateUserSwipe(params user_swipe.UserSwipe) (us user_swipe.UserSwipe, valErr error, txErr error) {
	us = params
	var timestampErr error
	if us.UpdatedAt.Before(us.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var userIdErr error
	if (us.UserID == user_id.UserID{}) {
		userIdErr = fmt.Errorf("invalid user ID for swipe sender")
	}
	var otherUserIdErr error
	if (us.OtherUserID == user_id.UserID{}) {
		otherUserIdErr = fmt.Errorf("invalid user ID for swipe receiver")
	}
	valErr = errors.Join(timestampErr, userIdErr, otherUserIdErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Omit(clause.Associations).Create(&us).Error
	if txErr != nil {
		tx.Rollback()
		return
	}

	// Check for a match only if the current swipe is a 'like'
	if us.Liked {
		var otherSwipe user_swipe.UserSwipe
		// This query checks for a reciprocal like.
		err := tx.Where("user_id = ? AND other_user_id = ? AND liked = ?", us.OtherUserID, us.UserID, true).First(&otherSwipe).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// Commit the creation of the user swipe, but don't continue
				tx.Commit()
				return
			}
			tx.Rollback()
			return
		}

		// Logic to handle a found reciprocal swipe, e.g., creating a match.
		// Grabs both users associated with the swipe, and inserts them into each other's matches list
		var userOne user.User
		c.database.First(&userOne, otherSwipe.UserID)

		var userTwo user.User
		c.database.First(&userTwo, otherSwipe.OtherUserID)

		txErr = c.database.Model(&userOne).Association("Matches").Append(&userTwo)
		if txErr != nil {
			tx.Rollback()
			return
		}

		txErr = c.database.Model(&userTwo).Association("Matches").Append(&userOne)
		if txErr != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return
}
