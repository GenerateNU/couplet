package controller

import (
	"couplet/internal/database/event_id"
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
func (c Controller) CreateUser(firstName string, lastName string, age uint8, images []user.UserImage) (user.User, error) {
	u := user.User{
		ID:        user_id.Wrap(uuid.New()),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Images:    images,
	}

	tx := c.database.Begin()

	result := tx.Create(&u)

	if result.Error != nil {
		tx.Rollback()
		return user.User{}, result.Error
	}

	tx.Commit()
	return u, nil
}

func (c Controller) SaveUser(updatedUser user.User, id user_id.UserID) (*user.User, error) {
	var user user.User

	tx := c.database.Begin()
	err := tx.First(&user, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	userUpdates := make(map[string]interface{})

	userUpdates["UpdatedAt"] = time.Now()
	userUpdates["Images"] = updatedUser.Images

	if updatedUser.FirstName != "" {
		userUpdates["FirstName"] = updatedUser.FirstName
	}

	if updatedUser.LastName != "" {
		userUpdates["LastName"] = updatedUser.LastName
	}

	if updatedUser.Age > 0 {
		userUpdates["Age"] = updatedUser.Age
	}

	if err := tx.Model(&user).Updates(userUpdates).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}

// Gets a user from the database by their ID
func (c Controller) GetUser(id user_id.UserID) (u user.User, txErr error) {
	txErr = c.database.First(&u, id).Error
	return
}

// Deletes a user from the database by its ID
func (c Controller) DeleteUser(id user_id.UserID) (u user.User, txErr error) {
	// TODO: Do this in one transaction
	tx := c.database.Begin()

	u, txErr = c.GetUser(id)
	if txErr != nil {
		tx.Rollback()
		return
	}

	txErr = tx.Delete(&u).Error
	if txErr != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// Updates a user in the database
func (c Controller) UpdateUser(params user.User) (u user.User, valErr error, txErr error) {
	// TODO: Write tests
	u = params
	valErr = u.Validate()

	tx := c.database.Begin()
	if valErr != nil {
		tx.Rollback()
		return
	}

	txErr = tx.Updates(&u).Error
	if txErr != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// Get Recommended Users
func (c Controller) GetRecommendationsUser(id user_id.UserID, limit uint8, offset uint32) ([]user.User, error) {
	// Return recommendedUsers
	var recommendedUsers []user.User

	// Get Current User
	var currentUser user.User
	err := c.database.Where("id = ?", id).Preload("EventSwipes").First(&currentUser).Error
	if err != nil {
		return nil, err
	}

	// Collect event IDs that the current user liked
	var likedEventIDs []event_id.EventID
	for _, eventSwipe := range currentUser.EventSwipes {
		if eventSwipe.Liked {
			likedEventIDs = append(likedEventIDs, eventSwipe.EventID)
		}
	}

	// Return all the Users that liked the same event as the current user
	if err := c.database.Where("id != ?", currentUser.ID).
		Joins("JOIN event_swipes ON users.id = event_swipes.user_id").
		Where("event_swipes.liked = ?", true).
		Where("event_swipes.event_id IN (?)", likedEventIDs).
		Limit(int(limit)).Offset(int(offset)).
		Find(&recommendedUsers).Error; err != nil {
		return nil, err
	}

	return recommendedUsers, nil
}
