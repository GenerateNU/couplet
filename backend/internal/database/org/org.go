package org

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
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
func (o *Org) BeforeCreate(tx *gorm.DB) (err error) {
	if (o.ID == org_id.OrgID{}) {
		o.ID = org_id.Wrap(uuid.New())
	}
	return
}

type OrgTag struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Orgs      []Org `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:orgs2tags"`
}
