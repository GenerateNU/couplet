package org

import (
	"couplet/internal/database/org_id"
	"time"

	"gorm.io/gorm"
)

type OrgImage struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time    `validate:"gtefield=CreatedAt"`
	Url       string       `validate:"required,url"`
	OrgID     org_id.OrgID `validate:"required"`
}

// Automatically rolls back transactions that save invalid data to the database
func (oi *OrgImage) BeforeSave(tx *gorm.DB) error {
	return oi.Validate()
}

// Ensures the image and its fields are valid
func (oi OrgImage) Validate() error {
	return validate.Struct(oi)
}
