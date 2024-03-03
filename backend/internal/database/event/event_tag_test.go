package event_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventTagValidate(t *testing.T) {
	validEventTag := event.EventTag{
		ID:        "tag",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Events:    []event.Event{{ID: event_id.Wrap(uuid.New())}, {ID: event_id.Wrap(uuid.New())}},
	}
	assert.Nil(t, validEventTag.Validate())

	idLengthCheck := validEventTag
	idLengthCheck.ID = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, idLengthCheck.Validate())
		} else {
			assert.Nil(t, idLengthCheck.Validate())
		}
		idLengthCheck.ID = idLengthCheck.ID + "a"
	}

	timesCheck := validEventTag
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
}
