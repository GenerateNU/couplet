package user_test

import (
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	validUser := user.User{
		ID:        user_id.Wrap(uuid.New()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
	}
	assert.Nil(t, validUser.Validate())

	idCheck := validUser
	idCheck.ID = user_id.UserID{}
	assert.NotNil(t, idCheck.Validate())

	timesCheck := validUser
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())

	firstNameLengthCheck := validUser
	firstNameLengthCheck.FirstName = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, firstNameLengthCheck.Validate())
		} else {
			assert.Nil(t, firstNameLengthCheck.Validate())
		}
		firstNameLengthCheck.FirstName = firstNameLengthCheck.FirstName + "a"
	}

	lastNameLengthCheck := validUser
	lastNameLengthCheck.LastName = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, lastNameLengthCheck.Validate())
		} else {
			assert.Nil(t, lastNameLengthCheck.Validate())
		}
		lastNameLengthCheck.LastName = lastNameLengthCheck.LastName + "a"
	}

	legalAgeCheck := validUser
	legalAgeCheck.Age = 0
	for i := 0; i <= 21; i++ {
		if i < 18 {
			assert.NotNil(t, legalAgeCheck.Validate())
		} else {
			assert.Nil(t, legalAgeCheck.Validate())
		}
		legalAgeCheck.Age = legalAgeCheck.Age + 1
	}
}
