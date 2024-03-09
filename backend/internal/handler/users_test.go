package handler_test

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/handler"
	"log/slog"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUsersPost(t *testing.T) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)
	h := handler.NewHandler(c, slog.Default())
	require.NotEmpty(t, h)

	firstName, lastName := "First", "Last"
	var age uint8 = 20

	// set up recorder to keep track of the auto-generated userID
	rec := dbtesting.NewValueRecorder()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := &api.UsersPostReq{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	res, err := h.UsersPost(context.Background(), user)

	require.NoError(t, err)

	resCreated, ok := res.(*api.UsersPostCreated)
	require.True(t, ok, "Expected UsersPostRes to be of type *api.UsersPostCreated")

	require.Equal(t, resCreated.FirstName, firstName)
	require.Equal(t, resCreated.LastName, lastName)
	require.Equal(t, resCreated.Age, age)
	require.NotNil(t, resCreated.ID)
}

func TestTooYoungUsersPost(t *testing.T) {
	db, _ := database.NewMockDB()
	c, _ := controller.NewController(db, nil)
	h := handler.NewHandler(c, slog.Default())
	require.NotEmpty(t, h)

	firstName, lastName := "First", "Last"
	var age uint8 = 16 // too young!

	user := &api.UsersPostReq{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	res, err := h.UsersPost(context.Background(), user)

	require.Error(t, err)
	require.ErrorContains(t, err, "must be at least 18 years old")
	require.Nil(t, res)
}

func TestInvalidUsersPost(t *testing.T) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)
	h := handler.NewHandler(c, slog.Default())
	require.NotEmpty(t, h)

	firstName := "John"
	var age uint8 = 21

	// should fail when a required field like LastName is not provided
	user := &api.UsersPostReq{
		FirstName: firstName,
		Age:       age,
	}

	// set database expectations. note the rollback at the end instead of a commit
	rec := dbtesting.NewValueRecorder()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, "", age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	res, err := h.UsersPost(context.Background(), user)

	// this will error because LastName is required
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to create user")
	require.Nil(t, res)
}

func TestUsersIDDelete(t *testing.T) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)
	h := handler.NewHandler(c, slog.Default())
	require.NotEmpty(t, h)

	firstName, lastName := "Johnny", "Appleseed"
	var age uint8 = 21

	// set up recorder to keep track of the auto-generated userID
	rec := dbtesting.NewValueRecorder()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := &api.UsersPostReq{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	_, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	// retrieve the user's ID
	userId := rec.Value("id").(string)

	// set expectations for the deletion.
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))

	mock.ExpectBegin()

	// expect the delete statement and delete the user
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	userRes, err := h.UsersIDDelete(context.Background(), api.UsersIDDeleteParams{ID: uuidUser})
	require.Nil(t, err)

	deletedUser, ok := userRes.(*api.UsersIDDeleteOK)
	require.True(t, ok, "Expected api.UsersIDDeleteRes to be of type *api.UsersIDDeleteOK")

	// ensure that the deleted user is returned and matches the info of the user that was created
	require.Equal(t, deletedUser.Age, age)
	require.Equal(t, deletedUser.FirstName, firstName)
	require.Equal(t, deletedUser.LastName, lastName)
}

// func TestUsersInvalidIDDelete(t *testing.T) {

// }
