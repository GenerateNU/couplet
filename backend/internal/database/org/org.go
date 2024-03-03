package org

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org_id"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Org struct {
	ID        org_id.OrgID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
	Name      string    `validate:"required,min=1,max=255"`
	Bio       string    `validate:"required,min=1,max=255"`
	Image     OrgImage
	OrgTags   []OrgTag      `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:orgs2tags" validate:"max=5"`
	Events    []event.Event `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

// Automatically generates a random ID if unset before creating
func (o *Org) BeforeCreate(tx *gorm.DB) error {
	if (o.ID == org_id.OrgID{}) {
		o.ID = org_id.Wrap(uuid.New())
	}
	return nil
}

// Rolls back transactions that save invalid data to the database
func (o *Org) AfterSave(tx *gorm.DB) error {
	return o.Validate()
}

// Ensures the organization and its fields are valid
func (o Org) Validate() error {
	return validate.Struct(o)
}
