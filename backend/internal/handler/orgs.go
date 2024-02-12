package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"errors"
	"fmt"
)

// Create a new organization.
// POST /orgs
func (h Handler) CreateOrg(ctx context.Context, req *api.CreateOrgReq) (api.CreateOrgRes, error) {
	// TODO: Write tests
	h.logger.Info("POST /orgs")

	var orgToCreate org.Org
	orgToCreate.Name = req.Name
	orgToCreate.Bio = req.Bio
	if req.Image.Set {
		orgToCreate.Image = req.Image.Value.String()
	}
	orgToCreate.OrgTags = []org.OrgTag{}
	for _, v := range req.Tags {
		orgToCreate.OrgTags = append(orgToCreate.OrgTags, org.OrgTag{ID: v})
	}

	o, valErr, txErr := h.controller.CreateOrg(orgToCreate)
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: valErr.Error(),
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to create organization")
	}

	res := api.CreateOrgCreated{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
		Tags: []string{},
	}
	for _, orgTag := range o.OrgTags {
		res.Tags = append(res.Tags, orgTag.ID)
	}

	return &res, nil
}

// Delete an organization by its ID.
// DELETE /orgs/{orgId}
func (h Handler) DeleteOrgById(ctx context.Context, params api.DeleteOrgByIdParams) (api.DeleteOrgByIdRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("DELETE /orgs/%s", params.ID))
	o, err := h.controller.DeleteOrg(org_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.DeleteOrgByIdOK{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
		Tags: []string{},
	}
	for _, orgTag := range o.OrgTags {
		res.Tags = append(res.Tags, orgTag.ID)
	}
	return &res, nil
}

// Get an organization by its ID.
// GET /orgs/{orgId}
func (h Handler) GetOrgById(ctx context.Context, params api.GetOrgByIdParams) (api.GetOrgByIdRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("GET /orgs/%s", params.ID))
	o, err := h.controller.GetOrg(org_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.GetOrgByIdOK{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
		Tags: []string{},
	}
	for _, orgTag := range o.OrgTags {
		res.Tags = append(res.Tags, orgTag.ID)
	}
	return &res, nil
}

// Get multiple organizations.
// GET /orgs
func (h Handler) GetOrgs(ctx context.Context, params api.GetOrgsParams) ([]api.GetOrgsOKItem, error) {
	// TODO: Write tests
	h.logger.Info("GET /orgs")
	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	orgs, err := h.controller.GetOrgs(limit, offset)
	res := []api.GetOrgsOKItem{}
	for _, o := range orgs {
		item := api.GetOrgsOKItem{
			ID:   o.ID.Unwrap(),
			Name: o.Name,
			Bio:  o.Bio,
			Tags: []string{},
		}
		for _, orgTag := range o.OrgTags {
			item.Tags = append(item.Tags, orgTag.ID)
		}
		res = append(res, item)
	}
	return res, err
}

// Partially update an organization by its ID.
// PATCH /orgs/{orgId}
func (h Handler) PartialUpdateOrgById(ctx context.Context, req *api.Org, params api.PartialUpdateOrgByIdParams) (api.PartialUpdateOrgByIdRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("PATCH /orgs/%s", params.ID))

	_, getErr := h.controller.GetOrg(org_id.OrgID(params.ID))
	doesNotExist := getErr != nil
	if doesNotExist {
		return &api.PartialUpdateOrgByIdNotFound{
			Code:    404,
			Message: getErr.Error(),
		}, nil
	}

	var reqOrg org.Org
	reqOrg.ID = org_id.OrgID(params.ID)
	if req.Name.Set {
		reqOrg.Name = req.Name.Value
	}
	if req.Bio.Set {
		reqOrg.Bio = req.Bio.Value
	}
	if req.Image.Set {
		reqOrg.Image = req.Image.Value.String()
	}
	if len(req.Tags) > 0 {
		reqOrg.OrgTags = []org.OrgTag{}
		for _, v := range req.Tags {
			reqOrg.OrgTags = append(reqOrg.OrgTags, org.OrgTag{ID: v})
		}
	}

	o, valErr, txErr := h.controller.UpdateOrg(reqOrg)
	if valErr != nil {
		return &api.PartialUpdateOrgByIdBadRequest{
			Code:    400,
			Message: valErr.Error(),
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to update organization")
	}
	res := api.PartialUpdateOrgByIdOK{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
		Tags: []string{},
	}
	for _, orgTag := range o.OrgTags {
		res.Tags = append(res.Tags, orgTag.ID)
	}
	return &res, nil
}

// Update an organization by its ID.
// PUT /orgs/{orgId}
func (h Handler) SaveOrgById(ctx context.Context, req *api.SaveOrgByIdReq, params api.SaveOrgByIdParams) (api.SaveOrgByIdRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("PUT /orgs/%s", params.ID))

	_, getErr := h.controller.GetOrg(org_id.OrgID(params.ID))
	alreadyExists := getErr == nil

	var reqOrg org.Org
	reqOrg.ID = org_id.OrgID(params.ID)
	reqOrg.Name = req.Name
	reqOrg.Bio = req.Bio
	if req.Image.Set {
		reqOrg.Image = req.Image.Value.String()
	}
	reqOrg.OrgTags = []org.OrgTag{}
	for _, v := range req.Tags {
		reqOrg.OrgTags = append(reqOrg.OrgTags, org.OrgTag{ID: v})
	}

	if alreadyExists {
		o, valErr, txErr := h.controller.UpdateOrg(reqOrg)
		if valErr != nil {
			return &api.Error{
				Code:    400,
				Message: valErr.Error(),
			}, nil
		}
		if txErr != nil {
			return nil, errors.New("failed to update organization")
		}
		res := api.SaveOrgByIdOK{
			ID:   o.ID.Unwrap(),
			Name: o.Name,
			Bio:  o.Bio,
			Tags: []string{},
		}
		for _, orgTag := range o.OrgTags {
			res.Tags = append(res.Tags, orgTag.ID)
		}
		return &res, nil
	}

	o, valErr, txErr := h.controller.CreateOrg(reqOrg)
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: valErr.Error(),
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to create organization")
	}

	res := api.SaveOrgByIdCreated{
		ID:   o.ID.Unwrap(),
		Name: o.Name,
		Bio:  o.Bio,
		Tags: []string{},
	}
	for _, orgTag := range o.OrgTags {
		res.Tags = append(res.Tags, orgTag.ID)
	}

	return &res, nil
}
