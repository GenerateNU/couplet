package swipe

import (
	"couplet/internal/database/user_id"
	"time"
)

type UserSwipe struct {
	SenderId   user_id.UserID `gorm:"primaryKey"`
	ReceiverId user_id.UserID `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Liked      bool
}
