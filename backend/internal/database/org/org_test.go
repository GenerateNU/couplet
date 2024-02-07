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
		Image:     "https://example.com/image.png",
		OrgTags:   []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}},
		Events:    []event.Event{{ID: event_id.Wrap(uuid.New())}},
	}
	assert.Nil(t, validOrg.Validate())

	invalidTimes := validOrg
	invalidTimes.CreatedAt = invalidTimes.UpdatedAt.Add(1)
	assert.NotNil(t, invalidTimes.Validate())

	invalidName := validOrg
	invalidName.Name = ""
	assert.NotNil(t, invalidName.Validate())
	for i := 0; i < 256; i++ {
		invalidName.Name = invalidName.Name + "a"
	}
	assert.NotNil(t, invalidName.Validate())

	validNoImage := validOrg
	validNoImage.Image = ""
	assert.Nil(t, validNoImage.Validate())

	invalidImage := validOrg
	invalidImage.Image = "invalid"
	assert.NotNil(t, invalidImage.Validate())

	invalidOrgTags := validOrg
	invalidOrgTags.OrgTags = []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}, {ID: "tag3"}, {ID: "tag4"}, {ID: "tag5"}, {ID: "tag6"}}
	assert.NotNil(t, invalidOrgTags.Validate())
}

func TestOrgTagValidate(t *testing.T) {
	validOrgTag := org.OrgTag{
		ID:        "tag",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Orgs:      []org.Org{{ID: org_id.Wrap(uuid.New())}, {ID: org_id.Wrap(uuid.New())}},
	}
	assert.Nil(t, validOrgTag.Validate())

	invalidTimes := validOrgTag
	invalidTimes.CreatedAt = invalidTimes.UpdatedAt.Add(1)
	assert.NotNil(t, invalidTimes.Validate())
}
