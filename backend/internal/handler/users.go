package handler

import (
	"context"
	"couplet/internal/api"
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
	// TODO: Write tests
	h.logger.Info("POST /users")

	if req.Age < 18 {
		return nil, errors.New("must be at least 18 years old")
	}
	images := make([]user.UserImage, len(req.Images))
	for i, v := range req.Images {
		images[i] = user.UserImage{Url: v.String()}
	}

	u, err := h.controller.CreateUser(req.FirstName, req.LastName, req.Age, images)
	// TODO: check for validation error from the controller and return 400
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	res := api.UsersPostCreated{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Images:    req.Images,
	}

	return &res, nil
}

// Deletes a user by their user ID.
// DELETE /users/{id}
func (h Handler) UsersIDDelete(ctx context.Context, params api.UsersIDDeleteParams) (api.UsersIDDeleteRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("DELETE /users/%s", params.ID))
	u, err := h.controller.DeleteUser(user_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.UsersIDDeleteOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}

	return &res, nil
}

// Gets a user by its ID.
// GET /users/{id}
func (h Handler) UsersIDGet(ctx context.Context, params api.UsersIDGetParams) (api.UsersIDGetRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("GET /users/%s", params.ID))
	u, err := h.controller.GetUser(user_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.UsersIDGetOK{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
	return &res, nil
}

// Gets multiple users.
// GET /users
func (h Handler) UsersGet(ctx context.Context, params api.UsersGetParams) ([]api.UsersGetOKItem, error) {
	// TODO: Write tests
	h.logger.Info("GET /users")
	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	users, err := h.controller.GetUsers(limit, offset)
	res := []api.UsersGetOKItem{}
	for _, u := range users {
		item := api.UsersGetOKItem{
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
func (h Handler) UsersIDPut(ctx context.Context, updatedUser *api.UsersIDPutReq, params api.UsersIDPutParams) (api.UsersIDPutRes, error) {
	// TODO: Write tests
	// Checks if user exists
	_, err := h.controller.GetUser(user_id.Wrap(params.ID))
	alreadyExists := err == nil

	images := make([]user.UserImage, len(updatedUser.Images))
	for i, v := range updatedUser.Images {
		images[i] = user.UserImage{Url: v.String()}
	}
	// TODO: Validate parameters
	if alreadyExists {
		responseUser, err := h.controller.SaveUser(user.User{FirstName: updatedUser.FirstName, LastName: updatedUser.LastName, Age: updatedUser.Age, Images: images}, user_id.Wrap(params.ID))
		if err != nil {
			return &api.Error{
				Code:    400,
				Message: err.Error(),
			}, nil
		}
		updatedUser := api.UsersIDPutOK{
			ID:        uuid.UUID(responseUser.ID),
			FirstName: responseUser.FirstName,
			LastName:  responseUser.LastName,
			Age:       responseUser.Age,
		}
		return &updatedUser, nil
	}

	responseUser, _ := h.controller.CreateUser(updatedUser.FirstName, updatedUser.LastName, updatedUser.Age, images)
	createdUser := api.UsersIDPutCreated{
		ID:        uuid.UUID(responseUser.ID),
		FirstName: responseUser.FirstName,
		LastName:  responseUser.LastName,
		Age:       responseUser.Age,
	}
	return &createdUser, nil
}

// Updates the specific user at their ID
// GET /users/{userId}
func (h Handler) UsersIDPatch(ctx context.Context, req *api.User, params api.UsersIDPatchParams) (api.UsersIDPatchRes, error) {
	h.logger.Info(fmt.Sprintf("PATCH /users/%s", params.ID))

	_, getErr := h.controller.GetUser(user_id.Wrap(params.ID))
	doesNotExist := getErr != nil
	if doesNotExist {
		return &api.UsersIDPatchNotFound{
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

	if req.Images != nil {
		images := make([]user.UserImage, len(req.Images))
		for i, v := range req.Images {
			images[i] = user.UserImage{Url: v.String()}
		}
		reqUser.Images = images
	}

	u, valErr, txErr := h.controller.UpdateUser(reqUser)
	if valErr != nil {
		return &api.UsersIDPatchBadRequest{
			Code:    400,
			Message: valErr.Error(),
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
