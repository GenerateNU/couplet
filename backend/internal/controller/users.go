package controller

import (
	"context"
	"couplet/internal/api"

	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
)

// Creates a new user.
// POST /users
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

// Gets all users.
// GET /users
func (c Controller) GetAllUsers(ctx context.Context) ([]api.User, error) {
	return []api.User{}, ht.ErrNotImplemented
}

// Gets a user by their user ID.
// GET /users/{userId}
func (c Controller) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	return &api.User{}, ht.ErrNotImplemented
}

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
