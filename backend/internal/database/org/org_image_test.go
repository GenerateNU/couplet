package org_test

import (
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrgImageValidate(t *testing.T) {
	validOrgImage := org.OrgImage{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Url:       "http://example.com/image.jpg",
		OrgID:     org_id.Wrap(uuid.New()),
	}
	assert.Nil(t, validOrgImage.Validate())
	assert.Nil(t, (&validOrgImage).BeforeSave(nil))

	timesCheck := validOrgImage
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))

	urlCheck := validOrgImage
	urlCheck.Url = ""
	assert.NotNil(t, urlCheck.Validate())
	assert.NotNil(t, (&urlCheck).BeforeSave(nil))

	ownerIdCheck := validOrgImage
	ownerIdCheck.OrgID = org_id.OrgID{}
	assert.NotNil(t, ownerIdCheck.Validate())
	assert.NotNil(t, (&ownerIdCheck).BeforeSave(nil))
}
