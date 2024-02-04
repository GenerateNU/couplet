package event

import (
	"couplet/internal/database/event/id"
	orgID "couplet/internal/database/org/id"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID        id.EventID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Bio       string
	OrgID     orgID.OrgID
}

// Automatically generates a random UUID before creating an event
func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = id.New(uuid.New())
	return
}
