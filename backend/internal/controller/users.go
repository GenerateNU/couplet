package controller

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"
)

// Creates a new user in the database
func (c Controller) CreateUser(params user.User) (u user.User, valErr error, txErr error) {
	u = params
	fmt.Print(u)
	var timestampErr error
	if u.UpdatedAt.Before(u.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var firstNameLengthErr error
	if len(u.FirstName) < 1 || 255 < len(u.FirstName) {
		firstNameLengthErr = fmt.Errorf("invalid first name length of %d, must be in range [1,255]", len(u.FirstName))
	}
	var lastNameLengthErr error
	if len(u.LastName) < 1 || 255 < len(u.LastName) {
		lastNameLengthErr = fmt.Errorf("invalid last name length of %d, must be in range [1,255]", len(u.LastName))
	}
	var ageLimitErr error
	if u.Age < 18 {
		ageLimitErr = fmt.Errorf("invalid age of %d, must be 18 or greater", u.Age)
	}
	var bioLengthErr error
	if len(u.Bio) < 1 || 255 < len(u.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(u.Bio))
	}
	var imageCountErr error
	if len(u.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(u.Images))
	}
	valErr = errors.Join(timestampErr, firstNameLengthErr, lastNameLengthErr, ageLimitErr, bioLengthErr, imageCountErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Create(&u).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Deletes a user from the database by its ID
func (c Controller) DeleteUser(id user_id.UserID) (u user.User, txErr error) {
	u.ID = id

	tx := c.database.Begin()
	txErr = tx.Clauses(clause.Returning{}).Delete(&u).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Gets a user from the database by its ID
func (c Controller) GetUser(id user_id.UserID) (u user.User, txErr error) {
	txErr = c.database.First(&u, id).Error
	return
}

// Gets several users from the database with pagination
func (c Controller) GetUsers(limit uint8, offset uint32) (users []user.User, txErr error) {
	txErr = c.database.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	return
}

// Creates a new user or updates an existing user in the database
func (c Controller) SaveUser(params user.User) (u user.User, valErr error, txErr error) {
	u = params
	var timestampErr error
	if u.UpdatedAt.Before(u.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var firstNameLengthErr error
	if len(u.FirstName) < 1 || 255 < len(u.FirstName) {
		firstNameLengthErr = fmt.Errorf("invalid first name length of %d, must be in range [1,255]", len(u.FirstName))
	}
	var lastNameLengthErr error
	if len(u.LastName) < 1 || 255 < len(u.LastName) {
		lastNameLengthErr = fmt.Errorf("invalid last name length of %d, must be in range [1,255]", len(u.LastName))
	}
	var ageLimitErr error
	if u.Age < 18 {
		ageLimitErr = fmt.Errorf("invalid age of %d, must be 18 or greater", u.Age)
	}
	var bioLengthErr error
	if len(u.Bio) < 1 || 255 < len(u.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(u.Bio))
	}
	var imageCountErr error
	if len(u.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(u.Images))
	}
	valErr = errors.Join(timestampErr, firstNameLengthErr, lastNameLengthErr, ageLimitErr, bioLengthErr, imageCountErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Save(&u).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Update one or many fields of an existing user in the database
func (c Controller) UpdateUser(params user.User) (u user.User, valErr error, txErr error) {
	u = params
	var timestampErr error
	if u.UpdatedAt.Before(u.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var firstNameLengthErr error
	if 255 < len(u.FirstName) {
		firstNameLengthErr = fmt.Errorf("invalid first name length of %d, must be in range [1,255]", len(u.FirstName))
	}
	var lastNameLengthErr error
	if 255 < len(u.LastName) {
		lastNameLengthErr = fmt.Errorf("invalid last name length of %d, must be in range [1,255]", len(u.LastName))
	}
	var ageLimitErr error
	if u.Age != 0 && u.Age < 18 {
		ageLimitErr = fmt.Errorf("invalid age of %d, must be 18 or greater", u.Age)
	}
	var bioLengthErr error
	if 255 < len(u.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(u.Bio))
	}
	var imageCountErr error
	if len(u.Images) != 0 && len(u.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(u.Images))
	}
	valErr = errors.Join(timestampErr, firstNameLengthErr, lastNameLengthErr, ageLimitErr, bioLengthErr, imageCountErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Clauses(clause.Returning{}).Updates(&u).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Get Recommended Users
func (c Controller) GetRecommendationsUser(id user_id.UserID, limit int, offset int) ([]user.User, error) {
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
	if err := c.database.Order("random()").Where("id != ?", currentUser.ID).
		Joins("JOIN event_swipes ON users.id = event_swipes.user_id").
		Where("event_swipes.liked = ?", true).
		Where("event_swipes.event_id IN (?)", likedEventIDs).
		Where("age BETWEEN ? AND ?", currentUser.Preference.AgeMin, currentUser.Preference.AgeMax).
		Limit(int(limit)).Offset(int(offset)).
		Find(&recommendedUsers).Error; err != nil {
		return nil, err
	}

	return recommendedUsers, nil
}
