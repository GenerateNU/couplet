package user

import (
	"couplet/internal/database/user_id"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type User struct {
	ID        user_id.UserID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Age       uint8
}

// Automatically generates a random ID if unset before creating
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if (u.ID == user_id.UserID{}) {
		u.ID = user_id.Wrap(uuid.New())
	}
	return nil
}

// Automatically rolls back transactions that save invalid data to the database
func (u *User) AfterSave(tx *gorm.DB) error {
	return u.Validate()
}

// Ensures the user and its fields are valid
func (u User) Validate() error {
	// TODO: Write tests
	return validate.Struct(u)
}
