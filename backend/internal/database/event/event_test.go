package event_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventValidate(t *testing.T) {
	validEvent := event.Event{
		ID:        event_id.Wrap(uuid.New()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		OrgID:     org_id.Wrap(uuid.New()),
	}
	assert.Nil(t, validEvent.Validate())

	timesCheck := validEvent
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())

	nameLengthCheck := validEvent
	nameLengthCheck.Name = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, nameLengthCheck.Validate())
		} else {
			assert.Nil(t, nameLengthCheck.Validate())
		}
		nameLengthCheck.Name = nameLengthCheck.Name + "a"
	}

	bioLengthCheck := validEvent
	bioLengthCheck.Bio = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, bioLengthCheck.Validate())
		} else {
			assert.Nil(t, bioLengthCheck.Validate())
		}
		bioLengthCheck.Bio = bioLengthCheck.Bio + "a"
	}

	eventTagsCheck := validEvent
	eventTagsCheck.EventTags = []event.EventTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}}
	assert.NotNil(t, eventTagsCheck.Validate())

	orgIdCheck := validEvent
	orgIdCheck.OrgID = org_id.OrgID{}
	assert.NotNil(t, orgIdCheck.Validate())
}
