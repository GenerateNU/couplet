package controller

import (
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"time"

	"github.com/google/uuid"
)

// Gets all the users in the database based on the limit and offset
func (c Controller) GetUsers(limit uint8, offset uint32) ([]user.User, error) {
	var users []user.User
	err := c.database.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Creates a new user.
func (c Controller) CreateUser(firstName string, lastName string, age uint8) (user.User, error) {
	u := user.User{
		ID:        user_id.Wrap(uuid.New()),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	result := c.database.Create(&u)

	if result.Error != nil {
		return user.User{}, result.Error
	}

	return u, nil
}

func (c Controller) SaveUser(updatedUser user.User, id user_id.UserID) (user.User, error) {
	user := user.User{}
	if err := c.database.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	user.UpdatedAt = time.Now()
	user.ID = id

	if updatedUser.FirstName != "" {
		user.FirstName = updatedUser.FirstName
	}

	if updatedUser.LastName != "" {
		user.LastName = updatedUser.LastName
	}

	if updatedUser.Age > 0 {
		user.Age = updatedUser.Age
	}

	if err := c.database.Model(&user).Updates(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Gets a user from the database by their ID
func (c Controller) GetUser(id user_id.UserID) (u user.User, txErr error) {
	txErr = c.database.First(&u, id).Error
	return
}

// Deletes a user from the database by its ID
func (c Controller) DeleteUser(id user_id.UserID) (u user.User, txErr error) {
	// TODO: Do this in one transaction
	u, txErr = c.GetUser(id)
	if txErr != nil {
		return
	}
	txErr = c.database.Delete(&u).Error
	return
}

// Updates a user in the database
func (c Controller) UpdateUser(params user.User) (u user.User, valErr error, txErr error) {
	// TODO: Write tests
	u = params
	// TODO: validate fields
	if valErr == nil {
		txErr = c.database.Updates(&u).Error
	}
	return
}
