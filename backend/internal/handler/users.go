package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"errors"
	"fmt"

	"net/url"

	"github.com/google/uuid"

)

// Creates a new user.
// POST /users
func (h Handler) UsersPost(ctx context.Context, req *api.UsersPostReq) (api.UsersPostRes, error) {
	if h.logger != nil {
		h.logger.Info("POST /users")
	}

	u, valErr, txErr := h.controller.CreateUser(user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Bio:       req.Bio,
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

	res := api.UsersPostCreated{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
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

	res := api.UsersIDDeleteOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
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

	res := api.UsersIDGetOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

// Gets multiple users.
// GET /users
func (h Handler) UsersGet(ctx context.Context, params api.UsersGetParams) ([]api.UsersGetOKItem, error) {
	if h.logger != nil {
		h.logger.Info("GET /users")
	}

	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	users, txErr := h.controller.GetUsers(limit, offset)
	if txErr != nil {
		return nil, errors.New("failed to get users")
	}
	res := []api.UsersGetOKItem{}
	for _, u := range users {
		item := api.UsersGetOKItem{
			ID:        u.ID.Unwrap(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			Bio:       u.Bio,
			Images:    u.Images.Unwrap(),
		}
		res = append(res, item)
	}
	return res, nil
}

// Partially updates a user by its ID.
// PATCH /users/{id}
func (h Handler) UsersIDPatch(ctx context.Context, req *api.UsersIDPatchReq, params api.UsersIDPatchParams) (api.UsersIDPatchRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PATCH /users/%s", params.ID))
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
	if req.Bio.Set {
		reqUser.Bio = req.Bio.Value
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

	res := api.UsersIDPatchOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

// Updates a user based on its ID.
// PUT /users/{id}
func (h Handler) UsersIDPut(ctx context.Context, req *api.UsersIDPutReq, params api.UsersIDPutParams) (api.UsersIDPutRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PUT /users/%s", params.ID))
	}

	u, valErr, txErr := h.controller.SaveUser(user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Bio:       req.Bio,
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

	res := api.UsersIDPutOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Images:    u.Images.Unwrap(),
	}
	return &res, nil
}

// RecommendationsUsersGet implements api.Handler.
func (h Handler) RecommendationsUsersGet(ctx context.Context, params api.RecommendationsUsersGetParams) ([]api.RecommendationsUsersGetOKItem, error) {
	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	users, err := h.controller.GetRecommendationsUser(user_id.Wrap(params.UserId), limit, offset)

	if err != nil {
		return nil, errors.New("Failed to get Recommended Users")
	}


	res := []api.RecommendationsUsersGetOKItem{}
	for _, u := range users {
		imagesUrl := []url.URL{}
		if u.Images != nil {
			for i := range u.Images {
				imagesUrl = append(imagesUrl, url.URL{Path: u.Images[i].Url})
			}
		}

		item := api.RecommendationsUsersGetOKItem{
			FirstName: api.NewOptString(u.FirstName),
			LastName:  api.NewOptString(u.LastName),
			Age:       api.NewOptUint8(u.Age),
			Images:    imagesUrl,
		}
		res = append(res, item)
	}
	
	return res, nil
}
