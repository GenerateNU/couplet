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

func TestOrgsPost(t *testing.T) {
	validTestCases := []struct {
		input api.OrgsPostReq
	}{{api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input api.OrgsPostReq
	}{{api.OrgsPostReq{
		Name:   "",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{},
	}}, {api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6"}},
	}}

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
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, url_slice.Wrap(v.input.Images)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := h.OrgsPost(context.Background(), &v.input)
		assert.NotEmpty(t, res)
		assert.IsType(t, &api.OrgsPostCreated{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.OrgsPost(context.Background(), &v.input)
		assert.IsType(t, &api.Error{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestOrgsIDDelete(t *testing.T) {
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
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "orgs" WHERE "orgs"."id" = $1 RETURNING *`)).WithArgs(invalidId)
	mock.ExpectRollback()
	deleteRes, err := h.OrgsIDDelete(context.Background(), api.OrgsIDDeleteParams{ID: invalidId})
	assert.IsType(t, &api.Error{}, deleteRes)
	assert.Nil(t, err)

	validOrg := api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, url_slice.Wrap(validOrg.Images)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	createRes, err := h.OrgsPost(context.Background(), &validOrg)
	require.IsType(t, &api.OrgsPostCreated{}, createRes)
	require.Nil(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "orgs" WHERE "orgs"."id" = $1 RETURNING *`)).
		WithArgs(createRes.(*api.OrgsPostCreated).ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "images"}).AddRow(createRes.(*api.OrgsPostCreated).ID, createRes.(*api.OrgsPostCreated).Name, createRes.(*api.OrgsPostCreated).Bio, url_slice.Wrap(createRes.(*api.OrgsPostCreated).Images)))
	mock.ExpectCommit()
	deleteRes, err = h.OrgsIDDelete(context.Background(), api.OrgsIDDeleteParams{ID: createRes.(*api.OrgsPostCreated).ID})
	assert.IsType(t, &api.OrgsIDDeleteOK{}, deleteRes)
	assert.Nil(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestOrgsIDGet(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	invalidId := uuid.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs" WHERE "orgs"."id" = $1 ORDER BY "orgs"."id" LIMIT 1`)).WithArgs(invalidId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "org_tags"}))
	getRes, err := h.OrgsIDGet(context.Background(), api.OrgsIDGetParams{ID: invalidId})
	assert.IsType(t, &api.Error{}, getRes)
	assert.Nil(t, err)

	validOrg := api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, url_slice.Wrap(validOrg.Images)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "org_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validOrg.Tags[0], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[1], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[2], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[3], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[4], sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs2tags" ("org_id","org_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validOrg.Tags[0], sqlmock.AnyArg(), validOrg.Tags[1], sqlmock.AnyArg(), validOrg.Tags[2], sqlmock.AnyArg(), validOrg.Tags[3], sqlmock.AnyArg(), validOrg.Tags[4]).
		WillReturnResult(sqlmock.NewResult(1, 1))
	createRes, err := h.OrgsPost(context.Background(), &validOrg)
	require.IsType(t, &api.OrgsPostCreated{}, createRes)
	require.Nil(t, err)
	created := createRes.(*api.OrgsPostCreated)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs" WHERE "orgs"."id" = $1 ORDER BY "orgs"."id" LIMIT 1`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "images"}).AddRow(created.ID, created.Name, created.Bio, url_slice.Wrap(created.Images)))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs2tags" WHERE "orgs2tags"."org_id" = $1`)).WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"org_id", "org_tag_id"}).
		AddRow(created.ID, created.Tags[0]).
		AddRow(created.ID, created.Tags[1]).
		AddRow(created.ID, created.Tags[2]).
		AddRow(created.ID, created.Tags[3]).
		AddRow(created.ID, created.Tags[4]))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "org_tags" WHERE "org_tags"."id" IN ($1,$2,$3,$4,$5)`)).WithArgs(created.Tags[0], created.Tags[1], created.Tags[2], created.Tags[3], created.Tags[4]).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(created.Tags[0]).
		AddRow(created.Tags[1]).
		AddRow(created.Tags[2]).
		AddRow(created.Tags[3]).
		AddRow(created.Tags[4]))
	getRes, err = h.OrgsIDGet(context.Background(), api.OrgsIDGetParams{ID: created.ID})
	require.IsType(t, &api.OrgsIDGetOK{}, getRes)
	assert.Nil(t, err)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestOrgsGet(t *testing.T) {
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

	validOrg := api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
	}

	for i := 0; i < 10; i++ {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, url_slice.Wrap(validOrg.Images)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "org_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
			WithArgs(validOrg.Tags[0], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[1], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[2], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[3], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[4], sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs2tags" ("org_id","org_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
			WithArgs(sqlmock.AnyArg(), validOrg.Tags[0], sqlmock.AnyArg(), validOrg.Tags[1], sqlmock.AnyArg(), validOrg.Tags[2], sqlmock.AnyArg(), validOrg.Tags[3], sqlmock.AnyArg(), validOrg.Tags[4]).
			WillReturnResult(sqlmock.NewResult(1, 1))
		createRes, err := h.OrgsPost(context.Background(), &validOrg)
		require.IsType(t, &api.OrgsPostCreated{}, createRes)
		require.Nil(t, err)
	}

	for _, v := range testCases {
		rows := sqlmock.NewRows([]string{"name", "bio", "images"})
		for i := 0; i < int(v.input.limit); i++ {
			rows.AddRow(validOrg.Name, validOrg.Bio, url_slice.Wrap(validOrg.Images))
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs"`)).
			WillReturnRows(rows)
		res, err := h.OrgsGet(context.Background(), api.OrgsGetParams{Limit: api.NewOptUint8(v.input.limit), Offset: api.NewOptUint32(v.input.offset)})
		assert.Nil(t, err)
		assert.Equal(t, len(res), int(v.input.limit))
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestOrgsIDPut(t *testing.T) {
	validTestCases := []struct {
		input api.OrgsIDPutReq
	}{{api.OrgsIDPutReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input api.OrgsIDPutReq
	}{{api.OrgsIDPutReq{
		Name:   "",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.OrgsIDPutReq{
		Name:   "The Orgs Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {api.OrgsIDPutReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{},
	}}, {api.OrgsIDPutReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6"},
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
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, url_slice.Wrap(v.input.Images)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := h.OrgsIDPut(context.Background(), &v.input, api.OrgsIDPutParams{ID: uuid.New()})
		assert.IsType(t, &api.OrgsIDPutOK{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.OrgsIDPut(context.Background(), &v.input, api.OrgsIDPutParams{ID: uuid.New()})
		assert.IsType(t, &api.Error{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestOrgsIDPatch(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)
	h := handler.NewHandler(c, nil)
	require.NotNil(t, h)

	validOrg := api.OrgsPostReq{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5"},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, url_slice.Wrap(validOrg.Images)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "org_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validOrg.Tags[0], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[1], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[2], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[3], sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Tags[4], sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs2tags" ("org_id","org_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validOrg.Tags[0], sqlmock.AnyArg(), validOrg.Tags[1], sqlmock.AnyArg(), validOrg.Tags[2], sqlmock.AnyArg(), validOrg.Tags[3], sqlmock.AnyArg(), validOrg.Tags[4]).
		WillReturnResult(sqlmock.NewResult(1, 1))
	createdRes, err := h.OrgsPost(context.Background(), &validOrg)
	require.IsType(t, &api.OrgsPostCreated{}, createdRes)
	require.Nil(t, err)

	len256 := ""
	for i := 0; i < 256; i++ {
		len256 += "a"
	}

	validTestCases := []struct {
		input api.OrgsIDPatchReq
	}{{api.OrgsIDPatchReq{
		Name:   api.NewOptString("The Orgs Company"),
		Bio:    api.NewOptString("At The Orgs Company, we connect people through orgs"),
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input api.OrgsIDPatchReq
	}{{api.OrgsIDPatchReq{
		Name: api.NewOptString(len256),
	}}, {api.OrgsIDPatchReq{
		Bio: api.NewOptString(len256),
	}}}

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "orgs"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio", "images"}).AddRow(createdRes.(*api.OrgsPostCreated).ID, createdRes.(*api.OrgsPostCreated).Name, createdRes.(*api.OrgsPostCreated).Bio, url_slice.Wrap(createdRes.(*api.OrgsPostCreated).Images)))
		res, err := h.OrgsIDPatch(context.Background(), &v.input, api.OrgsIDPatchParams{ID: createdRes.(*api.OrgsPostCreated).ID})
		assert.IsType(t, &api.OrgsIDPatchOK{}, res)
		assert.Nil(t, err)
	}

	for _, v := range invalidTestCases {
		res, err := h.OrgsIDPatch(context.Background(), &v.input, api.OrgsIDPatchParams{ID: createdRes.(*api.OrgsPostCreated).ID})
		assert.IsType(t, &api.OrgsIDPatchBadRequest{}, res)
		assert.Nil(t, err)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}
