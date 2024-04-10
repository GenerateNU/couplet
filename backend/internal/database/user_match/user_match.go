package user_match

import (
	"couplet/internal/database/user_id"
	"time"
)

type UserMatch struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    user_id.UserID `gorm:"index:pair,unique"`
	MatchID   user_id.UserID `gorm:"index:pair,unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Viewed    bool
}
