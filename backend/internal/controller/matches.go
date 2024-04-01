package controller

import (
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
)

// Gets a specified user's matches from the database
func (c Controller) GetUserMatches(id user_id.UserID) (matches []*user.User, txErr error) {
	var user user.User
	txErr = c.database.First(&user, id).Error
	if txErr != nil {
		return
	}

	txErr = c.database.Model(&user).Association("Matches").Find(&matches)
	return
}
