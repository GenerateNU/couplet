package handler_test

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/handler"
	"log"
	"log/slog"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const firstName, lastName = "Johnny", "Appleseed"
const age uint8 = 20

/* Set up mock database and handler. */
func setup(t *testing.T) (sqlmock.Sqlmock, api.Handler) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)
	h := handler.NewHandler(c, slog.Default())
	require.NotEmpty(t, h)
	return mock, h
}

func createDefaultUser(mock sqlmock.Sqlmock) (*api.UsersPostReq, dbtesting.ValueRecorder) {
	urlString := "https://example.com/image.png"
	imageUrl, _ := url.Parse(urlString)
	var images = []url.URL{*imageUrl}

	rec := dbtesting.NewValueRecorder()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), urlString, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	user := &api.UsersPostReq{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Images:    images,
	}
	return user, rec
}

func TestUsersGet(t *testing.T) {
	mock, h := setup(t)

	// setup user request and database expectations
	user, rec := createDefaultUser(mock)
	createdUserRes, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	createdUser, ok := createdUserRes.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")
	require.NotNil(t, createdUser)

	// create a second user
	user2, rec2 := createDefaultUser(mock)
	createdUserRes2, err2 := h.UsersPost(context.Background(), user2)
	require.Nil(t, err2)

	createdUser2, ok := createdUserRes2.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")
	require.NotNil(t, createdUser2)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" LIMIT 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(rec.Value("id"), rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age).
			AddRow(rec2.Value("id"), rec2.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))

	var limit api.OptUint8 = api.OptUint8{Value: 10, Set: true}
	var offset api.OptUint32 = api.OptUint32{Value: 0, Set: true}

	users, err := h.UsersGet(context.Background(), api.UsersGetParams{Limit: limit, Offset: offset})
	require.NotNil(t, users)
	require.Nil(t, err)

	require.Len(t, users, 2)

	idStr := rec.Value("id").(string)
	id, err := uuid.Parse(idStr)
	require.Nil(t, err)

	require.Equal(t, users[0].ID, id)
	require.Equal(t, users[0].FirstName, firstName)
	require.Equal(t, users[0].LastName, lastName)
	require.Equal(t, users[0].Age, age)

	idStr = rec2.Value("id").(string)
	id, err = uuid.Parse(idStr)
	require.Nil(t, err)

	require.Equal(t, users[1].ID, id)
	require.Equal(t, users[1].FirstName, firstName)
	require.Equal(t, users[1].LastName, lastName)
	require.Equal(t, users[1].Age, age)
}

func TestUsersIDGet(t *testing.T) {
	mock, h := setup(t)

	// setup user request and database expectations
	user, rec := createDefaultUser(mock)

	// use the handler to create the user
	createdUserRes, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	createdUser, ok := createdUserRes.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")

	// retrieve the user's ID
	userId := createdUser.ID.String()

	// set expectations for the retrieval
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))

	// parse the user's ID and use the handler to retrieve them
	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	userRes, err := h.UsersIDGet(context.Background(), api.UsersIDGetParams{ID: uuidUser})
	require.Nil(t, err)
	require.NotNil(t, userRes)

	retrievedUser, ok := userRes.(*api.UsersIDGetOK)
	require.True(t, ok, "Expected api.UsersIDGetRes to be of type *api.UsersIDGetOK")

	// ensure that the retrieved user matches the info of the user that was created
	require.Equal(t, retrievedUser.Age, age)
	require.Equal(t, retrievedUser.FirstName, firstName)
	require.Equal(t, retrievedUser.LastName, lastName)
}

func TestUsersPost(t *testing.T) {
	mock, h := setup(t)

	user, _ := createDefaultUser(mock)

	// set up recorder to keep track of the auto-generated userID

	res, err := h.UsersPost(context.Background(), user)

	require.NoError(t, err)

	resCreated, ok := res.(*api.UsersPostCreated)
	require.True(t, ok, "Expected UsersPostRes to be of type *api.UsersPostCreated")

	// require.Equal(t, resCreated.FirstName, firstName)
	// require.Equal(t, resCreated.LastName, lastName)
	// require.Equal(t, resCreated.Age, age)
	require.NotNil(t, resCreated.ID)
}

