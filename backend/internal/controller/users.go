package controller

import (
	"context"
	"couplet/internal/api"
	db "couplet/internal/database/user"
	"time"
)

// Searches the database for the specific user id
func (c Controller) GetUserById(ctx context.Context, params api.GetUserByIdParams) (db.User, error) {
	userId := params.UserId
	var user db.User
	result := c.database.First(&user, "id = ?", userId)
	return user, result.Error
}

// Updates a specific user in the database by their id
func (c Controller) PartialUpdateUserById(ctx context.Context, params api.PartialUpdateUserByIdParams) (db.User, error) {
	//Pull the fields from the parameter
	userId := params.UserId
	var user db.User
	result := c.database.First(&user, "id = ?", userId)
	//Update the fields of the user if applicable
	if params.CreatedAt.IsSet() {
		user.CreatedAt = params.CreatedAt.Value
	}
	if params.FirstName.IsSet() {
		user.FirstName = params.FirstName.Value
	}
	if params.LastName.IsSet() {
		user.LastName = params.LastName.Value
	}
	if params.Age.IsSet() {
		user.Age = uint8(params.Age.Value)
	}
	if params.CreatedAt.IsSet() || params.FirstName.IsSet() || params.LastName.IsSet() || params.Age.IsSet() {
		user.UpdatedAt = time.Now()
	}
	return user, result.Error
}

// Gets all the users in the database based on the limit and offset
func (c Controller) GetAllUsers(limit uint8, offset uint32) ([]db.User, error) {
	var users []db.User
	err := c.database.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c Controller) CreateUser(ctx context.Context, firstName string, lastName string, age uint8) (*db.User, error) {
	user := db.User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	result := c.database.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (c Controller) SaveUserById(ctx context.Context, updatedUser *api.User, userId string) (*db.User, error) {
	var user db.User
	err := c.database.First(&user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	userUpdates := make(map[string]interface{})

	userUpdates["UpdatedAt"] = time.Now()

	if updatedUser.FirstName != "" {
		userUpdates["FirstName"] = updatedUser.FirstName
	}

	if updatedUser.LastName != "" {
		userUpdates["LastName"] = updatedUser.LastName
	}

	if updatedUser.Age > 0 {
		userUpdates["Age"] = updatedUser.Age
	}

	if err := c.database.Model(&user).Updates(userUpdates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
