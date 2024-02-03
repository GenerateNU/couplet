package controller

import (
	"context"
	"couplet/internal/api"

	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
)

// Creates a new user.
// POST /users
func (c Controller) CreateUser(ctx context.Context, firstName string, lastName string, age uint8) (api.CreateUserRes, error) {
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

// Deletes a user by their user ID.
// DELETE /users/{userId}
func (c Controller) DeleteUserById(ctx context.Context, userId string) (api.DeleteUserByIdRes, error) {
    // Retrieve the user before deleting
    user := &api.User{}
    // if err := c.database.First(user, userId).Error; err != nil {
    //     return nil, err
    // }

    // Delete the user
	result := c.database.Delete(&api.User{}, userId)

    if result.Error != nil {
        return nil, result.Error
    }

    // Return the deleted user
    return user, nil
}
