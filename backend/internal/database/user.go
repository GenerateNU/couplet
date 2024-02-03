package database

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        UserID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Age       uint8
}

// Automatically generates a random UUID before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = UserID(uuid.New())
	return
}

type UserID uuid.UUID

func (id *UserID) Scan(src interface{}) error {
	baseId := uuid.UUID(*id)
	return (&baseId).Scan(src)
}

func (id UserID) Value() (driver.Value, error) {
	baseId := uuid.UUID(id)
	return baseId.Value()
}
