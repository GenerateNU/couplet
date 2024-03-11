package event_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventImageValidate(t *testing.T) {
	validEventImage := event.EventImage{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Url:       "http://example.com/image.jpg",
		EventID:   event_id.Wrap(uuid.New()),
	}
	assert.Nil(t, validEventImage.Validate())
	assert.Nil(t, (&validEventImage).BeforeSave(nil))

	timesCheck := validEventImage
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))

	urlCheck := validEventImage
	urlCheck.Url = ""
	assert.NotNil(t, urlCheck.Validate())
	assert.NotNil(t, (&urlCheck).BeforeSave(nil))

	ownerIdCheck := validEventImage
	ownerIdCheck.EventID = event_id.EventID{}
	assert.NotNil(t, ownerIdCheck.Validate())
	assert.NotNil(t, (&ownerIdCheck).BeforeSave(nil))
}
