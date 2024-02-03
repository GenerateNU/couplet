package controller

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database"
)

func (c Controller) GetUserById(ctx context.Context, params api.GetUserByIdParams) (database.User, error) {
	userId := params.UserId
	var user database.User
	result := c.database.Limit(1).First(&user, "id = ?", userId)
	if result.Error != nil {
		return database.User{}, result.Error
	} else {
		return user, nil
	}
}
