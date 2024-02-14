package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"time"

	"testing"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated eventID
	rec := dbtesting.NewValueRecorder()

	// set up example event data
	orgId := org_id.Wrap(uuid.New())
	exampleEventOne := event.Event{
		Name:  "Big event",
		Bio:   "Event description",
		OrgID: orgId,
	}
	exampleEventTwo := exampleEventOne

	// expect the insert statement and create the event
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "events" ("id","created_at","updated_at","name","bio","org_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idOne"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventOne.Name, exampleEventOne.Bio, exampleEventOne.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventOne, err := c.CreateEvent(exampleEventOne)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	assert.Equal(t, insertedEventOne.Name, exampleEventOne.Name)
	assert.Equal(t, insertedEventOne.Bio, exampleEventOne.Bio)

	// create a second event with the same data to show that repeated POST calls always creates new events
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
	INSERT INTO "events" ("id","created_at","updated_at","name","bio","org_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idTwo"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventTwo.Name, exampleEventTwo.Bio, exampleEventTwo.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventTwo, err := c.CreateEvent(exampleEventTwo)
	assert.Nil(t, err)

	assert.Equal(t, insertedEventTwo.Name, exampleEventTwo.Name)
	assert.Equal(t, insertedEventTwo.Bio, exampleEventTwo.Bio)

	// IMPORTANT! assert that internally, the second event id is not the same as the first event id
	assert.NotEqual(t, insertedEventTwo.ID, insertedEventOne.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}

func TestDeleteEvent(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated eventId and created/updated times
	rec := dbtesting.NewValueRecorder()

	// set up event data
	exampleEventOne := event.Event{
		Name:  "Big event",
		Bio:   "Event description",
		OrgID: org_id.Wrap(uuid.New()),
	}

	// expect the insert statement and create the event
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "events" ("id","created_at","updated_at","name","bio","org_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("eventId"), rec.Record("createdTime"), rec.Record("updatedTime"), exampleEventOne.Name, exampleEventOne.Bio, exampleEventOne.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	createdEvent, err := c.CreateEvent(exampleEventOne)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	assert.Equal(t, createdEvent.Name, exampleEventOne.Name)
	assert.Equal(t, createdEvent.Bio, exampleEventOne.Bio)

	// expect the initial select statement to store the event
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE id = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(rec.Value("eventId")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "org_id"}).
			AddRow(rec.Value("eventId"), rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), exampleEventOne.Name, exampleEventOne.Bio, exampleEventOne.OrgID))

	// expect the delete statement and delete the event
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "events"`)).
		WithArgs(rec.Value("eventId")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	deletedEvent, err := c.DeleteEvent(createdEvent.ID)
	assert.Nil(t, err)

	// ensure the deleted event value returned correctly
	assert.Equal(t, deletedEvent, createdEvent)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)

	// deleting an event that doesn't exist should return an error
	badId := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE id = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(badId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "org_id"})) // no rows added

	_, err = c.DeleteEvent(event_id.Wrap(badId))
	assert.Error(t, err)

	// ensure that all expectations are met in the mock
	errExpectations = mock.ExpectationsWereMet()

	assert.Nil(t, errExpectations)
}
