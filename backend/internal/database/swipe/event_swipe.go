package swipe

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user_id"
	"time"
)

type EventSwipe struct {
	UserId    user_id.UserID   `gorm:"primaryKey"`
	EventId   event_id.EventID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Liked     bool
}
