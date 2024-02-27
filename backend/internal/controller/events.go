package controller

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"fmt"
)

// Creates a new event in the database
func (c Controller) CreateEvent(params event.Event) (e *event.Event, err error) {
	newEvent := event.Event{
		Name:  params.Name,
		Bio:   params.Bio,
		OrgID: params.OrgID,
	}

	res := c.database.Create(&newEvent)
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("no event created")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return &newEvent, nil
}

// Deletes an event from the database by its ID
func (c Controller) DeleteEvent(id event_id.EventID) (event.Event, error) {
	// TODO: Do this in one transaction
	event := event.Event{}
	if err := c.database.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}

	res := c.database.Delete(&event)
	if res.RowsAffected == 0 {
		return event, fmt.Errorf("event with id=%s cannot be deleted because it doesn't exist", id)
	}

	if res.Error != nil {
		return event, res.Error
	}

	return event, nil
}

// GET (/events/:id) a single event by their id
func (c Controller) GetEvent(id event_id.EventID) (event.Event, error) {
	event := event.Event{}
	if err := c.database.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}
	return event, nil
}

// GET (/events) all events with pagination
func (c Controller) GetEvents(Limit int, Offset int) ([]event.Event, error) {
	events := []event.Event{}
	if err := c.database.Limit(Limit).Offset(Offset).Find(&events).Error; err != nil {
		return events, err
	}
	return events, nil
}

// PUT (/events/:id) to completely update an existing event, returning the created object if successful
func (c Controller) PutEvent(id event_id.EventID, params event.Event) (event.Event, error) {
	event := event.Event{}
	if err := c.database.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}
	event.Name = params.Name
	event.Bio = params.Bio
	event.OrgID = params.OrgID
	if err := c.database.Save(&event).Error; err != nil {
		return event, err
	}
	return event, nil
}

// PATCH (/events/:id) to partially update one or many fields of an existing event, returning the created object if successful
func (c Controller) PatchEvent(id event_id.EventID, params event.Event) (event.Event, error) {
	event := event.Event{}
	if err := c.database.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}
	if params.Name != "" {
		event.Name = params.Name
	}
	if params.Bio != "" {
		event.Bio = params.Bio
	}
	if params.OrgID.String() != "" {
		event.OrgID = params.OrgID
	}
	if err := c.database.Save(&event).Error; err != nil {
		return event, err
	}
	return event, nil
}
