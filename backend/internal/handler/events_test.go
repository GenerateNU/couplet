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

func TestEventsPost(t *testing.T) {
	validTestCases := []struct {
		input api.EventsPostReq
	}{{api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}}}
	invalidTestCases := []struct {
		input api.EventsPostReq
	}{{api.EventsPostReq{
		Name:   "",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}}, {api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}}, {api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{},
		OrgId:  uuid.New(),
	}}, {api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6"},
		OrgId:  uuid.New(),
	}}, {api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.UUID{},
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
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, url_slice.Wrap(v.input.Images), v.input.OrgId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := h.EventsPost(context.Background(), &v.input)
		assert.NotEmpty(t, res)
		assert.IsType(t, &api.EventsPostCreated{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.EventsPost(context.Background(), &v.input)
		assert.IsType(t, &api.Error{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestEventsIDDelete(t *testing.T) {
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
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "events" WHERE "events"."id" = $1 RETURNING *`)).WithArgs(invalidId)
	mock.ExpectRollback()
	deleteRes, err := h.EventsIDDelete(context.Background(), api.EventsIDDeleteParams{ID: invalidId})
	assert.IsType(t, &api.Error{}, deleteRes)
	assert.Nil(t, err)

	validEvent := api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, url_slice.Wrap(validEvent.Images), validEvent.OrgId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	createRes, err := h.EventsPost(context.Background(), &validEvent)
	require.IsType(t, &api.EventsPostCreated{}, createRes)
	require.Nil(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "events" WHERE "events"."id" = $1 RETURNING *`)).
		WithArgs(createRes.(*api.EventsPostCreated).ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "images", "org_id"}).AddRow(createRes.(*api.EventsPostCreated).ID, createRes.(*api.EventsPostCreated).Name, createRes.(*api.EventsPostCreated).Bio, url_slice.Wrap(createRes.(*api.EventsPostCreated).Images), createRes.(*api.EventsPostCreated).OrgId))
	mock.ExpectCommit()
	deleteRes, err = h.EventsIDDelete(context.Background(), api.EventsIDDeleteParams{ID: createRes.(*api.EventsPostCreated).ID})
	assert.IsType(t, &api.EventsIDDeleteOK{}, deleteRes)
	assert.Nil(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestEventsIDGet(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	invalidId := uuid.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE "events"."id" = $1 ORDER BY "events"."id" LIMIT 1`)).WithArgs(invalidId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "event_tags", "org_id"}))
	getRes, err := h.EventsIDGet(context.Background(), api.EventsIDGetParams{ID: invalidId})
	assert.IsType(t, &api.Error{}, getRes)
	assert.Nil(t, err)

	validEvent := api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
		OrgId:  uuid.New(),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, url_slice.Wrap(validEvent.Images), validEvent.OrgId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "event_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validEvent.Tags[0], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[1], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[2], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[3], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[4], sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events2tags" ("event_id","event_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validEvent.Tags[0], sqlmock.AnyArg(), validEvent.Tags[1], sqlmock.AnyArg(), validEvent.Tags[2], sqlmock.AnyArg(), validEvent.Tags[3], sqlmock.AnyArg(), validEvent.Tags[4]).
		WillReturnResult(sqlmock.NewResult(1, 1))
	createRes, err := h.EventsPost(context.Background(), &validEvent)
	require.IsType(t, &api.EventsPostCreated{}, createRes)
	require.Nil(t, err)
	created := createRes.(*api.EventsPostCreated)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE "events"."id" = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "images", "org_id"}).AddRow(created.ID, created.Name, created.Bio, url_slice.Wrap(created.Images), created.OrgId))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events2tags" WHERE "events2tags"."event_id" = $1`)).WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"event_id", "event_tag_id"}).
		AddRow(created.ID, created.Tags[0]).
		AddRow(created.ID, created.Tags[1]).
		AddRow(created.ID, created.Tags[2]).
		AddRow(created.ID, created.Tags[3]).
		AddRow(created.ID, created.Tags[4]))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "event_tags" WHERE "event_tags"."id" IN ($1,$2,$3,$4,$5)`)).WithArgs(created.Tags[0], created.Tags[1], created.Tags[2], created.Tags[3], created.Tags[4]).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(created.Tags[0]).
		AddRow(created.Tags[1]).
		AddRow(created.Tags[2]).
		AddRow(created.Tags[3]).
		AddRow(created.Tags[4]))
	getRes, err = h.EventsIDGet(context.Background(), api.EventsIDGetParams{ID: created.ID})
	require.IsType(t, &api.EventsIDGetOK{}, getRes)
	assert.Nil(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestEventsGet(t *testing.T) {
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

	validEvent := api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
		OrgId:  uuid.New(),
	}

	for i := 0; i < 10; i++ {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, url_slice.Wrap(validEvent.Images), validEvent.OrgId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "event_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
			WithArgs(validEvent.Tags[0], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[1], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[2], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[3], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[4], sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events2tags" ("event_id","event_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
			WithArgs(sqlmock.AnyArg(), validEvent.Tags[0], sqlmock.AnyArg(), validEvent.Tags[1], sqlmock.AnyArg(), validEvent.Tags[2], sqlmock.AnyArg(), validEvent.Tags[3], sqlmock.AnyArg(), validEvent.Tags[4]).
			WillReturnResult(sqlmock.NewResult(1, 1))
		createRes, err := h.EventsPost(context.Background(), &validEvent)
		require.IsType(t, &api.EventsPostCreated{}, createRes)
		require.Nil(t, err)
	}

	for _, v := range testCases {
		rows := sqlmock.NewRows([]string{"name", "bio", "images", "org_id"})
		for i := 0; i < int(v.input.limit); i++ {
			rows.AddRow(validEvent.Name, validEvent.Bio, url_slice.Wrap(validEvent.Images), validEvent.OrgId)
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events"`)).
			WillReturnRows(rows)
		res, err := h.EventsGet(context.Background(), api.EventsGetParams{Limit: api.NewOptUint8(v.input.limit), Offset: api.NewOptUint32(v.input.offset)})
		assert.Nil(t, err)
		assert.Equal(t, len(res), int(v.input.limit))
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestEventsIDPut(t *testing.T) {
	validTestCases := []struct {
		input api.EventsIDPutReq
	}{{api.EventsIDPutReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}}}
	invalidTestCases := []struct {
		input api.EventsIDPutReq
	}{{api.EventsIDPutReq{
		Name:   "",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}}, {api.EventsIDPutReq{
		Name:   "The Events Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.New(),
	}}, {api.EventsIDPutReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{},
		OrgId:  uuid.New(),
	}}, {api.EventsIDPutReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6"},
		OrgId:  uuid.New(),
	}}, {api.EventsIDPutReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  uuid.UUID{},
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
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, url_slice.Wrap(v.input.Images), v.input.OrgId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := h.EventsIDPut(context.Background(), &v.input, api.EventsIDPutParams{ID: uuid.New()})
		assert.IsType(t, &api.EventsIDPutOK{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.EventsIDPut(context.Background(), &v.input, api.EventsIDPutParams{ID: uuid.New()})
		assert.IsType(t, &api.Error{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestEventsIDPatch(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	validEvent := api.EventsPostReq{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
		OrgId:  uuid.New(),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, url_slice.Wrap(validEvent.Images), validEvent.OrgId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "event_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validEvent.Tags[0], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[1], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[2], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[3], sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Tags[4], sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events2tags" ("event_id","event_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validEvent.Tags[0], sqlmock.AnyArg(), validEvent.Tags[1], sqlmock.AnyArg(), validEvent.Tags[2], sqlmock.AnyArg(), validEvent.Tags[3], sqlmock.AnyArg(), validEvent.Tags[4]).
		WillReturnResult(sqlmock.NewResult(1, 1))
	createdRes, err := h.EventsPost(context.Background(), &validEvent)
	require.IsType(t, &api.EventsPostCreated{}, createdRes)
	require.Nil(t, err)

	len256 := ""
	for i := 0; i < 256; i++ {
		len256 += "a"
	}

	validTestCases := []struct {
		input api.Event
	}{{api.Event{
		Name:   api.NewOptString("The Events Company"),
		Bio:    api.NewOptString("At The Events Company, we connect people through events"),
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgId:  api.NewOptUUID(uuid.New()),
	}}}
	invalidTestCases := []struct {
		input api.Event
	}{{api.Event{
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png")},
	}}, {api.Event{
		Name: api.NewOptString(len256),
	}}, {api.Event{
		Bio: api.NewOptString(len256),
	}}}

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "events"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "images", "org_id"}).AddRow(createdRes.(*api.EventsPostCreated).ID, createdRes.(*api.EventsPostCreated).Name, createdRes.(*api.EventsPostCreated).Bio, url_slice.Wrap(createdRes.(*api.EventsPostCreated).Images), createdRes.(*api.EventsPostCreated).OrgId))
		res, err := h.EventsIDPatch(context.Background(), &v.input, api.EventsIDPatchParams{ID: createdRes.(*api.EventsPostCreated).ID})
		assert.IsType(t, &api.EventsIDPatchOK{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.EventsIDPatch(context.Background(), &v.input, api.EventsIDPatchParams{ID: createdRes.(*api.EventsPostCreated).ID})
		assert.IsType(t, &api.EventsIDPatchBadRequest{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}
