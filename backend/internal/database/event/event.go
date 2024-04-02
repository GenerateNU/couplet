package event

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID        event_id.EventID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Bio       string
	Images    url_slice.UrlSlice
	EventTags []EventTag `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:events2tags"`
	OrgID     org_id.OrgID
}

// Automatically generates a random ID if unset before creating
func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	if (e.ID == event_id.EventID{}) {
		e.ID = event_id.Wrap(uuid.New())
	}
	return
}

type EventTag struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Events    []Event `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:events2tags"`
}
