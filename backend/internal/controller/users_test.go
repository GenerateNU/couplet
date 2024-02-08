package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"time"

	"testing"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	// logger := slog.New(pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug)))
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID
	rec := dbtesting.NewValueRecorder()

	// set up user data
	firstName := "John"
	lastName := "Smith"
	var age uint8 = 20

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user, err := c.CreateUser(firstName, lastName, age)
	require.Nil(t, err)

	// ensure that all fields were set properly on the User object
	require.Equal(t, user.Age, age)
	require.Equal(t, user.FirstName, firstName)
	require.Equal(t, user.LastName, lastName)

	// create a second user with the same data to show that repeated POST calls always creates new users
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("newUserId"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newUser, err := c.CreateUser(firstName, lastName, age)
	require.Nil(t, err)

	require.Equal(t, newUser.Age, age)
	require.Equal(t, newUser.FirstName, firstName)
	require.Equal(t, newUser.LastName, lastName)

	// IMPORTANT! assert that internally, the second user is not the same as the first user
	require.NotEqual(t, newUser.ID, user.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}

func TestDeleteUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()

	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID and created/updated times
	rec := dbtesting.NewValueRecorder()

	// set up user data
	firstName := "firstName"
	lastName := "lastName"
	var age uint8 = 20

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), "firstName", "lastName", age).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	_, err = c.CreateUser("firstName", "lastName", age)
	require.Nil(t, err)

	// retrieve the user's ID
	userId := rec.Value("id").(string)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), "firstName", "lastName", 20))

	mock.ExpectBegin()

	// expect the delete statement and delete the user
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	deletedUser, err := c.DeleteUserById(uuidUser)
	require.Nil(t, err)

	// ensure that the deleted user is returned and matches the info of the user that was created
	require.Equal(t, deletedUser.Age, age)
	require.Equal(t, deletedUser.FirstName, firstName)
	require.Equal(t, deletedUser.LastName, lastName)

	// try deleting a fake user
	badId := uuid.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(badId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"})) // no rows added

	deletedUser, err = c.DeleteUserById(badId)
	require.Error(t, err)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}
