package controller

import (
	"couplet/internal/database/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Creates a new event swipe in the database
func (c Controller) CreateEventSwipe(params user.EventSwipe) (es user.EventSwipe, valErr error, txErr error) {
	// TODO: Write tests
	es = params
	valErr = es.Validate()
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Omit(clause.Associations).Create(&es).Error

	if txErr != nil {
		tx.Rollback()
	}

	tx.Commit()
	return
}

// Creates a new user swipe in the database and checks for a match
func (c Controller) CreateUserSwipe(params user.UserSwipe) (us user.UserSwipe, valErr error, txErr error) {
	// Validate the incoming swipe
	us = params
	valErr = us.Validate()
	if valErr != nil {
		return
	}

	// Start a transaction
	tx := c.database.Begin()

	// Save the new swipe
	txErr = tx.Omit(clause.Associations).Create(&us).Error
	if txErr != nil {
		tx.Rollback()
		return
	}

	// Check for a match only if the current swipe is a 'like'
	if us.Liked {
		var otherSwipe user.UserSwipe
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
		var userOne *user.User
		c.database.First(&userOne, otherSwipe.UserID)

		var userTwo *user.User
		c.database.First(&userTwo, otherSwipe.OtherUserID)

		txErr = c.database.Model(&userOne).Association("Matches").Append(userTwo)
		if txErr != nil {
			tx.Rollback()
			return
		}

		txErr = c.database.Model(&userTwo).Association("Matches").Append(userOne)
		if txErr != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return
}
