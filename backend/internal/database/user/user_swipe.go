package user

import (
	"couplet/internal/database/user_id"
	"time"

	"gorm.io/gorm"
)

type UserSwipe struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time      `validate:"gtefield=CreatedAt"`
	UserID      user_id.UserID `gorm:"index:pair,unique" validate:"required"` // Swipe sender
	OtherUserID user_id.UserID `gorm:"index:pair,unique" validate:"required"` // Swipe receiver
	Liked       bool
}

// Automatically rolls back transactions that save invalid data to the database
func (us *UserSwipe) BeforeSave(tx *gorm.DB) error {
	return us.Validate()
}

// Ensures the event swipe and its fields are valid
func (us UserSwipe) Validate() error {
	return validate.Struct(us)
}
