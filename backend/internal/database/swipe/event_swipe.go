package swipe

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/swipe_id"
	"couplet/internal/database/user_id"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type EventSwipe struct {
	ID        swipe_id.SwipeID `gorm:"primaryKey" validate:"required"`
	UserId    user_id.UserID   `gorm:"primaryKey"`
	EventId   event_id.EventID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Liked     bool
}

// Automatically generates a random ID if unset before creating
func (e *EventSwipe) BeforeCreate(tx *gorm.DB) error {
	if (e.ID == swipe_id.SwipeID{}) {
		e.ID = swipe_id.Wrap(uuid.New())
	}
	return nil
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
