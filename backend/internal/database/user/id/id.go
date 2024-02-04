package id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type UserID uuid.UUID

func New(uuid uuid.UUID) UserID {
	return UserID(uuid)
}

func (id UserID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id *UserID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = New(uuid)
	return err
}

func (id UserID) Value() (driver.Value, error) {
	return id.UUID().Value()
}

func (id UserID) String() string {
	return id.UUID().String()
}
