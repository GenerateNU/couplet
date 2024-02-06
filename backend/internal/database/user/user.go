package user

import (
	"couplet/internal/database/user/id"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        id.UserID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Age       uint8
}

// Automatically generates a random ID before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = id.New(uuid.New())
	return
}
