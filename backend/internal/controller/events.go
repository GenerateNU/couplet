package controller

import (
	"context"
	"couplet/internal/api"
)


func (c Controller) CreateEvent(ctx context.Context, event *api.Event) (*api.Event, error) {
	res := c.database.Create(&event)

	if res.Error != nil {
		return nil, res.Error
	}
	return event, nil
}


func (c Controller) DeleteEventById(ctx context.Context, event api.EventId) (error) {
	res := c.database.Delete(&api.Event{}, event)

	if res.Error != nil {
		return res.Error
	}
	return nil
}


