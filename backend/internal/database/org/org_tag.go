package org

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrgTag struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Orgs      []Org `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:orgs2tags"`
}

// Automatically rolls back transactions that save invalid org tags to the database
func (t *OrgTag) BeforeSave(tx *gorm.DB) error {
	return t.Validate()
}

// Ensures the org tag is valid
func (t OrgTag) Validate() error {
	var idLengthErr error
	if len(t.ID) < 1 || 255 < len(t.ID) {
		idLengthErr = fmt.Errorf("invalid ID length of %d, must be in range [1,255]", len(t.ID))
	}
	var timestampErr error
	if t.UpdatedAt.Before(t.CreatedAt) {
		return fmt.Errorf("invalid timestamps")
	}
	return errors.Join(idLengthErr, timestampErr)
}
