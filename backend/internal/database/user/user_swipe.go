package user

import (
	"couplet/internal/database/user_id"
	"time"

	"gorm.io/gorm"
)

type UserSwipe struct {
	UserID      user_id.UserID `gorm:"primaryKey" validate:"required"` // Swipe sender
	OtherUserID user_id.UserID `gorm:"primaryKey" validate:"required"` // Swipe receiver
	OtherUser   User           `gorm:"foreignKey:OtherUserID"`
	Liked       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time `validate:"gtefield=CreatedAt"`
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
