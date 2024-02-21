package swipe_id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type SwipeID uuid.UUID

// Wraps a UUID in an SwipeID to prevent misuse
func Wrap(uuid uuid.UUID) SwipeID {
	return SwipeID(uuid)
}

// Extracts the base UUID from an SwipeID
func (id SwipeID) Unwrap() uuid.UUID {
	return uuid.UUID(id)
}

func (id *SwipeID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = Wrap(uuid)
	return err
}

func (id SwipeID) Value() (driver.Value, error) {
	return id.Unwrap().Value()
}

func (id SwipeID) MarshalText() ([]byte, error) {
	return id.Unwrap().MarshalText()
}

func (id *SwipeID) UnmarshalText(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalText(data)
	*id = Wrap(uuid)
	return err
}

func (id SwipeID) MarshalBinary() ([]byte, error) {
	return id.Unwrap().MarshalBinary()
}

func (id *SwipeID) UnmarshalBinary(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalBinary(data)
	*id = Wrap(uuid)
	return err
}

func (id SwipeID) String() string {
	return id.Unwrap().String()
}
