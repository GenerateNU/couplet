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
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID
	rec := dbtesting.NewValueRecorder()

	// set up exampleEvent data
	exampleEvent := api.Event{
		Name:    "Big event",
		Bio:     "Event description",
		Address: "123 Something St, Boston",
	}

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "events" ("id","created_at","updated_at","name","bio","address","organization_id")
		VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(rec.Record("idOne"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEvent.Name, exampleEvent.Bio, exampleEvent.Address, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventOne, err := c.CreateEvent(context.Background(), &exampleEvent)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the User object
	assert.Equal(t, insertedEventOne.Name, exampleEvent.Name)
	assert.Equal(t, insertedEventOne.Bio, exampleEvent.Bio)
	assert.Equal(t, insertedEventOne.Address, exampleEvent.Address)

	// create a second user with the same data to show that repeated POST calls always creates new users
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(rec.Record("idTwo"), sqlmock.AnyArg(), sqlmock.AnyArg(), exampleEvent.Name, exampleEvent.Bio, exampleEvent.Address, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedEventTwo, err := c.CreateEvent(context.Background(), &exampleEvent)
	assert.Nil(t, err)

	assert.Equal(t, insertedEventTwo.Name, exampleEvent.Name)
	assert.Equal(t, insertedEventTwo.Bio, exampleEvent.Bio)
	assert.Equal(t, insertedEventTwo.Address, exampleEvent.Address)

	// IMPORTANT! assert that internally, the second user is not the same as the first user
	assert.NotEqual(t, insertedEventTwo.ID, insertedEventOne.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}

// func TestDeleteUser(t *testing.T) {
// 	// set up mock database
// 	db, mock := database.NewMockDB()
// 	c, err := controller.NewController(db)
// 	assert.NotEmpty(t, c)
// 	assert.Nil(t, err)

// 	// set up recorder to keep track of the auto-generated userID and created/updated times
// 	rec := dbtesting.NewValueRecorder()

// 	// set up user data
// 	firstName := "firstName"
// 	lastName := "lastName"
// 	age := 20

// 	// expect the insert statement and create the user
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
// 		VALUES ($1,$2,$3,$4,$5,$6)`)).
// 		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), "firstName", "lastName", age).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectCommit()

// 	_, err = c.CreateUser(context.Background(), "firstName", "lastName", 20)
// 	assert.Nil(t, err)

// 	// retrieve the user's ID
// 	userId := rec.Value("id").(string)

// 	// expect the select statement from the delete endpoint and return the user
// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
// 		WithArgs(userId).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
// 			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), "firstName", "lastName", 20))

// 	mock.ExpectBegin()

// 	// expect the delete statement and delete the user
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		DELETE FROM "users" WHERE id = $1`)).
// 		WithArgs(userId).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	deletedUser, err := c.DeleteUserById(context.Background(), userId)
// 	assert.Nil(t, err)

// 	// ensure that the deleted user is returned and matches the info of the user that was created
// 	assert.Equal(t, deletedUser.Age, uint8(age))
// 	assert.Equal(t, deletedUser.FirstName, firstName)
// 	assert.Equal(t, deletedUser.LastName, lastName)

// 	// ensure that all expectations are met in the mock
// 	errExpectations := mock.ExpectationsWereMet()
// 	assert.Nil(t, errExpectations)
// }
