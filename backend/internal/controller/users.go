package controller

import (
	"context"
	"couplet/internal/api"
	db "couplet/internal/database/user"
	"time"
	"github.com/google/uuid"
)

// Searches the database for the specific user id
func (c Controller) GetUserById(ctx context.Context, params api.GetUserByIdParams) (db.User, error) {
	userId := params.UserId
	var user db.User
	result := c.database.First(&user, "id = ?", userId)
	if result.Error != nil {
		return db.User{}, result.Error
	} else {
		return user, nil
	}
}

// Updates a specific user in the database by their id
func (c Controller) PartialUpdateUserById(ctx context.Context, params api.PartialUpdateUserByIdParams) (db.User, error) {
	//Pull the fields from the parameter
	userId := params.UserId
	var user db.User
	result := c.database.First(&user, "id = ?", userId)
	if result.Error != nil {
		return db.User{}, result.Error
	}
	createdAt := params.CreatedAt
	firstName := params.FirstName
	lastName := params.LastName
	age := params.Age
	//Update the fields of the user if applicable
	if createdAt.IsSet() {
		user.CreatedAt = createdAt.Value
	}
	if firstName.IsSet() {
		user.FirstName = firstName.Value
	}
	if lastName.IsSet() {
		user.LastName = lastName.Value
	}
	if age.IsSet() {
		user.Age = uint8(age.Value)
	}
	if createdAt.IsSet() || firstName.IsSet() || lastName.IsSet() || age.IsSet() {
		user.UpdatedAt = time.Now()
	}
	return user, nil
}

// Gets all the users in the database based on the limit and offset 
func (c Controller) GetAllUsers(limit uint8, offset uint32) ([]api.User, error) {
	var users []api.User
	err := c.database.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c Controller) CreateUser(ctx context.Context, firstName string, lastName string, age uint8) (*api.User, error) {
	user := api.User{
		ID:        uuid.New(),
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

func (c Controller) PutUserById(ctx context.Context, updatedUser *api.User, userId string) (*api.User, error) {
	var user api.User
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

	if updatedUser.Age >= 0 {
		userUpdates["Age"] = updatedUser.Age
	}

	if err := c.database.Model(&user).Updates(userUpdates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
