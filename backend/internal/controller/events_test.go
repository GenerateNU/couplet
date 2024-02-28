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
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

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
	require.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	require.Equal(t, insertedEventOne.Name, exampleEventOne.Name)
	require.Equal(t, insertedEventOne.Bio, exampleEventOne.Bio)

	// create a second event with the same data to show that repeated POST calls always creates new events
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
	INSERT INTO "events" ("id","created_at","updated_at","name","bio","org_id")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("idTwo"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEventTwo.Name, exampleEventTwo.Bio, exampleEventTwo.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventTwo, err := c.CreateEvent(exampleEventTwo)
	require.Nil(t, err)

	require.Equal(t, insertedEventTwo.Name, exampleEventTwo.Name)
	require.Equal(t, insertedEventTwo.Bio, exampleEventTwo.Bio)

	// IMPORTANT! require that internally, the second event id is not the same as the first event id
	require.NotEqual(t, insertedEventTwo.ID, insertedEventOne.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}

func TestDeleteEvent(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

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
	require.Nil(t, err)

	// ensure that all fields were set properly on the Event object
	require.Equal(t, createdEvent.Name, exampleEventOne.Name)
	require.Equal(t, createdEvent.Bio, exampleEventOne.Bio)

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

	// deletedEvent, err := c.DeleteEvent(createdEvent.ID)
	require.Nil(t, err)

	// ensure the deleted event value returned correctly
	// require.Equal(t, deletedEvent, createdEvent)

	// // ensure that all expectations are met in the mock
	// errExpectations := mock.ExpectationsWereMet()
	// require.Nil(t, errExpectations)

	// // deleting an event that doesn't exist should return an error
	// badId := uuid.New()

	// mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE id = $1 ORDER BY "events"."id" LIMIT 1`)).
	// 	WithArgs(exampleEventOne.ID).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
	// 		AddRow(exampleEventOne.ID, exampleEventOne.Name, exampleEventOne.Bio, exampleEventOne.OrgID))

	// _, err = c.DeleteEvent(event_id.Wrap(badId))
	// require.Error(t, err)

	// // ensure that all expectations are met in the mock
	// errExpectations = mock.ExpectationsWereMet()

	// require.Nil(t, errExpectations)
}

func TestGetEvent(t *testing.T) {
	// Set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.Nil(t, err)

	// Mocked event data
	orgId := org_id.Wrap(uuid.New())
	mockEvent := event.Event{
		ID:    event_id.Wrap(uuid.New()),
		Name:  "Sample Event",
		Bio:   "A brief description",
		OrgID: orgId,
	}

	// Set expectation: select query to fetch the event by ID
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE id = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(mockEvent.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEvent.ID, mockEvent.Name, mockEvent.Bio, mockEvent.OrgID))

	// Call GetEvent
	retrievedEvent, err := c.GetEvent(mockEvent.ID)
	require.Nil(t, err)

	// require that the retrieved event matches the mocked event
	require.Equal(t, mockEvent.ID, retrievedEvent.ID)
	require.Equal(t, mockEvent.Name, retrievedEvent.Name)
	require.Equal(t, mockEvent.Bio, retrievedEvent.Bio)
	require.Equal(t, mockEvent.OrgID, retrievedEvent.OrgID)

	// Ensure all expectations were met
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}

