package event

import (
	"couplet/internal/database/event_id"
	"time"

	"gorm.io/gorm"
)

type EventImage struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
	Url       string    `validate:"required,url"`
	EventID   event_id.EventID
}

// Automatically rolls back transactions that save invalid data to the database
func (i *EventImage) AfterSave(tx *gorm.DB) error {
	return i.Validate()
}

// Ensures the image and its fields are valid
func (i EventImage) Validate() error {
	return validate.Struct(i)
}
