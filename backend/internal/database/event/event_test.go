package event_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"couplet/internal/util"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventValidate(t *testing.T) {
	id := event_id.Wrap(uuid.New())
	validEvent := event.Event{
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
		OrgID:     org_id.Wrap(uuid.New()),
	}
	assert.Nil(t, validEvent.Validate())
	assert.Nil(t, (&validEvent).BeforeSave(nil))

	timesCheck := validEvent
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))

	nameLengthCheck := validEvent
	nameLengthCheck.Name = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, nameLengthCheck.Validate())
			assert.NotNil(t, (&nameLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, nameLengthCheck.Validate())
			assert.Nil(t, (&nameLengthCheck).BeforeSave(nil))
		}
		nameLengthCheck.Name = nameLengthCheck.Name + "a"
	}

	bioLengthCheck := validEvent
	bioLengthCheck.Bio = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, bioLengthCheck.Validate())
			assert.NotNil(t, (&bioLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, bioLengthCheck.Validate())
			assert.Nil(t, (&bioLengthCheck).BeforeSave(nil))
		}
		bioLengthCheck.Bio = bioLengthCheck.Bio + "a"
	}

	imageCountCheck := validEvent
	imageCountCheck.Images = url_slice.UrlSlice{}
	assert.NotNil(t, imageCountCheck.Validate())
	assert.NotNil(t, (&imageCountCheck).BeforeSave(nil))

	eventTagsCheck := validEvent
	eventTagsCheck.EventTags = []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}}
	assert.NotNil(t, eventTagsCheck.Validate())
	assert.NotNil(t, (&eventTagsCheck).BeforeSave(nil))

	orgIdCheck := validEvent
	orgIdCheck.OrgID = org_id.OrgID{}
	assert.NotNil(t, orgIdCheck.Validate())
	assert.NotNil(t, (&orgIdCheck).BeforeSave(nil))
}

func TestEventBeforeCreate(t *testing.T) {
	noIdEvent := event.Event{
		ID:        event_id.EventID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{},
		OrgID:     org_id.Wrap(uuid.New()),
	}
	require.Nil(t, (&noIdEvent).BeforeCreate(nil))
	assert.NotEmpty(t, noIdEvent.ID)
	id := noIdEvent.ID

	require.Nil(t, (&noIdEvent).BeforeCreate(nil))
	assert.Equal(t, id, noIdEvent.ID)
}
