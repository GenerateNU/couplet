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
func (h Handler) UsersPost(ctx context.Context, req *api.UserNoId) (api.UsersPostRes, error) {
	if h.logger != nil {
		h.logger.Info("POST /users")
	}
	pref := user.Preference{
		AgeMin:       req.Preference.AgeMin,
		AgeMax:       req.Preference.AgeMax,
		InterestedIn: string(req.Preference.InterestedIn),
	}

	fmt.Println(req)

	u, valErr, txErr := h.controller.CreateUser(user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Bio:       req.Bio,
		Gender:		 api.UserGender(req.Gender),
		Pronouns: req.Pronouns,
		Preference: pref,
		Location: req.Location,
		Work: req.Work,
		School: req.School,
		Height: req.Height,	
		PromptQuestion: req.PromptQuestion,
		PromptResponse: req.PromptResponse,
		RelationshipType: req.RelationshipType,
		Religion: req.Religion,
		PoliticalAffiliation: req.PoliticalAffiliation,
		AlcoholFrequency: req.AlcoholFrequency,
		SmokingFrequency: req.SmokingFrequency,
		DrugsFrequency: req.DrugsFrequency,
		CannabisFrequency: req.CannabisFrequency,
		InstagramUsername: req.InstagramUsername,		
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
		Gender:		 api.UserGender(req.Gender),
		Pronouns: req.Pronouns,
		Location: req.Location,
		Preference: req.Preference,
		Work: req.Work,
		School: req.School,
		Height: req.Height,	
		PromptQuestion: req.PromptQuestion,
		PromptResponse: req.PromptResponse,
		RelationshipType: req.RelationshipType,
		Religion: req.Religion,
		PoliticalAffiliation: req.PoliticalAffiliation,
		AlcoholFrequency: req.AlcoholFrequency,
		SmokingFrequency: req.SmokingFrequency,
		DrugsFrequency: req.DrugsFrequency,
		CannabisFrequency: req.CannabisFrequency,
		InstagramUsername: req.InstagramUsername,		
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

	pref := api.Preference{
		AgeMin:       u.Preference.AgeMin,
		AgeMax:       u.Preference.AgeMax,
		InterestedIn: api.PreferenceInterestedIn(u.Preference.InterestedIn),
	}


	res := api.User{
		ID:        u.ID.Unwrap(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Bio:       u.Bio,
		Gender:		 api.UserGender(u.Gender),
		Pronouns: u.Pronouns,
		Location: u.Location,
		Preference: pref,
		Work: u.Work,
		School: u.School,
		Height: u.Height,	
		PromptQuestion: u.PromptQuestion,
		PromptResponse: u.PromptResponse,
		RelationshipType: u.RelationshipType,
		Religion: u.Religion,
		PoliticalAffiliation: u.PoliticalAffiliation,
		AlcoholFrequency: u.AlcoholFrequency,
		SmokingFrequency: u.SmokingFrequency,
		DrugsFrequency: u.DrugsFrequency,
		CannabisFrequency: u.CannabisFrequency,
		InstagramUsername: u.InstagramUsername,		
		Images:    u.Images.Unwrap(),
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
		pref := api.Preference{
			AgeMin:       u.Preference.AgeMin,
			AgeMax:       u.Preference.AgeMax,
			InterestedIn: api.PreferenceInterestedIn(u.Preference.InterestedIn),
		}
		item := api.User{
			ID:        u.ID.Unwrap(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			Bio:       u.Bio,
			Images:    u.Images.Unwrap(),
			Gender:		 api.UserGender(u.Gender),
			Pronouns: u.Pronouns,
			Location: u.Location,
			Preference: pref,
			Work: u.Work,
			School: u.School,
			Height: u.Height,	
			PromptQuestion: u.PromptQuestion,
			PromptResponse: u.PromptResponse,
			RelationshipType: u.RelationshipType,
			Religion: u.Religion,
			PoliticalAffiliation: u.PoliticalAffiliation,
			AlcoholFrequency: u.AlcoholFrequency,
			SmokingFrequency: u.SmokingFrequency,
			DrugsFrequency: u.DrugsFrequency,
			CannabisFrequency: u.CannabisFrequency,
			InstagramUsername: u.InstagramUsername,		
		}
		res = append(res, item)
	}
	return res, nil
}

// Partially updates a user by its ID.
// PATCH /users/{id}
func (h Handler) UsersIDPatch(ctx context.Context, req *api.UserNoRequired, params api.UsersIDPatchParams) (api.UsersIDPatchRes, error) {
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

	res := api.User{
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
func (h Handler) UsersIDPut(ctx context.Context, req *api.User, params api.UsersIDPutParams) (api.UsersIDPutRes, error) {
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

	res := api.User{
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
func (h Handler) RecommendationsUsersGet(ctx context.Context, params api.RecommendationsUsersGetParams) ([]api.User, error) {
	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	users, err := h.controller.GetRecommendationsUser(user_id.Wrap(params.UserId), limit, offset)

	if err != nil {
		return nil, errors.New("Failed to get Recommended Users")
	}

	res := []api.User{}
	for _, u := range users {
		item := api.User{
			ID:        u.ID.Unwrap(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			Bio:       u.Bio,
			Images:    u.Images,
		}
		res = append(res, item)
	}

	return res, nil
}
