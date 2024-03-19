package user_test

import (
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserSwipeValidate(t *testing.T) {
	validUserSwipe := user.UserSwipe{
		UserID:      user_id.Wrap(uuid.New()),
		OtherUserID: user_id.Wrap(uuid.New()),
		Liked:       true,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	assert.Nil(t, validUserSwipe.Validate())
	assert.Nil(t, (&validUserSwipe).BeforeSave(nil))

	userIDCheck := validUserSwipe
	userIDCheck.UserID = user_id.UserID{}
	assert.NotNil(t, userIDCheck.Validate())
	assert.NotNil(t, (&userIDCheck).BeforeSave(nil))

	otherUserIDCheck := validUserSwipe
	otherUserIDCheck.OtherUserID = user_id.UserID{}
	assert.NotNil(t, otherUserIDCheck.Validate())
	assert.NotNil(t, (&otherUserIDCheck).BeforeSave(nil))

	timesCheck := validUserSwipe
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))
}
