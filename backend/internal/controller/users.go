package controller

import (
	"context"
	"couplet/internal/api"
	"fmt"

	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
	"gorm.io/gorm/clause"
)

// Creates a new user.
// POST /users
func (c Controller) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.User, error) {
	id := uuid.New();
	fmt.Println(id)

	user := api.User{
		ID:        api.UserId(id), // Generate a new UUID
		FirstName: req.FirstName.Value,
		LastName:  req.LastName.Value,
		Age:       req.Age.Value,
	}
	fmt.Println(user)

	result := c.database.Clauses(clause.Returning{}).Create(&user)

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
func (c Controller) DeleteUserById(ctx context.Context, params api.DeleteUserByIdParams) (api.DeleteUserByIdRes, error) {
	return nil, ht.ErrNotImplemented
}