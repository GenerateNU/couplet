package swipe

import (
	"couplet/internal/database/user_id"
	"time"
)

type UserSwipe struct {
	UserId      user_id.UserID `gorm:"primaryKey"` // Swipe sender
	UserSwipeId user_id.UserID `gorm:"primaryKey"` // Swipe receiver
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Liked       bool
}
