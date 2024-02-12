package controller

import (
	"context"
	"couplet/internal/api"=
	db "couplet/internal/database/user"
	"time"
)

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

// Creates a new user.
func (c Controller) CreateUser(firstName string, lastName string, age uint8) (*api.User, error) {
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

// Gets a user from the database by their ID
func (c Controller) GetUser(id uuid.UUID) (u api.User, txErr error) {
	txErr = c.database.First(&u, id).Error
	return
}

// Deletes a user by their user ID.
// DELETE /users/{userId}
func (c Controller) DeleteUserById(id uuid.UUID) (u api.User, txErr error) {
	u, txErr = c.GetUser(id)

	if txErr != nil {
		return api.User{}, txErr
	}

	txErr = c.database.Delete(&u).Error
	return
}