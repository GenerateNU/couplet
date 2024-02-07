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
	if valErr == nil {
		txErr = c.database.Create(&o).Error
	}
	return
}

// Deletes an organization from the database by its ID
func (c Controller) DeleteOrg(id org_id.OrgID) (o org.Org, txErr error) {
	// TODO: Write tests
	o, txErr = c.GetOrg(id)
	if txErr != nil {
		return
	}
	txErr = c.database.Delete(&o).Error
	return
}

// Gets an organization from the database by its ID
func (c Controller) GetOrg(id org_id.OrgID) (o org.Org, txErr error) {
	// TODO: Write tests
	txErr = c.database.Preload("OrgTags").First(&o, id).Error
	return
}

// Gets several organizations from the database
func (c Controller) GetOrgs(limit uint8, offset uint32) (orgs []org.Org, txErr error) {
	// TODO: Write tests
	txErr = c.database.Preload("OrgTags").Limit(int(limit)).Offset(int(offset)).Find(&orgs).Error
	return
}

// Updates an organization in the database
func (c Controller) UpdateOrg(params org.Org) (o org.Org, valErr error, txErr error) {
	// TODO: Write tests
	o = params
	valErr = o.Validate()
	if valErr == nil {
		txErr = c.database.Updates(&o).Error
	}
	return
}
