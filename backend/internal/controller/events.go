package controller

import (
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"fmt"
)

// Creates a new event in the database
func (c Controller) CreateEvent(params event.Event) (e event.Event, err error) {
	e = params

	tx := c.database.Begin()

	res := tx.Create(&e)
	if res.RowsAffected < 1 {
		tx.Rollback()
		err = fmt.Errorf("no event created")
		return
	}
	if res.Error != nil {
		tx.Rollback()
		err = res.Error
		return
	}

	tx.Commit()
	return
}

// Deletes an event from the database by its ID
func (c Controller) DeleteEvent(id event_id.EventID) (event.Event, error) {
	tx := c.database.Begin()
	event := event.Event{}
	if err := tx.Where("id = ?", id.String()).First(&event).Error; err != nil {
		tx.Rollback()
		return event, fmt.Errorf("event with id=%s cannot be deleted because it doesn't exist", id)
	}

	res := tx.Delete(&event)

	if res.Error != nil {
		tx.Rollback()
		return event, res.Error
	}

	tx.Commit()
	return event, nil
}

// GET (/events/:id) a single event by their id
func (c Controller) GetEvent(id event_id.EventID) (event.Event, error) {
	event := event.Event{}

	tx := c.database.Begin()
	if err := tx.Where("id = ?", id).First(&event).Error; err != nil {
		tx.Rollback()
		return event, err
	}

	tx.Commit()
	return event, nil
}

// GET (/events) all events with pagination
func (c Controller) GetEvents(Limit int, Offset int) ([]event.Event, error) {
	events := []event.Event{}

	tx := c.database.Begin()
	if err := tx.Limit(Limit).Offset(Offset).Find(&events).Error; err != nil {
		return events, err
	}

	tx.Commit()
	return events, nil
}

// PUT (/events/:id) to completely update an existing event, returning the created object if successful
func (c Controller) PutEvent(id event_id.EventID, params event.Event) (event.Event, error) {
	event := event.Event{}

	tx := c.database.Begin()
	if err := tx.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}

	event.Name = params.Name
	event.Bio = params.Bio
	event.OrgID = params.OrgID
	event.Images = params.Images

	if err := tx.Save(&event).Error; err != nil {
		tx.Rollback()
		return event, err
	}

	tx.Commit()
	return event, nil
}

// PATCH (/events/:id) to partially update one or many fields of an existing event, returning the created object if successful
func (c Controller) PatchEvent(id event_id.EventID, params event.Event) (event.Event, error) {
	event := event.Event{}

	tx := c.database.Begin()
	if err := tx.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}
	event.Images = params.Images
	if params.Name != "" {
		event.Name = params.Name
	}
	if params.Bio != "" {
		event.Bio = params.Bio
	}
	if params.OrgID.String() != "" {
		event.OrgID = params.OrgID
	}
	if err := tx.Save(&event).Error; err != nil {
		tx.Rollback()
		return event, err
	}

	tx.Commit()
	return event, nil
}
