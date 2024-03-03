package org

import (
	"time"

	"gorm.io/gorm"
)

type OrgTag struct {
	ID        string `gorm:"primaryKey" validate:"required,min=1,max=255"`
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
	Orgs      []Org     `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:orgs2tags"`
}

// Automatically rolls back transactions that save invalid org tags to the database
func (t *OrgTag) AfterSave(tx *gorm.DB) error {
	return t.Validate()
}

func (t OrgTag) Validate() error {
	return validate.Struct(t)
}
