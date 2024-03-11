package user_test

import (
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserValidate(t *testing.T) {
	id := user_id.Wrap(uuid.New())
	validUser := user.User{
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Images:    []user.UserImage{{Url: "https://example.com/image.png", UserID: id}},
	}
	assert.Nil(t, validUser.Validate())
	assert.Nil(t, (&validUser).BeforeSave(nil))

	timesCheck := validUser
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))

	firstNameLengthCheck := validUser
	firstNameLengthCheck.FirstName = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, firstNameLengthCheck.Validate())
			assert.NotNil(t, (&firstNameLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, firstNameLengthCheck.Validate())
			assert.Nil(t, (&firstNameLengthCheck).BeforeSave(nil))
		}
		firstNameLengthCheck.FirstName = firstNameLengthCheck.FirstName + "a"
	}

	lastNameLengthCheck := validUser
	lastNameLengthCheck.LastName = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, lastNameLengthCheck.Validate())
			assert.NotNil(t, (&lastNameLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, lastNameLengthCheck.Validate())
			assert.Nil(t, (&lastNameLengthCheck).BeforeSave(nil))
		}
		lastNameLengthCheck.LastName = lastNameLengthCheck.LastName + "a"
	}

	legalAgeCheck := validUser
	legalAgeCheck.Age = 0
	for i := 0; i <= 21; i++ {
		if i < 18 {
			assert.NotNil(t, legalAgeCheck.Validate())
			assert.NotNil(t, (&legalAgeCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, legalAgeCheck.Validate())
			assert.Nil(t, (&legalAgeCheck).BeforeSave(nil))
		}
		legalAgeCheck.Age = legalAgeCheck.Age + 1
	}
}

func TestUserBeforeCreate(t *testing.T) {
	noIdUser := user.User{
		ID:        user_id.UserID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
	}
	require.Nil(t, (&noIdUser).BeforeCreate(nil))
	assert.NotEmpty(t, noIdUser.ID)
	id := noIdUser.ID

	require.Nil(t, (&noIdUser).BeforeCreate(nil))
	assert.Equal(t, id, noIdUser.ID)
}
