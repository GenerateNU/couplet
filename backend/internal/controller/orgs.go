package controller

import (
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"
)

// Creates a new organization in the database
func (c Controller) CreateOrg(params org.Org) (o org.Org, valErr error, txErr error) {
	o = params
	var timestampErr error
	if o.UpdatedAt.Before(o.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if len(o.Name) < 1 || 255 < len(o.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(o.Name))
	}
	var bioLengthErr error
	if len(o.Bio) < 1 || 255 < len(o.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(o.Bio))
	}
	var imageCountErr error
	if len(o.Images) < 1 || 4 < len(o.Images) {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be in range [1,4]", len(o.Images))
	}

	var tagsCountErr error
	var tagsLengthErr error
	var tagsTimestampErr error
	if 5 < len(o.OrgTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(o.OrgTags))
	}
	for _, t := range o.OrgTags {
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
	txErr = tx.Create(&o).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Deletes an organization from the database by its ID
func (c Controller) DeleteOrg(id org_id.OrgID) (o org.Org, txErr error) {
	o.ID = id

	tx := c.database.Begin()
	txErr = tx.Clauses(clause.Returning{}).Preload("OrgTags").Delete(&o).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Gets an organization from the database by its ID
func (c Controller) GetOrg(id org_id.OrgID) (o org.Org, txErr error) {
	txErr = c.database.Preload("OrgTags").First(&o, id).Error
	return
}

// Gets several organizations from the database with pagination
func (c Controller) GetOrgs(limit uint8, offset uint32) (orgs []org.Org, txErr error) {
	txErr = c.database.Limit(int(limit)).Offset(int(offset)).Preload("OrgTags").Find(&orgs).Error
	return
}

// Creates a new org or updates an existing org in the database
func (c Controller) SaveOrg(params org.Org) (o org.Org, valErr error, txErr error) {
	o = params
	var timestampErr error
	if o.UpdatedAt.Before(o.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if len(o.Name) < 1 || 255 < len(o.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(o.Name))
	}
	var bioLengthErr error
	if len(o.Bio) < 1 || 255 < len(o.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(o.Bio))
	}
	var imageCountErr error
	if len(o.Images) < 1 || 4 < len(o.Images) {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be in range [1,4]", len(o.Images))
	}

	var tagsCountErr error
	var tagsLengthErr error
	var tagsTimestampErr error
	if 5 < len(o.OrgTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(o.OrgTags))
	}
	for _, t := range o.OrgTags {
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
	txErr = tx.Save(&o).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// Update one or many fields of an existing org in the database
func (c Controller) UpdateOrg(params org.Org) (o org.Org, valErr error, txErr error) {
	o = params
	var timestampErr error
	if o.UpdatedAt.Before(o.CreatedAt) {
		timestampErr = fmt.Errorf("invalid timestamps")
	}
	var nameLengthErr error
	if 255 < len(o.Name) {
		nameLengthErr = fmt.Errorf("invalid name length of %d, must be in range [1,255]", len(o.Name))
	}
	var bioLengthErr error
	if 255 < len(o.Bio) {
		bioLengthErr = fmt.Errorf("invalid bio length of %d, must be in range [1,255]", len(o.Bio))
	}
	var imageCountErr error
	if 4 < len(o.Images) {
		imageCountErr = fmt.Errorf("invalid image count of %d, must be in range [1,4]", len(o.Images))
	}

	var tagsCountErr error
	var tagsLengthErr error
	var tagsTimestampErr error
	if 5 < len(o.OrgTags) {
		tagsCountErr = fmt.Errorf("invalid tag count of %d, must be in range [0,5]", len(o.OrgTags))
	}
	for _, t := range o.OrgTags {
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
	txErr = tx.Clauses(clause.Returning{}).Updates(&o).Error
	if txErr != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
