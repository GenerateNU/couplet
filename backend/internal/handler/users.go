package handler

import (
	"context"
	"couplet/internal/api"

	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"
)

// Creates a new user.
// POST /users
func (h Handler) CreateUser(ctx context.Context, user *api.User) (api.CreateUserRes, error) {
	return &api.User{}, ht.ErrNotImplemented
}

// Gets all users.
// GET /users
func (h Handler) GetAllUsers(ctx context.Context, params api.GetAllUsersParams) ([]api.GetAllUsersOKItem, error) {
	// Gets all the users depending on page and limit
	limit := params.Limit.Value
	offset := params.Offset.Value
	users, err := h.controller.GetAllUsers(limit, offset)

	if err != nil {
		return nil, err
	}

	// Convert the database users into API users
	apiUsers := []api.GetAllUsersOKItem{}
	for _, user := range users {
		apiUser := api.GetAllUsersOKItem{
			ID:        uuid.UUID(user.ID),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Age:       user.Age,
		}
		apiUsers = append(apiUsers, apiUser)
	}

	return apiUsers, nil
}

// Updates the user based on their ID
// PUT /users/{userId}
func (h Handler) SaveUserById(ctx context.Context, updatedUser *api.User, params api.SaveUserByIdParams) (api.SaveUserByIdRes, error) {

	// Checks if user exists
	_, err := h.controller.GetUserById(ctx, api.GetUserByIdParams{UserId: uuid.UUID(params.UserId)})
	alreadyExists := err == nil

	if alreadyExists {
		responseUser, err := h.controller.SaveUserById(ctx, updatedUser, params.UserId.String())
		if err != nil {
			return &api.Error{
				Code:    400,
				Message: err.Error(),
			}, nil
		}
		updatedUser := api.SaveUserByIdOK{
			ID:        uuid.UUID(responseUser.ID),
			CreatedAt: responseUser.CreatedAt,
			UpdatedAt: responseUser.UpdatedAt,
			FirstName: responseUser.FirstName,
			LastName:  responseUser.LastName,
			Age:       responseUser.Age,
		}
		return &updatedUser, nil
	}

	// Replace with the CreateUser Endpoint
	responseUser, _ := h.controller.CreateUser(ctx, updatedUser.FirstName, updatedUser.LastName, updatedUser.Age)
	createdUser := api.SaveUserByIdCreated{
		ID:        uuid.UUID(responseUser.ID),
		CreatedAt: responseUser.CreatedAt,
		UpdatedAt: responseUser.UpdatedAt,
		FirstName: responseUser.FirstName,
		LastName:  responseUser.LastName,
		Age:       responseUser.Age,
	}
	return &createdUser, nil
}

// Gets a user by their user ID.
// GET /users/{userId}
func (h Handler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	//Grab the user from the database
	user, err := h.controller.GetUserById(ctx, params)
	//Return an empty user if there was an error
	if err != nil {
		return &api.User{}, err
	}
	//Convert the database user into an api user
	apiUser := api.User{
		ID:        uuid.UUID(user.ID),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
	}
	return &apiUser, nil
}

// Updates the specific user at their ID
// GET /users/{userId}
func (h Handler) PartialUpdateUserById(ctx context.Context, params api.PartialUpdateUserByIdParams) (api.PartialUpdateUserByIdRes, error) {
	//The user to be updated from the database
	user, err := h.controller.PartialUpdateUserById(ctx, params)
	//Return an empty user if there was an error
	if err != nil {
		return &api.User{}, err
	}
	//Convert the database user into an api user
	apiUser := api.User{
		ID:        uuid.UUID(user.ID),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
	}
	return &apiUser, nil
}
