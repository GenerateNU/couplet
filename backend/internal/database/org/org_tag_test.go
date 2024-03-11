package org_test

import (
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrgTagValidate(t *testing.T) {
	validOrgTag := org.OrgTag{
		ID:        "tag",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Orgs:      []org.Org{{ID: org_id.Wrap(uuid.New())}, {ID: org_id.Wrap(uuid.New())}},
	}
	assert.Nil(t, validOrgTag.Validate())
	assert.Nil(t, (&validOrgTag).BeforeSave(nil))

	idLengthCheck := validOrgTag
	idLengthCheck.ID = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, idLengthCheck.Validate())
			assert.NotNil(t, (&idLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, idLengthCheck.Validate())
			assert.Nil(t, (&idLengthCheck).BeforeSave(nil))
		}
		idLengthCheck.ID = idLengthCheck.ID + "a"
	}

	timesCheck := validOrgTag
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))
}
