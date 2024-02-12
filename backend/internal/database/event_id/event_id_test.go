package event_id_test

import (
	"couplet/internal/database/event_id"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func FuzzSqlParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		var eventId event_id.EventID

		// Test Scan()
		uuidErr := uuid.Scan(src)
		eventIdErr := eventId.Scan(src)
		assert.Equal(t, uuidErr, eventIdErr)
		assert.Equal(t, uuid, eventId.Unwrap())
		assert.Equal(t, event_id.Wrap(uuid), eventId)

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

func FuzzJsonParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		if uuid.Scan(src) != nil {
			return
		}

		// Test Marshal()
		uuidJson, uuidErr := json.Marshal(uuid)
		eventIdJson, eventIdErr := json.Marshal(uuid)
		assert.Equal(t, uuidErr, eventIdErr)
		assert.Equal(t, uuidJson, eventIdJson)

		// Test Unmarshal()
		var eventId event_id.EventID
		uuidErr = json.Unmarshal(uuidJson, &uuid)
		eventIdErr = json.Unmarshal(eventIdJson, &eventId)
		assert.Equal(t, uuidErr, eventIdErr)
		assert.Equal(t, uuid, eventId.Unwrap())
		assert.Equal(t, event_id.Wrap(uuid), eventId)
	})
}

func FuzzStringParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID

		uuidErr := uuid.Scan(src)
		if uuidErr != nil {
			assert.Equal(t, uuid.String(), event_id.Wrap(uuid).String())
		}
	})
}
