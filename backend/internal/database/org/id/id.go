package id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type OrgID uuid.UUID

func New(uuid uuid.UUID) OrgID {
	return OrgID(uuid)
}

func (id OrgID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id *OrgID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = New(uuid)
	return err
}

func (id OrgID) Value() (driver.Value, error) {
	return id.UUID().Value()
}
