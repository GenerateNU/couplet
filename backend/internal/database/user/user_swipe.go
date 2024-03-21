package user

import (
	"couplet/internal/database/user_id"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserSwipe struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      user_id.UserID `gorm:"index:pair,unique"` // Swipe sender
	OtherUserID user_id.UserID `gorm:"index:pair,unique"` // Swipe receiver
	Liked       bool
}

// Automatically rolls back transactions that save invalid data to the database
func (us *UserSwipe) BeforeSave(tx *gorm.DB) error {
	return us.Validate()
}

// Ensures the event swipe and its fields are valid
func (us UserSwipe) Validate() error {
	var timestampErr error
	if us.UpdatedAt.Before(us.CreatedAt) {
		return fmt.Errorf("invalid timestamps")
	}
	var userIdErr error
	if (us.UserID == user_id.UserID{}) {
		userIdErr = fmt.Errorf("invalid user ID for swipe sender")
	}
	var otherUserIdErr error
	if (us.OtherUserID == user_id.UserID{}) {
		otherUserIdErr = fmt.Errorf("invalid user ID for swipe receiver")
	}
	return errors.Join(timestampErr, userIdErr, otherUserIdErr)
}
