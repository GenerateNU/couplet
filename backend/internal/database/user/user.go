package user

import (
	"couplet/internal/api"
	"couplet/internal/database/event_swipe"
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user_id"
	"couplet/internal/database/user_swipe"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Perference struct {
	AgeMin uint8
	AgeMax uint8
	InterestedIn string
	Passions []string
}

type User struct {
	ID          user_id.UserID `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FirstName   string
	LastName    string
	Age         uint8
	Bio         string
	Gender			api.UserGender
	Preferences Perference
	Images      url_slice.UrlSlice
	UserSwipes  []user_swipe.UserSwipe
	EventSwipes []event_swipe.EventSwipe
	Matches     []*User `gorm:"many2many:user_matches"`
}

// Automatically generates a random ID if unset before creating
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if (u.ID == user_id.UserID{}) {
		u.ID = user_id.Wrap(uuid.New())
	}
	return
}
