package org

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Org struct {
	ID        org_id.OrgID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Bio       string
	Images    url_slice.UrlSlice
	OrgTags   []OrgTag      `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:orgs2tags"`
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
func (o *Org) BeforeSave(tx *gorm.DB) error {
	return o.Validate()
}

// Ensures the organization and its fields are valid
func (o Org) Validate() error {
	var timestampErr error
	if o.UpdatedAt.Before(o.CreatedAt) {
		return fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if len(o.Name) < 1 || 255 < len(o.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(o.Name))
	}
	var bioLengthErr error
	if len(o.Bio) < 1 || 255 < len(o.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(o.Bio))
	}
	var imageCountErr error
	if len(o.Images) < 1 || 4 < len(o.Images) {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be in range [1,4]", len(o.Images))
	}
	var tagsCountErr error
	if 5 < len(o.OrgTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(o.OrgTags))
	}
	return errors.Join(timestampErr, nameLengthErr, bioLengthErr, imageCountErr, tagsCountErr)
}
