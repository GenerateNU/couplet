package user

import (
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user_id"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          user_id.UserID `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FirstName   string
	LastName    string
	Age         uint8
	Bio         string
	Images      url_slice.UrlSlice
	UserSwipes  []UserSwipe
	EventSwipes []EventSwipe
	Matches     []*User `gorm:"many2many:user_matches"`
}

// Automatically generates a random ID if unset before creating
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if (u.ID == user_id.UserID{}) {
		u.ID = user_id.Wrap(uuid.New())
	}
	return nil
}

// Automatically rolls back transactions that save invalid data to the database
func (u *User) BeforeSave(tx *gorm.DB) error {
	return u.Validate()
}

// Ensures the user and its fields are valid
func (u User) Validate() error {
	var timestampErr error
	if u.UpdatedAt.Before(u.CreatedAt) {
		return fmt.Errorf("invalid timestamps")
	}
	var firstNameLengthErr error
	if len(u.FirstName) < 1 || 255 < len(u.FirstName) {
		firstNameLengthErr = fmt.Errorf("invalid first name length of %d, must be in range [1,255]", len(u.FirstName))
	}
	var lastNameLengthErr error
	if len(u.LastName) < 1 || 255 < len(u.LastName) {
		lastNameLengthErr = fmt.Errorf("invalid last name length of %d, must be in range [1,255]", len(u.LastName))
	}
	var ageLimitErr error
	if u.Age < 18 {
		ageLimitErr = fmt.Errorf("invalid age of %d, must be 18 or greater", u.Age)
	}
	var bioLengthErr error
	if len(u.Bio) < 1 || 255 < len(u.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(u.Bio))
	}
	var imageCountErr error
	if len(u.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(u.Images))
	}
	return errors.Join(timestampErr, firstNameLengthErr, lastNameLengthErr, ageLimitErr, bioLengthErr, imageCountErr)
}
