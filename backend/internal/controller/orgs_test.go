package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"couplet/internal/util"
	"regexp"
	"time"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrg(t *testing.T) {
	validTestCases := []struct {
		input org.Org
	}{{org.Org{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input org.Org
	}{{org.Org{
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
		Name:      "The Orgs Company",
		Bio:       "At The Orgs Company, we connect people through orgs",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {org.Org{
		Name:   "",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {org.Org{
		Name:   "The Orgs Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {org.Org{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{},
	}}, {org.Org{
		Name:    "The Orgs Company",
		Bio:     "At The Orgs Company, we connect people through orgs",
		Images:  url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgTags: []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}},
	}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, v.input.Images).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e, valErr, txErr := c.CreateOrg(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.CreateOrg(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteOrg(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	invalidId := org_id.Wrap(uuid.New())
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "orgs" WHERE "orgs"."id" = $1 RETURNING *`)).WithArgs(invalidId)
	mock.ExpectRollback()
	_, txErr := c.DeleteOrg(invalidId)
	assert.NotNil(t, txErr)

	validOrg := org.Org{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, validOrg.Images).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	created, valErr, txErr := c.CreateOrg(validOrg)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "orgs" WHERE "orgs"."id" = $1 RETURNING *`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.Name, created.Bio, created.Images))
	mock.ExpectCommit()
	deleted, txErr := c.DeleteOrg(created.ID)
	assert.Nil(t, txErr)
	assert.Equal(t, created, deleted)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetOrg(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	invalidId := org_id.Wrap(uuid.New())
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs" WHERE "orgs"."id" = $1 ORDER BY "orgs"."id" LIMIT 1`)).WithArgs(invalidId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "org_tags"}))
	_, txErr := c.GetOrg(invalidId)
	assert.NotNil(t, txErr)

	validOrg := org.Org{
		Name:    "The Orgs Company",
		Bio:     "At The Orgs Company, we connect people through orgs",
		Images:  url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgTags: []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, validOrg.Images).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "org_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validOrg.OrgTags[0].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[1].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[2].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[3].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[4].ID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs2tags" ("org_id","org_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validOrg.OrgTags[0].ID, sqlmock.AnyArg(), validOrg.OrgTags[1].ID, sqlmock.AnyArg(), validOrg.OrgTags[2].ID, sqlmock.AnyArg(), validOrg.OrgTags[3].ID, sqlmock.AnyArg(), validOrg.OrgTags[4].ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	created, valErr, txErr := c.CreateOrg(validOrg)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs" WHERE "orgs"."id" = $1 ORDER BY "orgs"."id" LIMIT 1`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.Name, created.Bio, created.Images))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs2tags" WHERE "orgs2tags"."org_id" = $1`)).WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"org_id", "org_tag_id", "created_at", "updated_at"}).
		AddRow(created.ID, created.OrgTags[0].ID, created.OrgTags[0].CreatedAt, created.OrgTags[0].UpdatedAt).
		AddRow(created.ID, created.OrgTags[1].ID, created.OrgTags[1].CreatedAt, created.OrgTags[1].UpdatedAt).
		AddRow(created.ID, created.OrgTags[2].ID, created.OrgTags[2].CreatedAt, created.OrgTags[2].UpdatedAt).
		AddRow(created.ID, created.OrgTags[3].ID, created.OrgTags[3].CreatedAt, created.OrgTags[3].UpdatedAt).
		AddRow(created.ID, created.OrgTags[4].ID, created.OrgTags[4].CreatedAt, created.OrgTags[4].UpdatedAt))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "org_tags" WHERE "org_tags"."id" IN ($1,$2,$3,$4,$5)`)).WithArgs(created.OrgTags[0].ID, created.OrgTags[1].ID, created.OrgTags[2].ID, created.OrgTags[3].ID, created.OrgTags[4].ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(created.OrgTags[0].ID, created.OrgTags[0].CreatedAt, created.OrgTags[0].UpdatedAt).
		AddRow(created.OrgTags[1].ID, created.OrgTags[1].CreatedAt, created.OrgTags[1].UpdatedAt).
		AddRow(created.OrgTags[2].ID, created.OrgTags[2].CreatedAt, created.OrgTags[2].UpdatedAt).
		AddRow(created.OrgTags[3].ID, created.OrgTags[3].CreatedAt, created.OrgTags[3].UpdatedAt).
		AddRow(created.OrgTags[4].ID, created.OrgTags[4].CreatedAt, created.OrgTags[4].UpdatedAt))
	retrieved, txErr := c.GetOrg(created.ID)
	assert.Nil(t, txErr)
	assert.Equal(t, created, retrieved)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetOrgs(t *testing.T) {
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

	validOrg := org.Org{
		Name:    "The Orgs Company",
		Bio:     "At The Orgs Company, we connect people through orgs",
		Images:  url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgTags: []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}},
	}

	for i := 0; i < 10; i++ {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, validOrg.Images).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "org_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
			WithArgs(validOrg.OrgTags[0].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[1].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[2].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[3].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[4].ID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs2tags" ("org_id","org_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
			WithArgs(sqlmock.AnyArg(), validOrg.OrgTags[0].ID, sqlmock.AnyArg(), validOrg.OrgTags[1].ID, sqlmock.AnyArg(), validOrg.OrgTags[2].ID, sqlmock.AnyArg(), validOrg.OrgTags[3].ID, sqlmock.AnyArg(), validOrg.OrgTags[4].ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		created, valErr, txErr := c.CreateOrg(validOrg)
		require.NotEmpty(t, created)
		require.Nil(t, valErr)
		require.Nil(t, txErr)
	}

	for _, v := range testCases {
		rows := sqlmock.NewRows([]string{"name", "bio", "images"})
		for i := 0; i < int(v.input.limit); i++ {
			rows.AddRow(validOrg.Name, validOrg.Bio, validOrg.Images)
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orgs"`)).
			WillReturnRows(rows)
		e, txErr := c.GetOrgs(v.input.limit, v.input.offset)
		assert.Nil(t, txErr)
		assert.Equal(t, len(e), int(v.input.limit))
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSaveOrg(t *testing.T) {
	validTestCases := []struct {
		input org.Org
	}{{org.Org{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input org.Org
	}{{org.Org{
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
		Name:      "The Orgs Company",
		Bio:       "At The Orgs Company, we connect people through orgs",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {org.Org{
		Name:   "",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {org.Org{
		Name:   "The Orgs Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}, {org.Org{
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{},
	}}, {org.Org{
		Name:    "The Orgs Company",
		Bio:     "At The Orgs Company, we connect people through orgs",
		Images:  url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgTags: []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}},
	}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, v.input.Images).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e, valErr, txErr := c.SaveOrg(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.SaveOrg(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateOrg(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	validOrg := org.Org{
		Name:    "The Orgs Company",
		Bio:     "At The Orgs Company, we connect people through orgs",
		Images:  url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgTags: []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}},
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs" ("id","created_at","updated_at","name","bio","images") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.Name, validOrg.Bio, validOrg.Images).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "org_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validOrg.OrgTags[0].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[1].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[2].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[3].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validOrg.OrgTags[4].ID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "orgs2tags" ("org_id","org_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validOrg.OrgTags[0].ID, sqlmock.AnyArg(), validOrg.OrgTags[1].ID, sqlmock.AnyArg(), validOrg.OrgTags[2].ID, sqlmock.AnyArg(), validOrg.OrgTags[3].ID, sqlmock.AnyArg(), validOrg.OrgTags[4].ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	created, valErr, txErr := c.CreateOrg(validOrg)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	len256 := ""
	for i := 0; i < 256; i++ {
		len256 += "a"
	}

	validTestCases := []struct {
		input org.Org
	}{{org.Org{
		ID:     created.ID,
		Name:   "The Orgs Company",
		Bio:    "At The Orgs Company, we connect people through orgs",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}}}
	invalidTestCases := []struct {
		input org.Org
	}{{org.Org{
		ID:        created.ID,
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
	}}, {org.Org{
		ID:   created.ID,
		Name: len256,
	}}, {org.Org{
		ID:  created.ID,
		Bio: len256,
	}}}

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "orgs"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.Name, created.Bio, created.Images))
		e, valErr, txErr := c.UpdateOrg(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.UpdateOrg(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}
