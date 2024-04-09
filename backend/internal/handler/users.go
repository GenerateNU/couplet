package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"errors"
	"fmt"
)

// Creates a new user.
// POST /users
func (h Handler) UsersPost(ctx context.Context, req *api.User) (api.UsersPostRes, error) {
	if h.logger != nil {
		h.logger.Info("POST /users")
	}

	u, valErr, txErr := h.controller.CreateUser(user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Bio:       req.Bio,
		Gender:    req.Gender, // Convert req.Gender to string
		Images:    url_slice.Wrap(req.Images),
	})
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: "failed to validate user",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to create user")
	}


	res := api.User{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Gender:		 u.Gender,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

// Deletes a user by its ID.
// DELETE /users/{id}
func (h Handler) UsersIDDelete(ctx context.Context, params api.UsersIDDeleteParams) (api.UsersIDDeleteRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("DELETE /users/%s", params.ID))
	}

	u, txErr := h.controller.DeleteUser(user_id.Wrap(params.ID))
	if txErr != nil {
		return &api.Error{
			Code:    404,
			Message: "user not found",
		}, nil
	}

	res := api.User{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Gender: 	 u.Gender,
		Bio:       u.Bio,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

// Gets a user by its ID.
// GET /users/{id}
func (h Handler) UsersIDGet(ctx context.Context, params api.UsersIDGetParams) (api.UsersIDGetRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("GET /users/%s", params.ID))
	}

	u, txErr := h.controller.GetUser(user_id.Wrap(params.ID))
	if txErr != nil {
		return &api.Error{
			Code:    404,
			Message: "user not found",
		}, nil
	}

	res := api.User{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Images:    u.Images.Unwrap(),
		Gender:		 u.Gender,
	}
	return &res, nil
}

// Gets multiple users.
// GET /users
func (h Handler) UsersGet(ctx context.Context, params api.UsersGetParams) ([]api.User, error) {
	if h.logger != nil {
		h.logger.Info("GET /users")
	}

	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	users, txErr := h.controller.GetUsers(limit, offset)
	if txErr != nil {
		return nil, errors.New("failed to get users")
	}
	res := []api.User{}
	for _, u := range users {
		item := api.User{
			ID:        u.ID.Unwrap(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			Bio:       u.Bio,
			Images:    u.Images.Unwrap(),
			Gender:		 u.Gender,
		}
		res = append(res, item)
	}
	return res, nil
}

// Partially updates a user by its ID.
// PATCH /users/{id}
func (h Handler) UsersIDPatch(ctx context.Context, req *api.User, params api.UsersIDPatchParams) (api.UsersIDPatchRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PATCH /users/%s", params.ID))
	}

	var reqUser user.User
	reqUser.ID = user_id.Wrap(params.ID)
	if req.FirstName != "" { // TODO fix 
		reqUser.FirstName = req.FirstName
	}
	if req.LastName!= "" { // TODO fix 
		reqUser.LastName = req.LastName
	}
	if req.Age > 0 { // TODO fix 
		reqUser.Age = req.Age
	}
	if req.Bio!= "" { // TODO fix 
		reqUser.Bio = req.Bio
	}
	reqUser.Images = url_slice.Wrap(req.Images)

	u, valErr, txErr := h.controller.UpdateUser(reqUser)
	if valErr != nil {
		return &api.UsersIDPatchBadRequest{
			Code:    400,
			Message: "failed to validate user",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to update user")
	}

	res := api.User{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Gender:		 u.Gender,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

// Updates a user based on its ID.
// PUT /users/{id}
func (h Handler) UsersIDPut(ctx context.Context, req *api.User, params api.UsersIDPutParams) (api.UsersIDPutRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PUT /users/%s", params.ID))
	}

	u, valErr, txErr := h.controller.SaveUser(user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Bio:       req.Bio,
		Gender:		 req.Gender,
		Images:    url_slice.Wrap(req.Images),
	})
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: "failed to validate user",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to save user")
	}

	res := api.User{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Gender:		 u.Gender,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

func (h Handler) UsersReccomendationGet(ctx context.Context, params api.UsersReccomendationGetParams) ([]api.User, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("GET /users/%s/reccomendations", params.UserId))
	}

	users, txErr := h.controller.GetReccomendations(user_id.Wrap(params.UserId))
	if txErr != nil {
		return nil, errors.New("failed to get reccomendations")
	}
	res := []api.User{}
	for _, u := range users {
		item := api.User{
			ID:        u.ID.Unwrap(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			Bio:       u.Bio,
			Gender:		 u.Gender,
			Images:    u.Images.Unwrap(),
		}
		res = append(res, item)
	}
	return res, nil
}