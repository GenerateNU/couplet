package controller

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database"
	"time"
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

// func (c Controller) GetAllUsers(ctx context.Context, limit int, offset int) ([]api.User, error) {
// 	var users []api.User

// 	err := c.database.Limit(limit).Offset(offset).Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

func (c Controller) PutUserById(ctx context.Context, userId api.UserId, updatedUser *api.User) (*api.User, error) {
	var user api.User

	// Do we create a new entry if entry doesn't exists? 
	err := c.database.Where("id = ?", userId).First(&user).Error 

	if err != nil {
		return nil, err
	}

	userUpdates := make(map[string]interface{})

	userUpdates["updatedAt"] = time.Now()

	if updatedUser.FirstName != "" {
		userUpdates["FirstName"] = updatedUser.FirstName
	}

	if updatedUser.LastName != "" {
		userUpdates["LastName"] = updatedUser.LastName
	}

	if updatedUser.Age >= 0 {
		userUpdates["Age"] = updatedUser.Age
	}

	if err := c.database.Model(&user).Updates(userUpdates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
