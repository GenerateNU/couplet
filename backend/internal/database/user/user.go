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

type Preference struct {
	AgeMin uint8 	`gorm:"column:age_min"`
	AgeMax uint8 	`gorm:"column:age_max"`
	InterestedIn string `gorm:"column:interested_in"`
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
	Preference Preference	`gorm:"embedded"`
	Pronouns		string `json:"pronouns"`
	Location		string `json:"location"`
	School			string `json:"school"`
	Work 				string
	Height			uint8
	PromptQuestion string
	PromptResponse string
	RelationshipType string
	Religion			string	
	PoliticalAffiliation	string
	AlcoholFrequency			string
	SmokingFrequency			string
	DrugsFrequency				string
	CannabisFrequency		string
	InstagramUsername	string
	Images      url_slice.UrlSlice
	UserSwipes  []user_swipe.UserSwipe
	EventSwipes []event_swipe.EventSwipe
	Matches     []*User `gorm:"many2many:user_matches"`
}

type PassionTag struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users    []User `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;many2many:users2tags"`
}


// Automatically generates a random ID if unset before creating
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if (u.ID == user_id.UserID{}) {
		u.ID = user_id.Wrap(uuid.New())
	}
	return
}
