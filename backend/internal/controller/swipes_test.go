package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserSwipeWithNoLike(t *testing.T) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)

	// Define test data
	user1ID := user_id.Wrap(uuid.New())
	user2ID := user_id.Wrap(uuid.New())
	userSwipe1 := user.UserSwipe{UserID: user1ID, OtherUserID: user2ID, Liked: true}

	// Mock database interactions
	mock.ExpectBegin()

	// Expectation: insert the first user swipe
	mock.ExpectQuery(`^INSERT INTO "user_swipes"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userSwipe1.UserID, userSwipe1.OtherUserID, userSwipe1.Liked).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Expectation: check if there's a swipe from the other user
	// Return no rows to simulate "record not found"
	mock.ExpectQuery(`^SELECT \* FROM "user_swipes" WHERE user_id = \$1 AND other_user_id = \$2 AND liked = \$3`).
		WithArgs(user2ID, user1ID, true).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "other_user_id", "liked"}))

	mock.ExpectCommit()

	// Execute the function being tested
	_, valErr, txErr := c.CreateUserSwipe(userSwipe1)

	// Assertions
	assert.Nil(t, valErr)
	assert.Nil(t, txErr)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestCreateUserSwipeWithReciprocalLike(t *testing.T) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)

	// Define test data
	user1ID := user_id.Wrap(uuid.New())
	user2ID := user_id.Wrap(uuid.New())
	userSwipe1 := user.UserSwipe{UserID: user1ID, OtherUserID: user2ID, Liked: true}

	// Mock database interactions
	mock.ExpectBegin()

	// Expectation: insert the first user swipe
	mock.ExpectQuery(`^INSERT INTO "user_swipes"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userSwipe1.UserID, userSwipe1.OtherUserID, userSwipe1.Liked).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Expectation: check if there's a reciprocal swipe from the other user
	// Return a row to simulate that the other user has already liked the first user
	mock.ExpectQuery(`^SELECT \* FROM "user_swipes" WHERE user_id = \$1 AND other_user_id = \$2 AND liked = \$3`).
		WithArgs(user2ID, user1ID, true).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "other_user_id", "liked"}).AddRow(1, user2ID, user1ID, true))

	mock.ExpectCommit()

	// Execute the function being tested
	_, valErr, txErr := c.CreateUserSwipe(userSwipe1)

	// Assertions
	assert.Nil(t, valErr)
	assert.Nil(t, txErr)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}
