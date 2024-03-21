package user

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user_id"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type EventSwipe struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    user_id.UserID   `gorm:"index:pair,unique"`
	EventID   event_id.EventID `gorm:"index:pair,unique"`
	Liked     bool
}

// Automatically rolls back transactions that save invalid data to the database
func (es *EventSwipe) BeforeSave(tx *gorm.DB) error {
	return es.Validate()
}

// Ensures the event swipe and its fields are valid
func (es EventSwipe) Validate() error {
	var timestampErr error
	if es.UpdatedAt.Before(es.CreatedAt) {
		return fmt.Errorf("invalid timestamps")
	}
	var userIdErr error
	if (es.UserID == user_id.UserID{}) {
		userIdErr = fmt.Errorf("invalid user ID")
	}
	var eventIdErr error
	if (es.EventID == event_id.EventID{}) {
		eventIdErr = fmt.Errorf("invalid event ID")
	}
	return errors.Join(timestampErr, userIdErr, eventIdErr)
}
