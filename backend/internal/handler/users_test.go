package handler_test

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/url_slice"
	"couplet/internal/handler"
	"couplet/internal/util"
	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersPost(t *testing.T) {
	validTestCases := []struct {
		input api.User
	}{{api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input api.User
	}{{api.User{
		FirstName: "",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.User{
		FirstName: "First",
		LastName:  "",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       5,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.User{
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
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.FirstName, v.input.LastName, v.input.Age, v.input.Bio, url_slice.Wrap(v.input.Images)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := h.UsersPost(context.Background(), &v.input)
		assert.NotEmpty(t, res)
		assert.IsType(t, &api.User{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.UsersPost(context.Background(), &v.input)
		assert.IsType(t, &api.Error{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUsersIDDelete(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	invalidId := uuid.New()
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1 RETURNING *`)).WithArgs(invalidId)
	mock.ExpectRollback()
	deleteRes, err := h.UsersIDDelete(context.Background(), api.UsersIDDeleteParams{ID: invalidId})
	assert.IsType(t, &api.Error{}, deleteRes)
	assert.Nil(t, err)

	validUser := api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, url_slice.Wrap(validUser.Images)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	createRes, err := h.UsersPost(context.Background(), &validUser)
	require.IsType(t, &api.User{}, createRes)
	require.Nil(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1 RETURNING *`)).
		WithArgs(createRes.(*api.User).ID).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "age", "bio", "images"}).AddRow(createRes.(*api.User).ID, createRes.(*api.User).FirstName, createRes.(*api.User).LastName, createRes.(*api.User).Age, createRes.(*api.User).Bio, url_slice.Wrap(createRes.(*api.User).Images)))
	mock.ExpectCommit()
	deleteRes, err = h.UsersIDDelete(context.Background(), api.UsersIDDeleteParams{ID: createRes.(*api.User).ID})
	assert.IsType(t, &api.User{}, deleteRes)
	assert.Nil(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUsersIDGet(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	invalidId := uuid.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).WithArgs(invalidId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age", "bio", "images"}))
	getRes, err := h.UsersIDGet(context.Background(), api.UsersIDGetParams{ID: invalidId})
	assert.IsType(t, &api.Error{}, getRes)
	assert.Nil(t, err)

	validUser := api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, url_slice.Wrap(validUser.Images)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	createRes, err := h.UsersPost(context.Background(), &validUser)
	require.IsType(t, &api.User{}, createRes)
	require.Nil(t, err)
	created := createRes.(*api.User)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "age", "bio", "images"}).AddRow(created.ID, created.FirstName, created.LastName, created.Age, created.Bio, url_slice.Wrap(created.Images)))
	getRes, err = h.UsersIDGet(context.Background(), api.UsersIDGetParams{ID: created.ID})
	require.IsType(t, &api.User{}, getRes)
	assert.Nil(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUsersGet(t *testing.T) {
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
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	validUser := api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	for i := 0; i < 10; i++ {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, url_slice.Wrap(validUser.Images)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		createRes, err := h.UsersPost(context.Background(), &validUser)
		require.IsType(t, &api.User{}, createRes)
		require.Nil(t, err)
	}

	for _, v := range testCases {
		rows := sqlmock.NewRows([]string{"first_name", "last_name", "age", "bio", "images"})
		for i := 0; i < int(v.input.limit); i++ {
			rows.AddRow(validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, url_slice.Wrap(validUser.Images))
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
			WillReturnRows(rows)
		res, err := h.UsersGet(context.Background(), api.UsersGetParams{Limit: api.NewOptUint8(v.input.limit), Offset: api.NewOptUint32(v.input.offset)})
		assert.Nil(t, err)
		assert.Equal(t, len(res), int(v.input.limit))
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

// func TestUsersIDPut(t *testing.T) {
// 	validTestCases := []struct {
// 		input api.User
// 	}{{api.User{
// 		FirstName: "First",
// 		LastName:  "Last",
// 		Age:       21,
// 		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
// 		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
// 	}}}
// 	invalidTestCases := []struct {
// 		input api.User
// 	}{{api.User{
// 		FirstName: "",
// 		LastName:  "Last",
// 		Age:       21,
// 		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
// 		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
// 	}}, {api.User{
// 		FirstName: "First",
// 		LastName:  "",
// 		Age:       21,
// 		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
// 		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
// 	}}, {api.User{
// 		FirstName: "First",
// 		LastName:  "Last",
// 		Age:       5,
// 		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
// 		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
// 	}}, {api.User{
// 		FirstName: "First",
// 		LastName:  "Last",
// 		Age:       21,
// 		Bio:       "",
// 		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
// 	}}, {api.User{
// 		FirstName: "First",
// 		LastName:  "Last",
// 		Age:       21,
// 		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
// 		Images:    url_slice.UrlSlice{},
// 	}}}

// 	db, mock := database.NewMockDB()
// 	require.NotNil(t, db)
// 	require.NotNil(t, mock)
// 	c, err := controller.NewController(db, nil)
// 	require.NotNil(t, c)
// 	require.Nil(t, err)
// 	h := handler.NewHandler(c, nil)
// 	require.NotNil(t, h)

// 	// for _, v := range validTestCases {
// 	// 	mock.ExpectBegin()
// 	// 	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
// 	// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.FirstName, v.input.LastName, v.input.Age, v.input.Bio, url_slice.Wrap(v.input.Images)).
// 	// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	// 	mock.ExpectCommit()

// 	// 	res, err := h.UsersIDPut(context.Background(), &v.input, api.UsersIDPutParams{ID: uuid.New()})
// 	// 	assert.IsType(t, &api.UsersIDPutOK{}, res)
// 	// 	assert.Nil(t, err)
// 	// }

// 	for _, v := range invalidTestCases {
// 		res, err := h.UsersIDPut(context.Background(), &v.input, api.UsersIDPutParams{ID: uuid.New()})
// 		assert.IsType(t, &api.Error{}, res)
// 		assert.Nil(t, err)
// 	}

// 	assert.Nil(t, mock.ExpectationsWereMet())
// }

func TestUsersIDPatch(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	validUser := api.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Bio:       "Hey everyone! I can't wait to go to an exciting event!",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age","bio","images") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validUser.FirstName, validUser.LastName, validUser.Age, validUser.Bio, url_slice.Wrap(validUser.Images)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	createdRes, err := h.UsersPost(context.Background(), &validUser)
	require.IsType(t, &api.User{}, createdRes)
	require.Nil(t, err)

	len256 := ""
	for i := 0; i < 256; i++ {
		len256 += "a"
	}

	// validTestCases := []struct {
	// 	input api.User
	// }{{api.User{
	// 	FirstName: api.NewOptString("First"),
	// 	LastName:  api.NewOptString("Last"),
	// 	Age:       api.NewOptUint8(21),
	// 	Bio:       api.NewOptString("Hey everyone! I can't wait to go to an exciting event!"),
	// 	Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	// }}}
	// invalidTestCases := []struct {
	// 	input api.User
	// }{{api.User{
	// 	Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png")},
	// }}, {api.User{
	// 	FirstName: api.NewOptString(len256),
	// }}, {api.User{
	// 	LastName: api.NewOptString(len256),
	// }}, {api.User{
	// 	Age: api.NewOptUint8(5),
	// }}, {api.User{
	// 	Bio: api.NewOptString(len256),
	// }}}

	// for _, v := range validTestCases {
	// 	mock.ExpectBegin()
	// 	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "users"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "age", "bio", "images"}).AddRow(createdRes.(*api.User).ID, createdRes.(*api.User).FirstName, createdRes.(*api.User).LastName, createdRes.(*api.User).Age, createdRes.(*api.User).Bio, url_slice.Wrap(createdRes.(*api.User).Images)))
	// 	res, err := h.UsersIDPatch(context.Background(), &v.input, api.UsersIDPatchParams{ID: createdRes.(*api.User).ID})
	// 	assert.IsType(t, &api.User{}, res)
	// 	assert.Nil(t, err)
	// }

	// for _, v := range invalidTestCases {
	// 	res, err := h.UsersIDPatch(context.Background(), &v.input, api.UsersIDPatchParams{ID: createdRes.(*api.User).ID})
	// 	assert.IsType(t, &api.UsersIDPatchBadRequest{}, res)
	// 	assert.Nil(t, err)
	// }

	// assert.Nil(t, mock.ExpectationsWereMet())
}
