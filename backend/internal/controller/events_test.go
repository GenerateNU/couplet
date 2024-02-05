package controller_test

import (
	"couplet/internal/api"
	"couplet/internal/controller"
	"couplet/internal/database"

	"testing"

	"context"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated eventID
	rec := dbtesting.NewValueRecorder()

	// set up example event data
	orgId := uuid.New()
	exampleEventOne := api.Event{
		Name:           "Big event",
		Bio:            "Event description",
		OrganizationID: orgId,
	}
	exampleEventTwo := api.Event{
		Name:           "Big event",
		Bio:            "Event description",
		OrganizationID: orgId,
	}

	// expect the insert statement and create the event
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "events" ("id","created_at","updated_at","name","bio","organization_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idOne"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventOne.Name, exampleEventOne.Bio, exampleEventOne.OrganizationID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventOne, err := c.CreateEvent(context.Background(), &exampleEventOne)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	assert.Equal(t, insertedEventOne.Name, exampleEventOne.Name)
	assert.Equal(t, insertedEventOne.Bio, exampleEventOne.Bio)

	// create a second user with the same data to show that repeated POST calls always creates new events
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
	INSERT INTO "events" ("id","created_at","updated_at","name","bio","organization_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idTwo"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventTwo.Name, exampleEventTwo.Bio, exampleEventTwo.OrganizationID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventTwo, err := c.CreateEvent(context.Background(), &exampleEventTwo)
	assert.Nil(t, err)

	assert.Equal(t, insertedEventTwo.Name, exampleEventTwo.Name)
	assert.Equal(t, insertedEventTwo.Bio, exampleEventTwo.Bio)

	// IMPORTANT! assert that internally, the second event id is not the same as the first event id
	assert.NotEqual(t, insertedEventTwo.ID, insertedEventOne.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}

func TestDeleteUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID and created/updated times
	rec := dbtesting.NewValueRecorder()

	// set up user data
	orgId := uuid.New()
	exampleEventOne := api.Event{
		Name:           "Big event",
		Bio:            "Event description",
		OrganizationID: orgId,
	}

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "events" ("id","created_at","updated_at","name","bio","organization_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), exampleEventOne.Name, exampleEventOne.Bio, exampleEventOne.OrganizationID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	createdEvent, err := c.CreateEvent(context.Background(), &exampleEventOne)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	assert.Equal(t, createdEvent.Name, exampleEventOne.Name)
	assert.Equal(t, createdEvent.Bio, exampleEventOne.Bio)

	// expect the delete statement and delete the user
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "events" WHERE id = $1`)).
		WithArgs(exampleEventOne.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = c.DeleteEventById(context.Background(), exampleEventOne.ID.Value)
	assert.Nil(t, err)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)

	// res := db.Find(exampleEventOne.ID.Value)

	// // ensure that the deleted user is returned and matches the info of the user that was created
	// assert.Equal(t, res.RowsAffected, int64(0))
}
