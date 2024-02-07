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
func (h Handler) GetAllUsers(ctx context.Context, params api.GetAllUsersParams) (api.GetAllUsersRes, error) {
	// Gets all the users depending on page and limit
	// limit := int(params.Limit.Value)
    // offset := int(params.Offset.Value)
	// users, err := h.controller.GetAllUsers(limit, offset)
	
	// if err != nil {
    //     return nil, err
    // }

    // // Convert the database users into API users
    // apiUsers := []api.User{}
    // for _, user := range users {
    //     apiUser := api.User{
    //         ID:        uuid.UUID(user.ID),
    //         CreatedAt: user.CreatedAt,
    //         UpdatedAt: user.UpdatedAt,
    //         FirstName: user.FirstName,
    //         LastName:  user.LastName,
    //         Age:       user.Age,
    //     }
    //     apiUsers = append(apiUsers, apiUser)
    // }

    // return apiUsers, nil
	return nil, ht.ErrNotImplemented
}

// Updates the user based on their ID
// PUT /users/{userId}
func (h Handler) PutUserById(ctx context.Context, updatedUser *api.User, params api.PutUserByIdParams) (api.PutUserByIdRes, error) {
	// Update the user and return it
	user, err := h.controller.PutUserById(ctx, updatedUser, params)
	// Returns an empty user if there was an error
	if err != nil {
        return &api.User{}, err
    }

    // Convert the updated database user into an API user
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
