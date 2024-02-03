package handler

import (
	"context"
	"couplet/internal/api"

	ht "github.com/ogen-go/ogen/http"
)

// GET (/users/:id) a single user by their id
// GET (/users) all users with pagination (ask if you don't know what this is)
// PUT (/users/:id) to completely update an existing user, returning the created object if successful
// PATCH (/users/:id) to partially update one or many fields of an existing user, returning the created object if successful

// Creates a new user.
// POST /users
// func (h Handler) CreateUser(ctx context.Context, user *api.User) (api.CreateUserRes, error) {
// 	return &api.User{}, ht.ErrNotImplemented
// }

// Gets all users.
// GET /users
// func (h Handler) GetAllUsers(ctx context.Context, params api.GetAllUsersParams) (api.GetAllUsersRes, error) {
// 	users, err := h.controller.GetAllUsers(ctx, params.Limit, params.Offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// Updates all the fields of an existing user
func (h Handler) PutUserById(ctx context.Context, params api.PutUserByIdParams, updatedUser *api.User) (database.PutUserByIdRes, error) {
	user, err := h.controller.PutUserById(ctx, params.UserId, updatedUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}


// Gets a user by their user ID.
// GET /users/{userId}
func (h Handler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	return h.controller.GetUserById(ctx, params)
}