func TestTooYoungUsersPost(t *testing.T) {
	_, h := setup(t)

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
	mock, h := setup(t)

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
	mock, h := setup(t)

	// setup user request and database expectations
	user, rec := createDefaultUser(mock)

	// use the handler to create the user
	createdUserRes, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	createdUser, ok := createdUserRes.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")

	// retrieve the user's ID
	userId := createdUser.ID.String()

	// set expectations for the deletion
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// parse the user's ID and use the handler to delete them
	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	userRes, err := h.UsersIDDelete(context.Background(), api.UsersIDDeleteParams{ID: uuidUser})
	require.Nil(t, err)
	require.NotNil(t, userRes)

	deletedUser, ok := userRes.(*api.UsersIDDeleteOK)
	require.True(t, ok, "Expected api.UsersIDDeleteRes to be of type *api.UsersIDDeleteOK")

	// ensure that the deleted user is returned and matches the info of the user that was created
	require.Equal(t, deletedUser.Age, age)
	require.Equal(t, deletedUser.FirstName, firstName)
	require.Equal(t, deletedUser.LastName, lastName)

	// try to find the user in the database again, it should fail
	// mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}))
	mock.ExpectCommit()

	deletedUserRes, _ := h.UsersIDGet(context.Background(), api.UsersIDGetParams{ID: uuidUser})
	// Assert that the returned value is of type *api.Error
	errRes, ok := deletedUserRes.(*api.Error)
	require.True(t, ok, "Expected api.UsersIDGetRes to be of type *api.Error")

	// Assert that the error message is correct
	require.Equal(t, "record not found", errRes.Message)
}

func TestUsersInvalidIDDelete(t *testing.T) {
	mock, h := setup(t)
	// rec := dbtesting.NewValueRecorder()

	userId := uuid.New()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}))
	mock.ExpectRollback()

	// parse the user's ID and use the handler to delete them

	userRes, err := h.UsersIDDelete(context.Background(), api.UsersIDDeleteParams{ID: userId})

	// Assert that an error was returned
	require.Nil(t, err)
	require.NotNil(t, userRes)

	// Assert that the returned value is of type *api.Error
	errRes, ok := userRes.(*api.Error)
	require.True(t, ok, "Expected api.UsersIDDeleteRes to be of type *api.Error")

	// Assert that the error message is correct
	require.Equal(t, "record not found", errRes.Message)
}

func TestUsersIDPut(t *testing.T) {
	mock, h := setup(t)

	// setup user request and database expectations
	user, rec := createDefaultUser(mock)

	// use the handler to create the user
	createdUserRes, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	createdUser, ok := createdUserRes.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")

	// retrieve the user's ID
	userId := createdUser.ID.String()

	updatedFirstName, updatedLastName := "Taylor", "Swift"
	var updatedAge uint8 = 34
	urlString := "https://example.com/image.png"
	imageUrl, _ := url.Parse(urlString)
	var images = []url.URL{*imageUrl}

	// set expectations for the update
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE "users" SET "age"=$1,"first_name"=$2,"last_name"=$3,"updated_at"=$4 WHERE "id" = $5`)).
		WithArgs(updatedAge, updatedFirstName, updatedLastName, sqlmock.AnyArg(), userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), urlString, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectCommit()

	// parse the user's ID and use the handler to update them
	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	userRes, err := h.UsersIDPut(context.Background(), &api.UsersIDPutReq{FirstName: updatedFirstName, LastName: updatedLastName, Age: updatedAge, Images: images}, api.UsersIDPutParams{ID: uuidUser})
	require.Nil(t, err)

	updatedUser, ok := userRes.(*api.UsersIDPutOK)
	require.True(t, ok, "Expected api.UsersIDPutRes to be of type *api.UsersIDPutOK")

	// ensure that the updated user is returned and matches the info of the user that was created
	require.Equal(t, updatedUser.Age, updatedAge)
	require.Equal(t, updatedUser.FirstName, updatedFirstName)
	require.Equal(t, updatedUser.LastName, updatedLastName)
}

func TestUsersPutNew(t *testing.T) {
	mock, h := setup(t)

	updatedFirstName, updatedLastName := "Taylor", "Swift"
	var updatedAge uint8 = 34
	urlString := "https://example.com/image.png"
	imageUrl, _ := url.Parse(urlString)
	var images = []url.URL{*imageUrl}
	userId := "00000000-0000-0000-0000-000000000000"
	uuid, _ := uuid.Parse(userId)

	// set expectations for the update
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}))
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), updatedFirstName, updatedLastName, updatedAge).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), urlString, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	userRes, err := h.UsersIDPut(context.Background(), &api.UsersIDPutReq{FirstName: updatedFirstName, LastName: updatedLastName, Age: updatedAge, Images: images}, api.UsersIDPutParams{ID: uuid})
	require.Nil(t, err)
	require.NotNil(t, userRes)

	updatedUser, ok := userRes.(*api.UsersIDPutCreated)
	require.True(t, ok, "Expected api.UsersIDPutRes to be of type *api.UsersIDPutCreated")

	// ensure that the updated user is returned and matches the info of the user that was created
	require.Equal(t, updatedUser.Age, updatedAge)
	require.Equal(t, updatedUser.FirstName, updatedFirstName)
	require.Equal(t, updatedUser.LastName, updatedLastName)
}

func TestUsersIDPatch(t *testing.T) {
	mock, h := setup(t)

	// setup user request and database expectations
	user, rec := createDefaultUser(mock)

	// use the handler to create the user
	createdUserRes, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	createdUser, ok := createdUserRes.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")

	// retrieve the user's ID
	userId := createdUser.ID.String()

	updatedFirstName, updatedLastName := "Taylor", "Swift"
	var updatedAge uint8 = 34
	urlString := "https://example.com/image.png"
	imageUrl, _ := url.Parse(urlString)
	var images = []url.URL{*imageUrl}

	// set expectations for the update
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE "users" SET "updated_at"=$1,"first_name"=$2,"last_name"=$3,"age"=$4 WHERE "id" = $5`)).
		WithArgs(sqlmock.AnyArg(), updatedFirstName, updatedLastName, updatedAge, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), urlString, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// parse the user's ID and use the handler to update them
	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	userPatch := &api.User{
		FirstName: api.OptString{Value: updatedFirstName, Set: true},
		LastName:  api.OptString{Value: updatedLastName, Set: true},
		Age:       api.OptUint8{Value: updatedAge, Set: true},
		Images:    images,
	}
	userRes, err := h.UsersIDPatch(context.Background(), userPatch, api.UsersIDPatchParams{ID: uuidUser})
	require.Nil(t, err)
	log.Printf("res: %v", userRes)

	updatedUser, ok := userRes.(*api.UsersIDPatchOK)
	require.True(t, ok, "Expected api.UsersIDPatchRes to be of type *api.UsersIDPatchOK")

	// ensure that the updated user is returned and matches the info of the user that was created
	require.Equal(t, updatedUser.Age, updatedAge)
	require.Equal(t, updatedUser.FirstName, updatedFirstName)
	require.Equal(t, updatedUser.LastName, updatedLastName)
}

