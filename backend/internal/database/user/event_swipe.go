package user

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/user_id"
	"time"

	"gorm.io/gorm"
)

type EventSwipe struct {
	UserID    user_id.UserID   `gorm:"primaryKey" validate:"required"`
	EventID   event_id.EventID `gorm:"primaryKey" validate:"required"`
	Event     event.Event      `gorm:"foreignKey:EventID"`
	Liked     bool
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
}

// Automatically rolls back transactions that save invalid data to the database
func (e *EventSwipe) AfterSave(tx *gorm.DB) error {
	return e.Validate()
}

// Ensures the user and its fields are valid
func (e EventSwipe) Validate() error {
	// TODO: Write tests
	return validate.Struct(e)
}
