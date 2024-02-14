package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Creates a new user.
// POST /users
func (h Handler) CreateUser(ctx context.Context, req *api.CreateUserReq) (api.CreateUserRes, error) {
	// TODO: Write tests
	h.logger.Info("POST /users")

	if req.Age < 18 {
		return nil, errors.New("must be at least 18 years old")
	}

	u, err := h.controller.CreateUser(req.FirstName, req.LastName, req.Age)
	// TODO: check for validation error from the controller and return 400
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	res := api.CreateUserCreated{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}

	return &res, nil
}

// Deletes a user by their user ID.
// DELETE /users/{id}
func (h Handler) DeleteUser(ctx context.Context, params api.DeleteUserParams) (api.DeleteUserRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("DELETE /users/%s", params.ID))
	u, err := h.controller.DeleteUser(user_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.DeleteUserOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}

	return &res, nil
}

// Gets a user by its ID.
// GET /users/{id}
func (h Handler) GetUser(ctx context.Context, params api.GetUserParams) (api.GetUserRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("GET /users/%s", params.ID))
	u, err := h.controller.GetUser(user_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.GetUserOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
	return &res, nil
}

// Gets multiple users.
// GET /users
func (h Handler) GetUsers(ctx context.Context, params api.GetUsersParams) ([]api.GetUsersOKItem, error) {
	// TODO: Write tests
	h.logger.Info("GET /users")
	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	users, err := h.controller.GetUsers(limit, offset)
	res := []api.GetUsersOKItem{}
	for _, u := range users {
		item := api.GetUsersOKItem{
			ID:        u.ID.Unwrap(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
		}
		res = append(res, item)
	}
	return res, err
}

// Updates the user based on their ID
// PUT /users/{userId}
func (h Handler) PutUser(ctx context.Context, updatedUser *api.PutUserReq, params api.PutUserParams) (api.PutUserRes, error) {
	// TODO: Write tests
	// Checks if user exists
	_, err := h.controller.GetUser(user_id.UserID(params.ID))
	alreadyExists := err == nil

	// TODO: Validate parameters
	if alreadyExists {
		responseUser, err := h.controller.SaveUser(user.User{FirstName: updatedUser.FirstName, LastName: updatedUser.LastName, Age: updatedUser.Age}, user_id.Wrap(params.ID))
		if err != nil {
			return &api.Error{
				Code:    400,
				Message: err.Error(),
			}, nil
		}
		updatedUser := api.PutUserOK{
			ID:        uuid.UUID(responseUser.ID),
			FirstName: responseUser.FirstName,
			LastName:  responseUser.LastName,
			Age:       responseUser.Age,
		}
		return &updatedUser, nil
	}

	responseUser, _ := h.controller.CreateUser(updatedUser.FirstName, updatedUser.LastName, updatedUser.Age)
	createdUser := api.PutUserCreated{
		ID:        uuid.UUID(responseUser.ID),
		FirstName: responseUser.FirstName,
		LastName:  responseUser.LastName,
		Age:       responseUser.Age,
	}
	return &createdUser, nil
}

// Updates the specific user at their ID
// GET /users/{userId}
func (h Handler) PatchUser(ctx context.Context, req *api.User, params api.PatchUserParams) (api.PatchUserRes, error) {
	h.logger.Info(fmt.Sprintf("PATCH /users/%s", params.ID))

	_, getErr := h.controller.GetUser(user_id.Wrap(params.ID))
	doesNotExist := getErr != nil
	if doesNotExist {
		return &api.PatchUserNotFound{
			Code:    404,
			Message: getErr.Error(),
		}, nil
	}

	var reqUser user.User
	reqUser.ID = user_id.Wrap(params.ID)
	if req.FirstName.Set {
		reqUser.FirstName = req.FirstName.Value
	}
	if req.LastName.Set {
		reqUser.LastName = req.LastName.Value
	}
	if req.Age.Set {
		reqUser.Age = req.Age.Value
	}

	u, valErr, txErr := h.controller.UpdateUser(reqUser)
	if valErr != nil {
		return &api.PatchUserBadRequest{
			Code:    400,
			Message: valErr.Error(),
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to update organization")
	}
	res := api.PatchUserOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
	return &res, nil
}
