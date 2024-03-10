package controller

import (
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
)

// Creates a new organization in the database
func (c Controller) CreateOrg(params org.Org) (o org.Org, valErr error, txErr error) {
	// TODO: Write tests
	o = params
	valErr = o.Validate()

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
	// TODO: Write tests
	tx := c.database.Begin()
	txErr = tx.Preload("OrgTags").First(&o, id).Error

	if txErr != nil {
		tx.Rollback()
		return
	}

	txErr = tx.Delete(&o).Error

	if txErr != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// Gets an organization from the database by its ID
func (c Controller) GetOrg(id org_id.OrgID) (o org.Org, txErr error) {
	// TODO: Write tests
	tx := c.database.Begin()

	txErr = tx.Preload("OrgTags").First(&o, id).Error

	if txErr != nil {
		tx.Rollback()
	}

	tx.Commit()
	return
}

// Gets several organizations from the database
func (c Controller) GetOrgs(limit uint8, offset uint32) (orgs []org.Org, txErr error) {
	// TODO: Write tests
	tx := c.database.Begin()

	txErr = tx.Preload("OrgTags").Limit(int(limit)).Offset(int(offset)).Find(&orgs).Error

	if txErr != nil {
		tx.Rollback()
	}

	tx.Commit()
	return
}

// Updates an organization in the database
func (c Controller) UpdateOrg(params org.Org) (o org.Org, valErr error, txErr error) {
	// TODO: Write tests
	o = params
	valErr = o.Validate()

	if valErr != nil {
		return
	}

	tx := c.database.Begin()
	txErr = tx.Updates(&o).Error

	if txErr != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}
