package id_test

import (
	"couplet/internal/database/event/id"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func FuzzEventIDParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		var eventId id.EventID

		// Test Scan()
		uuidErr := uuid.Scan(src)
		eventIdErr := eventId.Scan(src)
		assert.Equal(t, uuidErr, eventIdErr)
		assert.Equal(t, uuid, eventId.UUID())
		assert.Equal(t, id.New(uuid), eventId)

		if uuidErr != nil && eventIdErr != nil {
			// If successful scan, test Value()
			expected, expectedErr := uuid.Value()
			actual, actualErr := uuid.Value()
			assert.Nil(t, expectedErr)
			assert.Nil(t, actualErr)
			assert.Equal(t, expectedErr, actualErr)
			assert.Equal(t, expected, actual)
		}
	})
}
