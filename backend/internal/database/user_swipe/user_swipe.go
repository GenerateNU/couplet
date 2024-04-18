package user_swipe

import (
	"couplet/internal/database/user_id"
	"time"
)

type UserSwipe struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      user_id.UserID `gorm:"index:pair,unique"` // Swipe sender
	OtherUserID user_id.UserID `gorm:"index:pair,unique"` // Swipe receiver
	Liked       bool
}
