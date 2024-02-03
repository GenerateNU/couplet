package handler

import (
	"context"
	"couplet/internal/api"

	ht "github.com/ogen-go/ogen/http"
)

// Creates a new user.
// POST /users
func (h Handler) CreateUser(ctx context.Context, user *api.User) (api.CreateUserRes, error) {
	return &api.User{}, ht.ErrNotImplemented
}

// Gets all users.
// GET /users
func (h Handler) GetAllUsers(ctx context.Context) ([]api.User, error) {
	return []api.User{}, ht.ErrNotImplemented
}

// Gets a user by their user ID.
// GET /users/{userId}
func (h Handler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	return h.controller.GetUserById(ctx, params)
}
