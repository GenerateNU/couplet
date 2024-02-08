package handler

import (
	"context"
	"couplet/internal/api"
	"errors"
	"fmt"

	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
)

// Creates a new user.
// POST /users
func (h Handler) CreateUser(ctx context.Context, user *api.CreateUserRequest) (api.CreateUserRes, error) {

	if user.Age < 18 {
		return nil, errors.New("must be at least 18 years old")
	}

	return h.controller.CreateUser(user.FirstName, user.LastName, user.Age)
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
	uuid, err := uuid.Parse(params.UserId.String())
	if err != nil {
		return nil, errors.New("invalid userID provided")
	}

	h.logger.Info(fmt.Sprintf("DELETE /users/%s", uuid))
	o, err := h.controller.DeleteUserById(uuid)
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	return &o, nil
}