func TestGetEvents(t *testing.T) {
	// Set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.Nil(t, err)

	// Mocked event data
	orgId := org_id.Wrap(uuid.New())
	mockEventOne := event.Event{
		ID:    event_id.Wrap(uuid.New()),
		Name:  "Sample Event",
		Bio:   "A brief description",
		OrgID: orgId,
	}
	mockEventTwo := event.Event{
		ID:    event_id.Wrap(uuid.New()),
		Name:  "Another Event",
		Bio:   "A different description",
		OrgID: orgId,
	}
	mockEventThree := event.Event{
		ID:    event_id.Wrap(uuid.New()),
		Name:  "Third Event",
		Bio:   "A third description",
		OrgID: orgId,
	}
	mockEventFour := event.Event{
		ID:    event_id.Wrap(uuid.New()),
		Name:  "Fourth Event",
		Bio:   "A fourth description",
		OrgID: orgId,
	}
	mockEventFive := event.Event{
		ID:    event_id.Wrap(uuid.New()),
		Name:  "Fifth Event",
		Bio:   "A fifth description",
		OrgID: orgId,
	}

	// Set expectation: select query to fetch all events
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEventOne.ID, mockEventOne.Name, mockEventOne.Bio, mockEventOne.OrgID).
			AddRow(mockEventTwo.ID, mockEventTwo.Name, mockEventTwo.Bio, mockEventTwo.OrgID).
			AddRow(mockEventThree.ID, mockEventThree.Name, mockEventThree.Bio, mockEventThree.OrgID).
			AddRow(mockEventFour.ID, mockEventFour.Name, mockEventFour.Bio, mockEventFour.OrgID).
			AddRow(mockEventFive.ID, mockEventFive.Name, mockEventFive.Bio, mockEventFive.OrgID))

	// Call GetEvents
	retrievedAllEvents1, err := c.GetEvents(10, 0)
	require.Nil(t, err)

	// require that all of the retrieved events match the mocked events
	require.Equal(t, 5, len(retrievedAllEvents1))
	require.Equal(t, mockEventOne.ID, retrievedAllEvents1[0].ID)
	require.Equal(t, mockEventOne.Name, retrievedAllEvents1[0].Name)
	require.Equal(t, mockEventOne.Bio, retrievedAllEvents1[0].Bio)
	require.Equal(t, mockEventOne.OrgID, retrievedAllEvents1[0].OrgID)
	require.Equal(t, mockEventTwo.ID, retrievedAllEvents1[1].ID)
	require.Equal(t, mockEventTwo.Name, retrievedAllEvents1[1].Name)
	require.Equal(t, mockEventTwo.Bio, retrievedAllEvents1[1].Bio)
	require.Equal(t, mockEventTwo.OrgID, retrievedAllEvents1[1].OrgID)
	require.Equal(t, mockEventThree.ID, retrievedAllEvents1[2].ID)
	require.Equal(t, mockEventThree.Name, retrievedAllEvents1[2].Name)
	require.Equal(t, mockEventThree.Bio, retrievedAllEvents1[2].Bio)
	require.Equal(t, mockEventThree.OrgID, retrievedAllEvents1[2].OrgID)
	require.Equal(t, mockEventFour.ID, retrievedAllEvents1[3].ID)
	require.Equal(t, mockEventFour.Name, retrievedAllEvents1[3].Name)
	require.Equal(t, mockEventFour.Bio, retrievedAllEvents1[3].Bio)
	require.Equal(t, mockEventFour.OrgID, retrievedAllEvents1[3].OrgID)
	require.Equal(t, mockEventFive.ID, retrievedAllEvents1[4].ID)
	require.Equal(t, mockEventFive.Name, retrievedAllEvents1[4].Name)
	require.Equal(t, mockEventFive.Bio, retrievedAllEvents1[4].Bio)
	require.Equal(t, mockEventFive.OrgID, retrievedAllEvents1[4].OrgID)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEventOne.ID, mockEventOne.Name, mockEventOne.Bio, mockEventOne.OrgID).
			AddRow(mockEventTwo.ID, mockEventTwo.Name, mockEventTwo.Bio, mockEventTwo.OrgID))

	// Call GetEvents with a limit
	retrievedAllEvents2, err := c.GetEvents(2, 0)
	require.Nil(t, err)

	require.Equal(t, 2, len(retrievedAllEvents2))
	require.Equal(t, mockEventOne.ID, retrievedAllEvents2[0].ID)
	require.Equal(t, mockEventOne.Name, retrievedAllEvents2[0].Name)
	require.Equal(t, mockEventOne.Bio, retrievedAllEvents2[0].Bio)
	require.Equal(t, mockEventOne.OrgID, retrievedAllEvents2[0].OrgID)
	require.Equal(t, mockEventTwo.ID, retrievedAllEvents2[1].ID)
	require.Equal(t, mockEventTwo.Name, retrievedAllEvents2[1].Name)
	require.Equal(t, mockEventTwo.Bio, retrievedAllEvents2[1].Bio)
	require.Equal(t, mockEventTwo.OrgID, retrievedAllEvents2[1].OrgID)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEventThree.ID, mockEventThree.Name, mockEventThree.Bio, mockEventThree.OrgID).
			AddRow(mockEventFour.ID, mockEventFour.Name, mockEventFour.Bio, mockEventFour.OrgID))

	// Call GetEvents with a limit and offset
	retrievedAllEvents3, err := c.GetEvents(2, 2)
	require.Nil(t, err)

	require.Equal(t, 2, len(retrievedAllEvents3))
	require.Equal(t, mockEventThree.ID, retrievedAllEvents3[0].ID)
	require.Equal(t, mockEventThree.Name, retrievedAllEvents3[0].Name)
	require.Equal(t, mockEventThree.Bio, retrievedAllEvents3[0].Bio)
	require.Equal(t, mockEventThree.OrgID, retrievedAllEvents3[0].OrgID)
	require.Equal(t, mockEventFour.ID, retrievedAllEvents3[1].ID)
	require.Equal(t, mockEventFour.Name, retrievedAllEvents3[1].Name)
	require.Equal(t, mockEventFour.Bio, retrievedAllEvents3[1].Bio)
	require.Equal(t, mockEventFour.OrgID, retrievedAllEvents3[1].OrgID)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEventThree.ID, mockEventThree.Name, mockEventThree.Bio, mockEventThree.OrgID).
			AddRow(mockEventFour.ID, mockEventFour.Name, mockEventFour.Bio, mockEventFour.OrgID).
			AddRow(mockEventFive.ID, mockEventFive.Name, mockEventFive.Bio, mockEventFive.OrgID))

	// Call GetEvents with an offset, but no limit
	retrievedAllEvents4, err := c.GetEvents(10, 2)
	require.Nil(t, err)

	require.Equal(t, 3, len(retrievedAllEvents4))
	require.Equal(t, mockEventThree.ID, retrievedAllEvents4[0].ID)
	require.Equal(t, mockEventThree.Name, retrievedAllEvents4[0].Name)
	require.Equal(t, mockEventThree.Bio, retrievedAllEvents4[0].Bio)
	require.Equal(t, mockEventThree.OrgID, retrievedAllEvents4[0].OrgID)
	require.Equal(t, mockEventFour.ID, retrievedAllEvents4[1].ID)
	require.Equal(t, mockEventFour.Name, retrievedAllEvents4[1].Name)
	require.Equal(t, mockEventFour.Bio, retrievedAllEvents4[1].Bio)
	require.Equal(t, mockEventFour.OrgID, retrievedAllEvents4[1].OrgID)
	require.Equal(t, mockEventFive.ID, retrievedAllEvents4[2].ID)
	require.Equal(t, mockEventFive.Name, retrievedAllEvents4[2].Name)
	require.Equal(t, mockEventFive.Bio, retrievedAllEvents4[2].Bio)
	require.Equal(t, mockEventFive.OrgID, retrievedAllEvents4[2].OrgID)

	// Ensure all expectations were met
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}

