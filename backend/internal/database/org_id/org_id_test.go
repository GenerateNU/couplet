package org_id_test

import (
	"couplet/internal/database/org_id"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWrapAndUnwrap(t *testing.T) {
	id := uuid.UUID{}
	assert.Equal(t, id, org_id.Wrap(id).Unwrap())

	id = uuid.New()
	assert.Equal(t, id, org_id.Wrap(id).Unwrap())
}

func FuzzSqlParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		var orgId org_id.OrgID

		// Test Scan()
		uuidErr := uuid.Scan(src)
		orgIdErr := orgId.Scan(src)
		assert.Equal(t, uuidErr, orgIdErr)
		assert.Equal(t, uuid, orgId.Unwrap())
		assert.Equal(t, org_id.Wrap(uuid), orgId)

		if uuidErr != nil && orgIdErr != nil {
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
		orgIdJson, orgIdErr := json.Marshal(uuid)
		assert.Equal(t, uuidErr, orgIdErr)
		assert.Equal(t, uuidJson, orgIdJson)

		// Test Unmarshal()
		var orgId org_id.OrgID
		uuidErr = json.Unmarshal(uuidJson, &uuid)
		orgIdErr = json.Unmarshal(orgIdJson, &orgId)
		assert.Equal(t, uuidErr, orgIdErr)
		assert.Equal(t, uuid, orgId.Unwrap())
		assert.Equal(t, org_id.Wrap(uuid), orgId)
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
			assert.Equal(t, uuid.String(), org_id.Wrap(uuid).String())
		}
	})
}
