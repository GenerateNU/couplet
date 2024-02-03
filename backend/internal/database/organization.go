package database

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID        OrganizationID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Events    []Event
	Bio       string
}

// Automatically generates a random UUID before creating an organization
func (o *Organization) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = OrganizationID(uuid.New())
	return
}

type OrganizationID uuid.UUID

func (id *OrganizationID) Scan(src interface{}) error {
	baseId := uuid.UUID(*id)
	return (&baseId).Scan(src)
}

func (id OrganizationID) Value() (driver.Value, error) {
	baseId := uuid.UUID(id)
	return baseId.Value()
}
