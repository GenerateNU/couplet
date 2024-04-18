package event_test

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventBeforeCreate(t *testing.T) {
	noIdEvent := event.Event{
		ID:        event_id.EventID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "The Events Company",
		Address:   "1234 Main St",
		Bio:       "At The Events Company, we connect people through events",
		Images:    url_slice.UrlSlice{},
		OrgID:     org_id.Wrap(uuid.New()),
	}
	require.Nil(t, (&noIdEvent).BeforeCreate(nil))
	assert.NotEmpty(t, noIdEvent.ID)
	id := noIdEvent.ID

	require.Nil(t, (&noIdEvent).BeforeCreate(nil))
	assert.Equal(t, id, noIdEvent.ID)
}
