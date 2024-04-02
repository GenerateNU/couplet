package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
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

func TestCreateEvent(t *testing.T) {
	validTestCases := []struct {
		input event.Event
	}{{event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}}
	invalidTestCases := []struct {
		input event.Event
	}{{event.Event{
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:     org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "The Events Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{},
		OrgID:  org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		EventTags: []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}},
		OrgID:     org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.OrgID{},
	}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, v.input.Images, v.input.OrgID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e, valErr, txErr := c.CreateEvent(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.CreateEvent(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteEvent(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	invalidId := event_id.Wrap(uuid.New())
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "events" WHERE "events"."id" = $1 RETURNING *`)).WithArgs(invalidId)
	mock.ExpectRollback()
	_, txErr := c.DeleteEvent(invalidId)
	assert.NotNil(t, txErr)

	validEvent := event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, validEvent.Images, validEvent.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	created, valErr, txErr := c.CreateEvent(validEvent)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "events" WHERE "events"."id" = $1 RETURNING *`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "org_id"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.Name, created.Bio, created.Images, created.OrgID))
	mock.ExpectCommit()
	deleted, txErr := c.DeleteEvent(created.ID)
	assert.Nil(t, txErr)
	assert.Equal(t, created, deleted)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetEvent(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	invalidId := event_id.Wrap(uuid.New())
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE "events"."id" = $1 ORDER BY "events"."id" LIMIT 1`)).WithArgs(invalidId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "event_tags", "org_id"}))
	_, txErr := c.GetEvent(invalidId)
	assert.NotNil(t, txErr)

	validEvent := event.Event{
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		EventTags: []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}},
		OrgID:     org_id.Wrap(uuid.New()),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, validEvent.Images, validEvent.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "event_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validEvent.EventTags[0].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[1].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[2].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[3].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[4].ID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events2tags" ("event_id","event_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validEvent.EventTags[0].ID, sqlmock.AnyArg(), validEvent.EventTags[1].ID, sqlmock.AnyArg(), validEvent.EventTags[2].ID, sqlmock.AnyArg(), validEvent.EventTags[3].ID, sqlmock.AnyArg(), validEvent.EventTags[4].ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	created, valErr, txErr := c.CreateEvent(validEvent)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events" WHERE "events"."id" = $1 ORDER BY "events"."id" LIMIT 1`)).
		WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "org_id"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.Name, created.Bio, created.Images, created.OrgID))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events2tags" WHERE "events2tags"."event_id" = $1`)).WithArgs(created.ID).WillReturnRows(sqlmock.NewRows([]string{"event_id", "event_tag_id", "created_at", "updated_at"}).
		AddRow(created.ID, created.EventTags[0].ID, created.EventTags[0].CreatedAt, created.EventTags[0].UpdatedAt).
		AddRow(created.ID, created.EventTags[1].ID, created.EventTags[1].CreatedAt, created.EventTags[1].UpdatedAt).
		AddRow(created.ID, created.EventTags[2].ID, created.EventTags[2].CreatedAt, created.EventTags[2].UpdatedAt).
		AddRow(created.ID, created.EventTags[3].ID, created.EventTags[3].CreatedAt, created.EventTags[3].UpdatedAt).
		AddRow(created.ID, created.EventTags[4].ID, created.EventTags[4].CreatedAt, created.EventTags[4].UpdatedAt))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "event_tags" WHERE "event_tags"."id" IN ($1,$2,$3,$4,$5)`)).WithArgs(created.EventTags[0].ID, created.EventTags[1].ID, created.EventTags[2].ID, created.EventTags[3].ID, created.EventTags[4].ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(created.EventTags[0].ID, created.EventTags[0].CreatedAt, created.EventTags[0].UpdatedAt).
		AddRow(created.EventTags[1].ID, created.EventTags[1].CreatedAt, created.EventTags[1].UpdatedAt).
		AddRow(created.EventTags[2].ID, created.EventTags[2].CreatedAt, created.EventTags[2].UpdatedAt).
		AddRow(created.EventTags[3].ID, created.EventTags[3].CreatedAt, created.EventTags[3].UpdatedAt).
		AddRow(created.EventTags[4].ID, created.EventTags[4].CreatedAt, created.EventTags[4].UpdatedAt))
	retrieved, txErr := c.GetEvent(created.ID)
	assert.Nil(t, txErr)
	assert.Equal(t, created, retrieved)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetEvents(t *testing.T) {
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

	validEvent := event.Event{
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		EventTags: []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}},
		OrgID:     org_id.Wrap(uuid.New()),
	}

	for i := 0; i < 10; i++ {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, validEvent.Images, validEvent.OrgID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "event_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
			WithArgs(validEvent.EventTags[0].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[1].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[2].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[3].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[4].ID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events2tags" ("event_id","event_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
			WithArgs(sqlmock.AnyArg(), validEvent.EventTags[0].ID, sqlmock.AnyArg(), validEvent.EventTags[1].ID, sqlmock.AnyArg(), validEvent.EventTags[2].ID, sqlmock.AnyArg(), validEvent.EventTags[3].ID, sqlmock.AnyArg(), validEvent.EventTags[4].ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		created, valErr, txErr := c.CreateEvent(validEvent)
		require.NotEmpty(t, created)
		require.Nil(t, valErr)
		require.Nil(t, txErr)
	}

	for _, v := range testCases {
		rows := sqlmock.NewRows([]string{"name", "bio", "images", "org_id"})
		for i := 0; i < int(v.input.limit); i++ {
			rows.AddRow(validEvent.Name, validEvent.Bio, validEvent.Images, validEvent.OrgID)
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "events"`)).
			WillReturnRows(rows)
		e, txErr := c.GetEvents(v.input.limit, v.input.offset)
		assert.Nil(t, txErr)
		assert.Equal(t, len(e), int(v.input.limit))
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSaveEvent(t *testing.T) {
	validTestCases := []struct {
		input event.Event
	}{{event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}}
	invalidTestCases := []struct {
		input event.Event
	}{{event.Event{
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:     org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "The Events Company",
		Bio:    "",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{},
		OrgID:  org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		EventTags: []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}},
		OrgID:     org_id.Wrap(uuid.New()),
	}}, {event.Event{
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.OrgID{},
	}}}

	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), v.input.Name, v.input.Bio, v.input.Images, v.input.OrgID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e, valErr, txErr := c.SaveEvent(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.SaveEvent(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateEvent(t *testing.T) {
	db, mock := database.NewMockDB()
	require.NotNil(t, db)
	require.NotNil(t, mock)
	c, err := controller.NewController(db, nil)
	require.NotNil(t, c)
	require.Nil(t, err)

	validEvent := event.Event{
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		EventTags: []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}},
		OrgID:     org_id.Wrap(uuid.New()),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events" ("id","created_at","updated_at","name","bio","images","org_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.Name, validEvent.Bio, validEvent.Images, validEvent.OrgID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "event_tags" ("id","created_at","updated_at") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9),($10,$11,$12),($13,$14,$15) ON CONFLICT DO NOTHING`)).
		WithArgs(validEvent.EventTags[0].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[1].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[2].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[3].ID, sqlmock.AnyArg(), sqlmock.AnyArg(), validEvent.EventTags[4].ID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "events2tags" ("event_id","event_tag_id") VALUES ($1,$2),($3,$4),($5,$6),($7,$8),($9,$10) ON CONFLICT DO NOTHING`)).
		WithArgs(sqlmock.AnyArg(), validEvent.EventTags[0].ID, sqlmock.AnyArg(), validEvent.EventTags[1].ID, sqlmock.AnyArg(), validEvent.EventTags[2].ID, sqlmock.AnyArg(), validEvent.EventTags[3].ID, sqlmock.AnyArg(), validEvent.EventTags[4].ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	created, valErr, txErr := c.CreateEvent(validEvent)
	require.NotEmpty(t, created)
	require.Nil(t, valErr)
	require.Nil(t, txErr)

	len256 := ""
	for i := 0; i < 256; i++ {
		len256 += "a"
	}

	validTestCases := []struct {
		input event.Event
	}{{event.Event{
		ID:     created.ID,
		Name:   "The Events Company",
		Bio:    "At The Events Company, we connect people through events",
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:  org_id.Wrap(uuid.New()),
	}}}
	invalidTestCases := []struct {
		input event.Event
	}{{event.Event{
		ID:        created.ID,
		CreatedAt: time.Now().Add(10 * time.Hour),
		UpdatedAt: time.Now(),
	}}, {event.Event{
		ID:     created.ID,
		Images: url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png")},
	}}, {event.Event{
		ID:   created.ID,
		Name: len256,
	}}, {event.Event{
		ID:  created.ID,
		Bio: len256,
	}}}

	for _, v := range validTestCases {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "events"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "bio", "images", "org_id"}).AddRow(created.ID, created.CreatedAt, created.UpdatedAt, created.Name, created.Bio, created.Images, created.OrgID))
		e, valErr, txErr := c.UpdateEvent(v.input)
		assert.NotEmpty(t, e)
		assert.Nil(t, valErr)
		assert.Nil(t, txErr)
	}

	for _, v := range invalidTestCases {
		_, valErr, txErr := c.UpdateEvent(v.input)
		assert.NotNil(t, valErr)
		assert.Nil(t, txErr)
	}

	assert.Nil(t, mock.ExpectationsWereMet())
}
