package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"time"

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

	// set up user data
	firstName := "John"
	lastName := "Smith"
	age := 20

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user, err := c.CreateUser(context.Background(), firstName, lastName, uint8(age))
	assert.Nil(t, err)

	// ensure that all fields were set properly on the User object
	assert.Equal(t, user.Age, uint8(age))
	assert.Equal(t, user.FirstName, firstName)
	assert.Equal(t, user.LastName, lastName)

	// create a second user with the same data to show that repeated POST calls always creates new users
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("newUserId"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newUser, err := c.CreateUser(context.Background(), firstName, lastName, uint8(age))
	assert.Nil(t, err)

	assert.Equal(t, newUser.Age, uint8(age))
	assert.Equal(t, newUser.FirstName, firstName)
	assert.Equal(t, newUser.LastName, lastName)

	// IMPORTANT! assert that internally, the second user is not the same as the first user
	assert.NotEqual(t, newUser.ID, user.ID)

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
	firstName := "firstName"
	lastName := "lastName"
	age := 20

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), "firstName", "lastName", age).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	_, err = c.CreateUser(context.Background(), "firstName", "lastName", 20)
	assert.Nil(t, err)

	// retrieve the user's ID
	userId := rec.Value("id").(string)

	// expect the select statement from the delete endpoint and return the user
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), "firstName", "lastName", 20))

	mock.ExpectBegin()

	// expect the delete statement and delete the user
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "users" WHERE id = $1`)).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	deletedUser, err := c.DeleteUserById(context.Background(), userId)
	assert.Nil(t, err)

	// ensure that the deleted user is returned and matches the info of the user that was created
	assert.Equal(t, deletedUser.Age, uint8(age))
	assert.Equal(t, deletedUser.FirstName, firstName)
	assert.Equal(t, deletedUser.LastName, lastName)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}
