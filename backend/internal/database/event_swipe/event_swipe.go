package event_swipe

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user_id"
	"time"
)

type EventSwipe struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    user_id.UserID   `gorm:"index:pair,unique"`
	EventID   event_id.EventID `gorm:"index:pair,unique"`
	Liked     bool
}
