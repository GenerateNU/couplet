package event

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Event struct {
	ID        event_id.EventID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Bio       string
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
func (e *Event) AfterSave(tx *gorm.DB) error {
	return e.Validate()
}

// Ensures the event and its fields are valid
func (e Event) Validate() error {
	// TODO: Write tests
	return validate.Struct(e)
}
