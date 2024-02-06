package controller

import (
	"context"
	"couplet/internal/api"
	db "couplet/internal/database/user"
	"time"
)

// Searches the database for the specific user id
func (c Controller) GetUserById(ctx context.Context, params api.GetUserByIdParams) (db.User, error) {
	userId := params.UserId
	var user db.User
	result := c.database.First(&user, "id = ?", userId)
	if result.Error != nil {
		return db.User{}, result.Error
	} else {
		return user, nil
	}
}

// Updates a specific user in the database by their id
func (c Controller) PartialUpdateUserById(ctx context.Context, params api.PartialUpdateUserByIdParams) (db.User, error) {
	//Pull the fields from the parameter
	userId := params.UserId
	var user db.User
	result := c.database.First(&user, "id = ?", userId)
	if result.Error != nil {
		return db.User{}, result.Error
	}
	createdAt := params.CreatedAt
	firstName := params.FirstName
	lastName := params.LastName
	age := params.Age
	//Update the fields of the user if applicable
	if createdAt.IsSet() {
		user.CreatedAt = createdAt.Value
	}
	if firstName.IsSet() {
		user.FirstName = firstName.Value
	}
	if lastName.IsSet() {
		user.LastName = lastName.Value
	}
	if age.IsSet() {
		user.Age = uint8(age.Value)
	}
	if createdAt.IsSet() || firstName.IsSet() || lastName.IsSet() || age.IsSet() {
		user.UpdatedAt = time.Now()
	}
	return user, nil
}

// func (c Controller) GetAllUsers(ctx context.Context, limit int, offset int) ([]api.User, error) {
// 	var users []api.User

// 	err := c.database.Limit(limit).Offset(offset).Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
// func (c Controller) PutUserById(ctx context.Context, userId api.User.ID, updatedUser *api.User) (*api.User, error) {
// 	var user api.User

// 	// Do we create a new entry if entry doesn't exists?
// 	err := c.database.Where("id = ?", userId).First(&user).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	userUpdates := make(map[string]interface{})

// 	userUpdates["updatedAt"] = time.Now()

// 	if updatedUser.FirstName != "" {
// 		userUpdates["FirstName"] = updatedUser.FirstName
// 	}

// 	if updatedUser.LastName != "" {
// 		userUpdates["LastName"] = updatedUser.LastName
// 	}

// 	if updatedUser.Age >= 0 {
// 		userUpdates["Age"] = updatedUser.Age
// 	}

// 	if err := c.database.Model(&user).Updates(userUpdates).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
