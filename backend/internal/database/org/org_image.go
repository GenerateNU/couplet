package org

import (
	"couplet/internal/database/org_id"
	"time"

	"gorm.io/gorm"
)

type OrgImage struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
	Url       string    `validate:"required,url"`
	OrgID     org_id.OrgID
}

// Automatically rolls back transactions that save invalid data to the database
func (i *OrgImage) AfterSave(tx *gorm.DB) error {
	return i.Validate()
}

// Ensures the image and its fields are valid
func (i OrgImage) Validate() error {
	return validate.Struct(i)
}
