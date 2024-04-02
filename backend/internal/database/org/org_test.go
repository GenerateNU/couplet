package org_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"couplet/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrgBeforeCreate(t *testing.T) {
	noIdOrg := org.Org{
		ID:        org_id.OrgID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png")},
		OrgTags:   []org.OrgTag{{ID: "tag1"}, {ID: "tag2"}},
		Events:    []event.Event{},
	}
	require.Nil(t, (&noIdOrg).BeforeCreate(nil))
	assert.NotEmpty(t, noIdOrg.ID)
	id := noIdOrg.ID

	require.Nil(t, (&noIdOrg).BeforeCreate(nil))
	assert.Equal(t, id, noIdOrg.ID)
}
