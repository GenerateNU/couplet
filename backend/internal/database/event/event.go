package event

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"errors"
	"fmt"

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

// Automatically rolls back transactions that save invalid data to the database
func (e *Event) BeforeSave(tx *gorm.DB) error {
	return e.Validate()
}

// Ensures the event and its fields are valid
func (e Event) Validate() error {
	var timestampErr error
	if e.UpdatedAt.Before(e.CreatedAt) {
		return fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if len(e.Name) < 1 || 255 < len(e.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(e.Name))
	}
	var bioLengthErr error
	if len(e.Bio) < 1 || 255 < len(e.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(e.Bio))
	}
	var imageCountErr error
	if len(e.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(e.Images))
	}
	var tagsCountErr error
	if 5 < len(e.EventTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(e.EventTags))
	}
	var orgIdErr error
	if (e.OrgID == org_id.OrgID{}) {
		orgIdErr = fmt.Errorf("invalid org ID")
	}
	return errors.Join(timestampErr, nameLengthErr, bioLengthErr, imageCountErr, tagsCountErr, orgIdErr)
}
