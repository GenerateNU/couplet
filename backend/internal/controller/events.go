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

func (c Controller) DeleteEventById(ctx context.Context, apiEvent api.EventId) (*api.Event, error) {
	eventObj := &event.Event{}
	if err := c.database.Where("id = ?", eventId.EventID(apiEvent)).First(&eventObj).Error; err != nil {
		return nil, err
	}

	res := c.database.Delete(&eventObj)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("event with id=%v cannot be deleted because it doesn't exist", uuid.UUID(apiEvent).String())
	}

	if res.Error != nil {
		return nil, res.Error
	}

	deletedEvent := &api.Event{
		ID:             api.NewOptEventId(apiEvent),
		CreatedAt:      api.NewOptDateTime(eventObj.CreatedAt),
		UpdatedAt:      api.NewOptDateTime(eventObj.UpdatedAt),
		Name:           eventObj.Name,
		Bio:            eventObj.Bio,
		OrganizationID: eventObj.OrgID.UUID(),
	}

	return deletedEvent, nil
}
