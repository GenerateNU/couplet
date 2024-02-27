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

	userIDCheck := validUserSwipe
	userIDCheck.UserID = user_id.UserID{}
	assert.NotNil(t, userIDCheck.Validate())

	otherUserIDCheck := validUserSwipe
	otherUserIDCheck.OtherUserID = user_id.UserID{}
	assert.NotNil(t, otherUserIDCheck.Validate())

	timesCheck := validUserSwipe
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
}
