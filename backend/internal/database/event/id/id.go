package id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type EventID uuid.UUID

func New(uuid uuid.UUID) EventID {
	return EventID(uuid)
}

func (id EventID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id *EventID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = New(uuid)
	return err
}

func (id EventID) Value() (driver.Value, error) {
	return id.UUID().Value()
}