func TestUsersIDBadPatch(t *testing.T) {
	mock, h := setup(t)

	// setup user request and database expectations
	user, rec := createDefaultUser(mock)

	// use the handler to create the user
	createdUserRes, err := h.UsersPost(context.Background(), user)
	require.Nil(t, err)

	createdUser, ok := createdUserRes.(*api.UsersPostCreated)
	require.True(t, ok, "expected createdUserRes to be of type *api.UsersPostCreated")

	// retrieve the user's ID
	userId := createdUser.ID.String()

	updatedFirstName, updatedLastName := "Taylor", "Swift"
	var updatedAge uint8 = 16 // TOO YOUNG
	urlString := "https://example.com/image.png"
	imageUrl, _ := url.Parse(urlString)
	var images = []url.URL{*imageUrl}

	// set expectations for the update
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), firstName, lastName, age))
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE "users" SET "updated_at"=$1,"first_name"=$2,"last_name"=$3,"age"=$4 WHERE "id" = $5`)).
		WithArgs(sqlmock.AnyArg(), updatedFirstName, updatedLastName, updatedAge, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), urlString, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// parse the user's ID and use the handler to update them
	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	userPatch := &api.User{
		FirstName: api.OptString{Value: updatedFirstName, Set: true},
		LastName:  api.OptString{Value: updatedLastName, Set: true},
		Age:       api.OptUint8{Value: updatedAge, Set: true},
		Images:    images,
	}
	userRes, err := h.UsersIDPatch(context.Background(), userPatch, api.UsersIDPatchParams{ID: uuidUser})
	require.Nil(t, err)

	errRes, ok := userRes.(*api.UsersIDPatchBadRequest)
	require.True(t, ok, "Expected api.UsersIDPatchRes to be of type *api.UsersIDPatchBadRequest")

	// Assert that the error message is correct
	require.Equal(t, "Key: 'User.Age' Error:Field validation for 'Age' failed on the 'min' tag", errRes.Message)
}

func TestUsersBadIDPatch(t *testing.T) {
	mock, h := setup(t)

	// setup user request and database expectations
	updatedFirstName, updatedLastName := "Taylor", "Swift"
	var updatedAge uint8 = 34
	urlString := "https://example.com/image.png"
	imageUrl, _ := url.Parse(urlString)
	var images = []url.URL{*imageUrl}

	userId := "00000000-0000-0000-0000-000000000000"
	uuid, _ := uuid.Parse(userId)

	userPatch := &api.User{
		FirstName: api.OptString{Value: updatedFirstName, Set: true},
		LastName:  api.OptString{Value: updatedLastName, Set: true},
		Age:       api.OptUint8{Value: updatedAge, Set: true},
		Images:    images,
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}))

	userRes, _ := h.UsersIDPatch(context.Background(), userPatch, api.UsersIDPatchParams{ID: uuid})
	require.NotNil(t, userRes)

	errRes, ok := userRes.(*api.UsersIDPatchNotFound)
	require.True(t, ok, "Expected api.UsersIDPatchRes to be of type *api.UsersIDPatchNotFound")

	// Assert that the error message is correct
	require.Equal(t, "record not found", errRes.Message)
}
