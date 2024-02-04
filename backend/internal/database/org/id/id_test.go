package id_test

import (
	"couplet/internal/database/org/id"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func FuzzOrgIDParity(f *testing.F) {
	f.Add("5e91507e-5630-4efd-9fd4-799178870b10")
	f.Add("f47ac10b-58cc-0372-8567-0e02b2c3d4")
	f.Add("")

	f.Fuzz(func(t *testing.T, src string) {
		var uuid uuid.UUID
		var orgId id.OrgID

		// Test Scan()
		uuidErr := uuid.Scan(src)
		orgIdErr := orgId.Scan(src)
		assert.Equal(t, uuidErr, orgIdErr)
		assert.Equal(t, uuid, orgId.UUID())
		assert.Equal(t, id.New(uuid), orgId)

		if uuidErr != nil && orgIdErr != nil {
			// If successful scan, test Value()
			expected, expectedErr := uuid.Value()
			actual, actualErr := uuid.Value()
			assert.Nil(t, expectedErr)
			assert.Nil(t, actualErr)
			assert.Equal(t, expectedErr, actualErr)
			assert.Equal(t, expected, actual)

			// If successful scan, test String()
			assert.Equal(t, uuid.String(), orgId.String())
		}
	})
}
