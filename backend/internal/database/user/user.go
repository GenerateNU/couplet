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
	ID          user_id.UserID `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time   `validate:"gtefield=CreatedAt"`
	FirstName   string      `validate:"required,min=1,max=255"`
	LastName    string      `validate:"required,min=1,max=255"`
	Age         uint8       `validate:"required,min=18"`
	Images      []UserImage `validate:"max=5"`
	UserSwipes  []UserSwipe
	EventSwipes []EventSwipe
	Matches     []*User `gorm:"many2many:user_matches;"`
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
	return validate.Struct(u)
}
