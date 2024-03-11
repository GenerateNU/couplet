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
