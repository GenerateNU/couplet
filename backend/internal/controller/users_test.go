package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"couplet/internal/util"
	"regexp"
	"time"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	validTestCases := []struct {
		input user.User
	}{{user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input user.User
	}{{user.User{
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       5,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{},
	}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.FirstName, v.input.LastName, v.input.Age, v.input.Bio, v.input.Images).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e, valErr, txErr := c.CreateUser(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.CreateUser(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	invalidId := user_id.Wrap(uuid.New())
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1 RETURNING *`)).WithArgs(invalidId)
	mock.ExpectRollback()
	_, txErr := c.DeleteUser(invalidId)
	assert.NotNil(t, txErr)

	validUser := user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, validUser.Images).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	created, valErr, txErr := c.CreateUser(validUser)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1 RETURNING *`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age", "bio", "images"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.FirstName, created.LastName, created.Age, created.Bio, created.Images))
	mock.ExpectCommit()
	deleted, txErr := c.DeleteUser(created.ID)
	assert.Nil(t, txErr)
	assert.Equal(t, created, deleted)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetUser(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	invalidId := user_id.Wrap(uuid.New())
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).WithArgs(invalidId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age", "bio", "images"}))
	_, txErr := c.GetUser(invalidId)
	assert.NotNil(t, txErr)

	validUser := user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, validUser.Images).
		WillReturnResult(sqlmock.NewResult(1, 1))
	created, valErr, txErr := c.CreateUser(validUser)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age", "bio", "images"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.FirstName, created.LastName, created.Age, created.Bio, created.Images))
	retrieved, txErr := c.GetUser(created.ID)
	assert.Nil(t, txErr)
	assert.Equal(t, created, retrieved)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetUsers(t *testing.T) {
	testCases := []struct {
		input struct {
			limit  uint8
			offset uint32
		}
	}{{input: struct {
		limit  uint8
		offset uint32
	}{limit: 0, offset: 1}},
		{input: struct {
			limit  uint8
			offset uint32
		}{limit: 10, offset: 13}},
		{input: struct {
			limit  uint8
			offset uint32
		}{limit: 4, offset: 340}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	validUser := user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	for i := 0; i < 10; i++ {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, validUser.Images).
			WillReturnResult(sqlmock.NewResult(1, 1))
		created, valErr, txErr := c.CreateUser(validUser)
		require.NotEmpty(t, created)
		require.Nil(t, valErr)
		require.Nil(t, txErr)
	}

	for _, v := range testCases {
		rows := sqlmock.NewRows([]string{"first_name", "last_name", "age", "bio", "images"})
		for i := 0; i < int(v.input.limit); i++ {
			rows.AddRow(validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, validUser.Images)
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
			WillReturnRows(rows)
		e, txErr := c.GetUsers(v.input.limit, v.input.offset)
		assert.Nil(t, txErr)
		assert.Equal(t, len(e), int(v.input.limit))
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSaveUser(t *testing.T) {
	validTestCases := []struct {
		input user.User
	}{{user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input user.User
	}{{user.User{
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       5,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{},
	}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.FirstName, v.input.LastName, v.input.Age, v.input.Bio, v.input.Images).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e, valErr, txErr := c.SaveUser(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.SaveUser(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	validUser := user.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, validUser.Images).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	created, valErr, txErr := c.CreateUser(validUser)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	len256 := ""
	for i := 0; i < 256; i++ {
		len256 += "a"
	}

	validTestCases := []struct {
		input user.User
	}{{user.User{
		ID:        created.ID,
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input user.User
	}{{user.User{
		ID:        created.ID,
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
	}}, {user.User{
		ID:     created.ID,
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png")},
	}}, {user.User{
		ID:        created.ID,
		FirstName: len256,
	}}, {user.User{
		ID:       created.ID,
		LastName: len256,
	}}, {user.User{
		ID:  created.ID,
		Age: 5,
	}}, {user.User{
		ID:  created.ID,
		Bio: len256,
	}}}

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "users"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age", "bio", "images"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.FirstName, created.LastName, created.Age, created.Bio, created.Images))
		e, valErr, txErr := c.UpdateUser(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.UpdateUser(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}
