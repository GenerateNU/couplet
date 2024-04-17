package controller

import (
	"couplet/internal/api"
	"couplet/internal/database/event"
	"couplet/internal/database/event_id"
	"couplet/internal/database/org_id"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"
)

// Creates a new event in the database
func (c Controller) CreateEvent(params event.Event) (e event.Event, valErr error, txErr error) {
	e = params
	var timestampErr error
	if e.UpdatedAt.Before(e.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if len(e.Name) < 1 || 255 < len(e.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(e.Name))
	}
	var bioLengthErr error
	if len(e.Bio) < 1 || 255 < len(e.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(e.Bio))
	}
	var imageCountErr error
	if len(e.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(e.Images))
	}

	var tagsCountErr error
	var tagsLengthErr error
	var tagsTimestampErr error
	if 5 < len(e.EventTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(e.EventTags))
	}
	for _, t := range e.EventTags {
		if len(t.ID) < 1 || 255 < len(t.ID) {
			tagsLengthErr = fmt.Errorf("invalid ID length of %d, must be in range [1,255]", len(t.ID))
		}
		if t.UpdatedAt.Before(t.CreatedAt) {
			tagsTimestampErr = fmt.Errorf("invalid timestamps")
		}
	}
	tagsErr := errors.Join(tagsCountErr, tagsLengthErr, tagsTimestampErr)
	var addressLengthError error
	if len(e.Address) < 1 || 255 < len(e.Address) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(e.Address))
	}
	var orgIdErr error
	if (e.OrgID == org_id.OrgID{}) {
		orgIdErr = fmt.Errorf("invalid org ID")
	}
	valErr = errors.Join(timestampErr, nameLengthErr, bioLengthErr, imageCountErr, tagsErr, addressLengthError, orgIdErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Create(&e).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Deletes an event from the database by its ID
func (c Controller) DeleteEvent(id event_id.EventID) (e event.Event, txErr error) {
	e.ID = id

	tx := c.database.Begin()
	txErr = tx.Clauses(clause.Returning{}).Preload("EventTags").Delete(&e).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Gets an event from the database by its ID
func (c Controller) GetEvent(id event_id.EventID) (e event.Event, txErr error) {
	txErr = c.database.Preload("EventTags").First(&e, id).Error
	return
}

// Gets several events from the database with pagination
func (c Controller) GetEvents(limit uint8, offset uint32) (events []event.Event, txErr error) {
	txErr = c.database.Limit(int(limit)).Offset(int(offset)).Preload("EventTags").Find(&events).Error
	return
}

// Creates a new event or updates an existing event in the database
func (c Controller) SaveEvent(params event.Event) (e event.Event, valErr error, txErr error) {
	e = params
	var timestampErr error
	if e.UpdatedAt.Before(e.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if len(e.Name) < 1 || 255 < len(e.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(e.Name))
	}
	var bioLengthErr error
	if len(e.Bio) < 1 || 255 < len(e.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(e.Bio))
	}
	var imageCountErr error
	if len(e.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(e.Images))
	}
	var tagsCountErr error
	var tagsLengthErr error
	var tagsTimestampErr error
	if 5 < len(e.EventTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(e.EventTags))
	}
	for _, t := range e.EventTags {
		if len(t.ID) < 1 || 255 < len(t.ID) {
			tagsLengthErr = fmt.Errorf("invalid ID length of %d, must be in range [1,255]", len(t.ID))
		}
		if t.UpdatedAt.Before(t.CreatedAt) {
			tagsTimestampErr = fmt.Errorf("invalid timestamps")
		}
	}
	tagsErr := errors.Join(tagsCountErr, tagsLengthErr, tagsTimestampErr)
	var orgIdErr error
	if (e.OrgID == org_id.OrgID{}) {
		orgIdErr = fmt.Errorf("invalid org ID")
	}
	valErr = errors.Join(timestampErr, nameLengthErr, bioLengthErr, imageCountErr, tagsErr, orgIdErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Save(&e).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Update one or many fields of an existing event in the database
func (c Controller) UpdateEvent(params event.Event) (e event.Event, valErr error, txErr error) {
	e = params
	var timestampErr error
	if e.UpdatedAt.Before(e.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if 255 < len(e.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(e.Name))
	}
	var bioLengthErr error
	if 255 < len(e.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(e.Bio))
	}
	var imageCountErr error
	if len(e.Images) != 0 && len(e.Images) != 4 {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be 4", len(e.Images))
	}
	var tagsCountErr error
	var tagsLengthErr error
	var tagsTimestampErr error
	if 5 < len(e.EventTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(e.EventTags))
	}
	for _, t := range e.EventTags {
		if len(t.ID) < 1 || 255 < len(t.ID) {
			tagsLengthErr = fmt.Errorf("invalid ID length of %d, must be in range [1,255]", len(t.ID))
		}
		if t.UpdatedAt.Before(t.CreatedAt) {
			tagsTimestampErr = fmt.Errorf("invalid timestamps")
		}
	}
	tagsErr := errors.Join(tagsCountErr, tagsLengthErr, tagsTimestampErr)
	valErr = errors.Join(timestampErr, nameLengthErr, bioLengthErr, imageCountErr, tagsErr)
	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Clauses(clause.Returning{}).Updates(&e).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (c Controller) GetRandomEvents(params api.RecommendationEventsGetParams) ([]event.Event, error) {
	var events []event.Event
	tx := c.database.Begin()
	query := tx.Order("random()")

	// Apply filters based on query parameters
	if params.Like.IsSet() && params.Like.Value {
		query = query.Where("liked = ?", true)
	}

	if len(params.Tags) > 0 {
		query = query.Where("event_tags.id IN (?)", params.Tags)
	}

	if params.Date.IsSet() {
		query = query.Where("date >= ?", params.Date.Value)
	}

	result := query.Offset(params.Offset.Value).Limit(params.Limit.Value).Find(&events)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return events, nil
}
