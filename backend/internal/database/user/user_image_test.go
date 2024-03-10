package user_test

import (
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserImageValidate(t *testing.T) {
	validUserImage := user.UserImage{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Url:       "http://example.com/image.jpg",
		UserID:    user_id.Wrap(uuid.New()),
	}
	assert.Nil(t, validUserImage.Validate())
	assert.Nil(t, (&validUserImage).BeforeSave(nil))

	timesCheck := validUserImage
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))

	urlCheck := validUserImage
	urlCheck.Url = ""
	assert.NotNil(t, urlCheck.Validate())
	assert.NotNil(t, (&urlCheck).BeforeSave(nil))

	ownerIdCheck := validUserImage
	ownerIdCheck.UserID = user_id.UserID{}
	assert.NotNil(t, ownerIdCheck.Validate())
	assert.NotNil(t, (&ownerIdCheck).BeforeSave(nil))
}
