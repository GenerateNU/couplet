package user_test

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventSwipeValidate(t *testing.T) {
	validEventSwipe := user.EventSwipe{
		UserID:    user_id.Wrap(uuid.New()),
		EventID:   event_id.Wrap(uuid.New()),
		Liked:     true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	assert.Nil(t, validEventSwipe.Validate())
	assert.Nil(t, (&validEventSwipe).BeforeSave(nil))

	userIDCheck := validEventSwipe
	userIDCheck.UserID = user_id.UserID{}
	assert.NotNil(t, userIDCheck.Validate())
	assert.NotNil(t, (&userIDCheck).BeforeSave(nil))

	eventIDCheck := validEventSwipe
	eventIDCheck.EventID = event_id.EventID{}
	assert.NotNil(t, eventIDCheck.Validate())
	assert.NotNil(t, (&eventIDCheck).BeforeSave(nil))

	timesCheck := validEventSwipe
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))
}
