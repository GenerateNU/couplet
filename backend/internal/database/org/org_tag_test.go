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

	idLengthCheck := validOrgTag
	idLengthCheck.ID = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, idLengthCheck.Validate())
		} else {
			assert.Nil(t, idLengthCheck.Validate())
		}
		idLengthCheck.ID = idLengthCheck.ID + "a"
	}

	timesCheck := validOrgTag
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
}
