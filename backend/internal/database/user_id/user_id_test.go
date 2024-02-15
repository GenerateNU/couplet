package user_id_test

import (
	"couplet/internal/database/event_id"
	"couplet/internal/database/user_id"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWrapAndUnwrap(t *testing.T) {
	id := uuid.UUID{}
	assert.Equal(t, id, user_id.Wrap(id).Unwrap())

	id = uuid.New()
	assert.Equal(t, id, user_id.Wrap(id).Unwrap())
}

func FuzzSqlParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		var userId user_id.UserID

		// Test Scan()
		uuidErr := uuid.Scan(src)
		userIdErr := userId.Scan(src)
		assert.Equal(t, uuidErr, userIdErr)
		assert.Equal(t, uuid, userId.Unwrap())
		assert.Equal(t, user_id.Wrap(uuid), userId)

		if uuidErr != nil && userIdErr != nil {
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
		userIdJson, userIdErr := json.Marshal(uuid)
		assert.Equal(t, uuidErr, userIdErr)
		assert.Equal(t, uuidJson, userIdJson)

		// Test Unmarshal()
		var eventId event_id.EventID
		uuidErr = json.Unmarshal(uuidJson, &uuid)
		userIdErr = json.Unmarshal(userIdJson, &eventId)
		assert.Equal(t, uuidErr, userIdErr)
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
			assert.Equal(t, uuid.String(), user_id.Wrap(uuid).String())
		}
	})
}
