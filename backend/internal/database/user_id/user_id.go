package user_id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type UserID uuid.UUID

// Wraps a UUID in an UserID to prevent misuse
func Wrap(uuid uuid.UUID) UserID {
	return UserID(uuid)
}

// Extracts the base UUID from an UserID
func (id UserID) Unwrap() uuid.UUID {
	return uuid.UUID(id)
}

func (id *UserID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = Wrap(uuid)
	return err
}

func (id UserID) Value() (driver.Value, error) {
	return id.Unwrap().Value()
}

func (id UserID) MarshalText() ([]byte, error) {
	return id.Unwrap().MarshalText()
}

func (id *UserID) UnmarshalText(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalText(data)
	*id = Wrap(uuid)
	return err
}

func (id UserID) MarshalBinary() ([]byte, error) {
	return id.Unwrap().MarshalBinary()
}

func (id *UserID) UnmarshalBinary(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalBinary(data)
	*id = Wrap(uuid)
	return err
}

func (id UserID) String() string {
	return id.Unwrap().String()
}
