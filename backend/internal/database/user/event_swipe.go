package user

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user_id"
	"time"

	"gorm.io/gorm"
)

type EventSwipe struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time        `validate:"gtefield=CreatedAt"`
	UserID    user_id.UserID   `gorm:"index:pair,unique" validate:"required"`
	EventID   event_id.EventID `gorm:"index:pair,unique" validate:"required"`
	Liked     bool
}

// Automatically rolls back transactions that save invalid data to the database
func (es *EventSwipe) BeforeSave(tx *gorm.DB) error {
	return es.Validate()
}

// Ensures the user and its fields are valid
func (es EventSwipe) Validate() error {
	return validate.Struct(es)
}