func TestPutEvent(t *testing.T) {
	// Set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// Mocked event data
	orgId := org_id.Wrap(uuid.New())
	eventId := uuid.New()
	mockEvent := event.Event{
		ID:    event_id.Wrap(eventId),
		Name:  "Sample Event",
		Bio:   "A brief description",
		OrgID: orgId,
	}

	// Set expectation: select query to get the event
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE id = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(mockEvent.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEvent.ID, mockEvent.Name, mockEvent.Bio, mockEvent.OrgID))

	// Set expectation: begin transaction
	mock.ExpectBegin()

	// Set expectation: update query to update the event
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "events" SET "created_at"=$1,"updated_at"=$2,"name"=$3,"bio"=$4,"org_id"=$5 WHERE "id" = $6`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), mockEvent.Name, mockEvent.Bio, mockEvent.OrgID, mockEvent.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Set expectation: commit transaction
	mock.ExpectCommit()

	// Call PutEvent
	updatedEvent, err := c.PutEvent(event_id.Wrap(eventId), mockEvent)
	if err != nil {
		t.Errorf("Error was not expected while updating event: %s", err)
	}

	// require that the updated event matches the mocked event data
	require.Equal(t, mockEvent.ID, updatedEvent.ID)
	require.Equal(t, mockEvent.Name, updatedEvent.Name)
	require.Equal(t, mockEvent.Bio, updatedEvent.Bio)
	require.Equal(t, mockEvent.OrgID, updatedEvent.OrgID)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestPatchEvent(t *testing.T) {
	// Set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// Mocked event data
	orgId := org_id.Wrap(uuid.New())
	eventId := uuid.New()
	mockEvent := event.Event{
		ID:    event_id.Wrap(eventId),
		Name:  "Sample Event",
		Bio:   "A brief description",
		OrgID: orgId,
	}

	// Set expectation: select query to get the event
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE id = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(mockEvent.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "org_id"}).
			AddRow(mockEvent.ID, mockEvent.Name, mockEvent.Bio, mockEvent.OrgID))

	// Set expectation: begin transaction
	mock.ExpectBegin()

	// Set expectation: update query to update the event
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "events" SET "created_at"=$1,"updated_at"=$2,"name"=$3,"bio"=$4,"org_id"=$5 WHERE "id" = $6`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), mockEvent.Name, mockEvent.Bio, mockEvent.OrgID, mockEvent.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Set expectation: commit transaction
	mock.ExpectCommit()

	// Call PatchEvent
	updatedEvent, err := c.PatchEvent(event_id.Wrap(eventId), mockEvent)
	if err != nil {
		t.Errorf("Error was not expected while updating event: %s", err)
	}

	// require that the updated event matches the mocked event data
	require.Equal(t, mockEvent.ID, updatedEvent.ID)
	require.Equal(t, mockEvent.Name, updatedEvent.Name)
	require.Equal(t, mockEvent.Bio, updatedEvent.Bio)
	require.Equal(t, mockEvent.OrgID, updatedEvent.OrgID)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
