package event_id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type EventID uuid.UUID

// Wraps a UUID in an EventID to prevent misuse
func Wrap(uuid uuid.UUID) EventID {
	return EventID(uuid)
}

// Extracts the base UUID from an EventID
func (id EventID) Unwrap() uuid.UUID {
	return uuid.UUID(id)
}

func (id *EventID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = Wrap(uuid)
	return err
}

func (id EventID) Value() (driver.Value, error) {
	return id.Unwrap().Value()
}

func (id EventID) MarshalText() ([]byte, error) {
	return id.Unwrap().MarshalText()
}

func (id *EventID) UnmarshalText(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalText(data)
	*id = Wrap(uuid)
	return err
}

func (id EventID) MarshalBinary() ([]byte, error) {
	return id.Unwrap().MarshalBinary()
}

func (id *EventID) UnmarshalBinary(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalBinary(data)
	*id = Wrap(uuid)
	return err
}

func (id EventID) String() string {
	return id.Unwrap().String()
}
