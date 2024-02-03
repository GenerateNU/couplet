package controller

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database"
	"fmt"
	"time"
)

func (c Controller) CreateEvent(ctx context.Context, event *api.Event) (*api.Event, error) {
	newEvent := database.Event{
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Name:           event.GetName(),
		Bio:            event.GetBio(),
		OrganizationID: database.OrganizationID(event.OrganizationID),
	}

	res := c.database.Create(&newEvent)
	if res.Error != nil {
		return nil, res.Error
	}

	event.SetID(api.NewOptEventId(api.EventId(newEvent.ID)))
	event.SetCreatedAt(api.NewOptDateTime(newEvent.CreatedAt))
	event.SetUpdatedAt(api.NewOptDateTime(newEvent.UpdatedAt))

	return event, nil
}

func (c Controller) DeleteEventById(ctx context.Context, event api.EventId) error {
	res := c.database.Delete(&database.Event{}, event)
	if res.RowsAffected == 0 {
		return fmt.Errorf("row with id=%v cannot be deleted because it doesn't exist", event)
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}
