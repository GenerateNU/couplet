package org_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrgValidate(t *testing.T) {
	validOrg := org.Org{
		ID:        org_id.Wrap(uuid.New()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Image:     org.OrgImage{Url: "https://example.com/image.png"},
		OrgTags:   []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}},
		Events:    []event.Event{{ID: event_id.Wrap(uuid.New())}},
	}
	assert.Nil(t, validOrg.Validate())

	timesCheck := validOrg
	timesCheck.CreatedAt = timesCheck.UpdatedAt.Add(1)
	assert.NotNil(t, timesCheck.Validate())

	nameLengthCheck := validOrg
	nameLengthCheck.Name = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, nameLengthCheck.Validate())
		} else {
			assert.Nil(t, nameLengthCheck.Validate())
		}
		nameLengthCheck.Name = nameLengthCheck.Name + "a"
	}

	bioLengthCheck := validOrg
	bioLengthCheck.Bio = ""
	for i := 0; i <= 256; i++ {
		if i < 1 || i > 255 {
			assert.NotNil(t, bioLengthCheck.Validate())
		} else {
			assert.Nil(t, bioLengthCheck.Validate())
		}
		bioLengthCheck.Bio = bioLengthCheck.Bio + "a"
	}

	noImageCheck := validOrg
	noImageCheck.Image = org.OrgImage{}
	assert.Nil(t, noImageCheck.Validate())

	imageUrlCheck := validOrg
	imageUrlCheck.Image = org.OrgImage{Url: "invalid"}
	assert.NotNil(t, imageUrlCheck.Validate())

	orgTagsCheck := validOrg
	orgTagsCheck.OrgTags = []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}}
	assert.NotNil(t, orgTagsCheck.Validate())
}
