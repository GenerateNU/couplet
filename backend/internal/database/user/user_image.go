package user

import (
	"couplet/internal/database/user_id"
	"time"

	"gorm.io/gorm"
)

type UserImage struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time `validate:"gtefield=CreatedAt"`
	Url       string    `validate:"required,url"`
	UserID    user_id.UserID
}

// Automatically rolls back transactions that save invalid data to the database
func (i *UserImage) AfterSave(tx *gorm.DB) error {
	return i.Validate()
}

// Ensures the image and its fields are valid
func (i UserImage) Validate() error {
	return validate.Struct(i)
}
