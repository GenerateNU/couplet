package org

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org/id"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Org struct {
	ID        id.OrgID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Events    []event.Event
	Bio       string
}

// Automatically generates a random ID before creating an organization
func (o *Org) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = id.New(uuid.New())
	return
}
