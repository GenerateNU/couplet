package controller

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/event"
	eventId "couplet/internal/database/event/id"
	orgId "couplet/internal/database/org/id"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c Controller) CreateEvent(ctx context.Context, apiEvent *api.Event) (*api.Event, error) {
	newEvent := event.Event{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      apiEvent.GetName(),
		Bio:       apiEvent.GetBio(),
		OrgID:     orgId.OrgID(apiEvent.GetOrganizationID()),
	}

	res := c.database.Create(&newEvent)
	fmt.Println("created:", newEvent.ID)
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("no event created")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	apiEvent.SetID(api.NewOptEventId(api.EventId(newEvent.ID)))
	apiEvent.SetCreatedAt(api.NewOptDateTime(newEvent.CreatedAt))
	apiEvent.SetUpdatedAt(api.NewOptDateTime(newEvent.UpdatedAt))

	return apiEvent, nil
}

func (c Controller) DeleteEventById(ctx context.Context, apiEvent api.EventId) error {
	res := c.database.Where("id = ?", eventId.EventID(apiEvent)).Delete(&event.Event{})
	if res.RowsAffected == 0 {
		return fmt.Errorf("404 - event with id=%v cannot be deleted because it doesn't exist", uuid.UUID(apiEvent).String())
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}
