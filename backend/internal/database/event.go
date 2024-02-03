package database

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID             EventID `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	Bio            string
	OrganizationID OrganizationID
}

// Automatically generates a random UUID before creating an event
func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = EventID(uuid.New())
	return
}

type EventID uuid.UUID

func (id *EventID) Scan(src interface{}) error {
	baseId := uuid.UUID(*id)
	return (&baseId).Scan(src)
}

func (id EventID) Value() (driver.Value, error) {
	baseId := uuid.UUID(id)
	return baseId.Value()
}
