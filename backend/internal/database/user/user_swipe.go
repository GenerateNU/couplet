package user

import (
	"couplet/internal/database/user_id"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSwipe struct {
	ID          uuid.UUID      `gorm:"primaryKey"`
	UserID      user_id.UserID `gorm:"index:user,unique" validate:"required"` // Swipe sender
	OtherUserID user_id.UserID `gorm:"index:user,unique" validate:"required"` // Swipe receiver
	Liked       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time `validate:"gtefield=CreatedAt"`
}

// Automatically generates a random ID if unset before creating
func (u *UserSwipe) BeforeCreate(tx *gorm.DB) error {
	if (u.ID == uuid.UUID{}) {
		u.ID = uuid.New()
	}
	return nil
}

// Automatically rolls back transactions that save invalid data to the database
func (u *UserSwipe) AfterSave(tx *gorm.DB) error {
	return u.Validate()
}

// Ensures the event swipe and its fields are valid
func (u UserSwipe) Validate() error {
	// TODO: Write tests
	return validate.Struct(u)
}
