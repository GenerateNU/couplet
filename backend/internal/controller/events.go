package controller

import (
	"context"
	"couplet/internal/api"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GET (/events/:id) a single event by their id
func (c *Controller) GetEventById(ctx context.Context, params api.GetEventByIdParams) (api.GetEventByIdRes, error) {
	var event api.Event
	result := c.database.First(&event, "id = ?", params.EventId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("event with id %s not found", params.EventId)
		}
		return nil, result.Error
	}

	return &event, nil
}

// GET (/events) all events with pagination (ask if you don't know what this is)
func (c *Controller) GetAllEvents(ctx context.Context, params api.GetAllEventsParams) (api.GetAllEventsRes, error) {
	var events []api.Event
	result := c.database.Limit(params.Limit.Value).Offset(params.Offset.Value).Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}

	return events, nil // fix this issue
}

// PUT (/events/:id) to completely update an existing event, returning the created object if successful
func (c *Controller) UpdateEventById(ctx context.Context, req *api.Event, params api.UpdateEventByIdParams) (api.UpdateEventByIdRes, error) {
	result := c.database.Model(&api.Event{}).Where("id = ?", params.EventId).Updates(req)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("event with id %s not found", params.EventId)
	}

	return req, nil
}

// PATCH (/events/:id) to partially update one or many fields of an existing event, returning the created object if successful
func (c *Controller) PatchEventById(ctx context.Context, req *api.Event, params api.PatchEventByIdParams) (api.PatchEventByIdRes, error) {
	result := c.database.Model(&api.Event{}).Where("id = ?", params.EventId).Updates(req)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("event with id %s not found", params.EventId)
	}

	return req, nil
}
