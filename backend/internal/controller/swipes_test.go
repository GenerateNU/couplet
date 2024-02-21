package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/event_id"
	"couplet/internal/database/swipe"
	"couplet/internal/database/user_id"

	"testing"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateEventSwipe(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated eventID
	rec := dbtesting.NewValueRecorder()

	// set up example event data
	event_id := event_id.Wrap(uuid.New())
	user_id := user_id.Wrap(uuid.New())

	// set up example event data
	// orgId := org_id.Wrap(uuid.New())
	// exampleEvent := event.Event{
	// 	Name:  "Big event",
	// 	Bio:   "Event description",
	// 	OrgID: orgId,
	// }

	exampleEventSwipe := swipe.EventSwipe{
		UserId:  user_id,
		EventId: event_id,
		Liked:   true,
	}

	// expect the insert statement and create the event
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "events" ("id","created_at","updated_at","user_id","event_id","liked")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idOne"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventSwipe.UserId, exampleEventSwipe.EventId, exampleEventSwipe.Liked).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventSwipe, err := c.CreateEventSwipe(exampleEventSwipe)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	assert.Equal(t, insertedEventSwipe.UserId, exampleEventSwipe.UserId)
	assert.Equal(t, insertedEventSwipe.EventId, exampleEventSwipe.EventId)
	assert.Equal(t, insertedEventSwipe.Liked, exampleEventSwipe.Liked)

	// // create a second event with the same data to show that repeated POST calls always creates new events
	// mock.ExpectBegin()
	// mock.ExpectExec(regexp.QuoteMeta(`
	// INSERT INTO "events" ("id","created_at","updated_at","name","bio","org_id")
	// 	VALUES ($1,$2,$3,$4,$5,$6)`)).
	// 	WithArgs(rec.Record("idTwo"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventTwo.Name, exampleEventTwo.Bio, exampleEventTwo.OrgID).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	// insertedEventTwo, err := c.CreateEvent(exampleEventTwo)
	// assert.Nil(t, err)

	// assert.Equal(t, insertedEventTwo.Name, exampleEventTwo.Name)
	// assert.Equal(t, insertedEventTwo.Bio, exampleEventTwo.Bio)

	// // IMPORTANT! assert that internally, the second event id is not the same as the first event id
	// assert.NotEqual(t, insertedEventTwo.ID, insertedEventOne.ID)

	// // ensure that all expectations are met in the mock
	// errExpectations := mock.ExpectationsWereMet()
	// assert.Nil(t, errExpectations)
}

func TestCreateUserSwipe(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated eventID
	rec := dbtesting.NewValueRecorder()

	// set up example event data
	other_user_id := user_id.Wrap(uuid.New())
	user_id := user_id.Wrap(uuid.New())

	// set up example event data
	exampleUserSwipe := swipe.UserSwipe{
		UserId:      user_id,
		UserSwipeId: other_user_id,
		Liked:       true,
	}

	// expect the insert statement and create the event
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "user_swipes" ("id","created_at","updated_at","user_id","user_swipe_id","liked")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idOne"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleUserSwipe.UserId, exampleUserSwipe.UserSwipeId, exampleUserSwipe.Liked).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedUserSwipe, err := c.CreateUserSwipe(exampleUserSwipe)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	assert.Equal(t, insertedUserSwipe.UserId, exampleUserSwipe.UserId)
	assert.Equal(t, insertedUserSwipe.UserSwipeId, exampleUserSwipe.UserSwipeId)
	assert.Equal(t, insertedUserSwipe.Liked, exampleUserSwipe.Liked)

}
