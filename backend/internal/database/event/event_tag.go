package event

import (
	"time"

	"gorm.io/gorm"
)

type EventTag struct {
	ID        string `gorm:"primaryKey" validate:"required,min=1,max=255"`
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
	Events    []Event   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:events2tags"`
}

// Automatically rolls back transactions that save invalid event tags to the database
func (t *EventTag) BeforeSave(tx *gorm.DB) error {
	return t.Validate()
}

func (t EventTag) Validate() error {
	return validate.Struct(t)
}
