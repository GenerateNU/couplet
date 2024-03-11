package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserSwipe(t *testing.T) {
	// Set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, c)

	// Set up example user swipe data
	userID := user_id.Wrap(uuid.New())
	otherUserID := user_id.Wrap(uuid.New())
	exampleUserSwipe := user.UserSwipe{
		UserID:      userID,
		OtherUserID: otherUserID,
		Liked:       true,
	}

	// Expect the insert statement and create the user swipe
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`
        INSERT INTO "user_swipes" ("created_at","updated_at","user_id","other_user_id","liked")
        VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), exampleUserSwipe.UserID, exampleUserSwipe.OtherUserID, exampleUserSwipe.Liked).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// Call CreateUserSwipe and check for no error
	insertedUserSwipe, valErr, txErr := c.CreateUserSwipe(exampleUserSwipe)
	assert.Nil(t, valErr)
	assert.Nil(t, txErr)
	assert.NotNil(t, insertedUserSwipe)

	// Verify that the fields are correctly set
	assert.Equal(t, exampleUserSwipe.UserID, insertedUserSwipe.UserID)
	assert.Equal(t, exampleUserSwipe.OtherUserID, insertedUserSwipe.OtherUserID)
	assert.Equal(t, exampleUserSwipe.Liked, insertedUserSwipe.Liked)

	// Ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}
