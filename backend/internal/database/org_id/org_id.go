package org_id

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID wrapper to prevent confusion among other UUIDs
type OrgID uuid.UUID

// Wraps a UUID in an OrgID to prevent misuse
func Wrap(uuid uuid.UUID) OrgID {
	return OrgID(uuid)
}

// Extracts the base UUID from an OrgID
func (id OrgID) Unwrap() uuid.UUID {
	return uuid.UUID(id)
}

func (id *OrgID) Scan(src interface{}) error {
	var uuid uuid.UUID
	err := uuid.Scan(src)
	*id = Wrap(uuid)
	return err
}

func (id OrgID) Value() (driver.Value, error) {
	return id.Unwrap().Value()
}

func (id OrgID) MarshalText() ([]byte, error) {
	return id.Unwrap().MarshalText()
}

func (id *OrgID) UnmarshalText(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalText(data)
	*id = Wrap(uuid)
	return err
}

func (id OrgID) MarshalBinary() ([]byte, error) {
	return id.Unwrap().MarshalBinary()
}

func (id *OrgID) UnmarshalBinary(data []byte) error {
	var uuid uuid.UUID
	err := uuid.UnmarshalBinary(data)
	*id = Wrap(uuid)
	return err
}

func (id OrgID) String() string {
	return id.Unwrap().String()
}
