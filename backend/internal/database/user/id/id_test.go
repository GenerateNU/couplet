package id_test

import (
	"couplet/internal/database/user/id"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func FuzzUserIDParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		var userId id.UserID

		// Test Scan()
		uuidErr := uuid.Scan(src)
		userIdErr := userId.Scan(src)
		assert.Equal(t, uuidErr, userIdErr)
		assert.Equal(t, uuid, userId.UUID())
		assert.Equal(t, id.New(uuid), userId)

		if uuidErr != nil && userIdErr != nil {
			// If successful scan, test Value()
			expected, expectedErr := uuid.Value()
			actual, actualErr := uuid.Value()
			assert.Nil(t, expectedErr)
			assert.Nil(t, actualErr)
			assert.Equal(t, expectedErr, actualErr)
			assert.Equal(t, expected, actual)

			// If successful scan, test String()
			assert.Equal(t, uuid.String(), userId.String())
		}
	})
}
