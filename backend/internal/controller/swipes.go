package controller

import (
	"couplet/internal/database/user"

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

// Creates a new user swipe in the database
func (c Controller) CreateUserSwipe(params user.UserSwipe) (us user.UserSwipe, valErr error, txErr error) {
	// TODO: Write tests
	us = params
	valErr = us.Validate()
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Omit(clause.Associations).Create(&us).Error

	if txErr != nil {
		tx.Rollback()
	}

	tx.Commit()
	return
}
