package org_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrgValidate(t *testing.T) {
	id := org_id.Wrap(uuid.New())
	validOrg := org.Org{
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Image:     org.OrgImage{Url: "https://example.com/image.png", OrgID: id},
		OrgTags:   []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}},
		Events:    []event.Event{{OrgID: id}},
	}
	assert.Nil(t, validOrg.Validate())
	assert.Nil(t, (&validOrg).BeforeSave(nil))

	timesCheck := validOrg
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())
	assert.NotNil(t, (&timesCheck).BeforeSave(nil))

	nameLengthCheck := validOrg
	nameLengthCheck.Name = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, nameLengthCheck.Validate())
			assert.NotNil(t, (&nameLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, nameLengthCheck.Validate())
			assert.Nil(t, (&nameLengthCheck).BeforeSave(nil))
		}
		nameLengthCheck.Name = nameLengthCheck.Name + "a"
	}

	bioLengthCheck := validOrg
	bioLengthCheck.Bio = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, bioLengthCheck.Validate())
			assert.NotNil(t, (&bioLengthCheck).BeforeSave(nil))
		} else {
			assert.Nil(t, bioLengthCheck.Validate())
			assert.Nil(t, (&bioLengthCheck).BeforeSave(nil))
		}
		bioLengthCheck.Bio = bioLengthCheck.Bio + "a"
	}

	noImageCheck := validOrg
	noImageCheck.Image = org.OrgImage{}
	assert.Nil(t, noImageCheck.Validate())
	assert.Nil(t, (&noImageCheck).BeforeSave(nil))

	orgTagsCheck := validOrg
	orgTagsCheck.OrgTags = []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}}
	assert.NotNil(t, orgTagsCheck.Validate())
	assert.NotNil(t, (&orgTagsCheck).BeforeSave(nil))
}

func TestOrgBeforeCreate(t *testing.T) {
	noIdOrg := org.Org{
		ID:        org_id.OrgID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Image:     org.OrgImage{},
		OrgTags:   []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}},
		Events:    []event.Event{},
	}
	require.Nil(t, (&noIdOrg).BeforeCreate(nil))
	assert.NotEmpty(t, noIdOrg.ID)
	id := noIdOrg.ID

	require.Nil(t, (&noIdOrg).BeforeCreate(nil))
	assert.Equal(t, id, noIdOrg.ID)
}
