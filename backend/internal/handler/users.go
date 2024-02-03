package handler

import (
	"context"
	"couplet/internal/api"
	"errors"

	ht "github.com/ogen-go/ogen/http"
)

// Creates a new user.
// POST /users
func (h Handler) CreateUser(ctx context.Context, user *api.CreateUserRequest) (api.CreateUserRes, error) {

	if user.Age.Value < 18 {
		return nil, errors.New("must be at least 18 years old")
	}

	return h.controller.CreateUser(ctx, user.FirstName.Value, user.LastName.Value, user.Age.Value)
}

// Gets all users.
// GET /users
func (h Handler) GetAllUsers(ctx context.Context) ([]api.User, error) {
	return []api.User{}, ht.ErrNotImplemented
}

// Gets a user by their user ID.
// GET /users/{userId}
func (h Handler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	return &api.User{}, ht.ErrNotImplemented
}

// Deletes a user by their user ID.
// DELETE /users/{userId}
func (h Handler) DeleteUserById(ctx context.Context, params api.DeleteUserByIdParams) (api.DeleteUserByIdRes, error) {
	if params.UserId.String() == "" {
		return nil, errors.New("user ID must not be empty")
	}

	return h.controller.DeleteUserById(ctx, params.UserId.String());
}
