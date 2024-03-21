package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/org"
	"couplet/internal/database/org_id"
	"couplet/internal/database/url_slice"
	"errors"
	"fmt"
)

// Creates a new organization.
// POST /orgs
func (h Handler) OrgsPost(ctx context.Context, req *api.OrgsPostReq) (api.OrgsPostRes, error) {
	// TODO: Write tests
	h.logger.Info("POST /orgs")

	var orgToCreate org.Org
	orgToCreate.Name = req.Name
	orgToCreate.Bio = req.Bio
	orgToCreate.Images = url_slice.Wrap(req.Images)
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

	res := api.OrgsPostCreated{
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

// Deletes an organization by its ID.
// DELETE /orgs/{id}
func (h Handler) OrgsIDDelete(ctx context.Context, params api.OrgsIDDeleteParams) (api.OrgsIDDeleteRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("DELETE /orgs/%s", params.ID))
	o, err := h.controller.DeleteOrg(org_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.OrgsIDDeleteOK{
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

// Gets an organization by its ID.
// GET /orgs/{id}
func (h Handler) OrgsIDGet(ctx context.Context, params api.OrgsIDGetParams) (api.OrgsIDGetRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("GET /orgs/%s", params.ID))
	o, err := h.controller.GetOrg(org_id.Wrap(params.ID))
	if err != nil {
		return &api.Error{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	res := api.OrgsIDGetOK{
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

// Gets multiple organizations.
// GET /orgs
func (h Handler) OrgsGet(ctx context.Context, params api.OrgsGetParams) ([]api.OrgsGetOKItem, error) {
	// TODO: Write tests
	h.logger.Info("GET /orgs")
	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	orgs, err := h.controller.GetOrgs(limit, offset)
	res := []api.OrgsGetOKItem{}
	for _, o := range orgs {
		item := api.OrgsGetOKItem{
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

// Partially updates an organization by its ID.
// PATCH /orgs/{id}
func (h Handler) OrgsIDPatch(ctx context.Context, req *api.Org, params api.OrgsIDPatchParams) (api.OrgsIDPatchRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("PATCH /orgs/%s", params.ID))

	_, getErr := h.controller.GetOrg(org_id.Wrap(params.ID))
	doesNotExist := getErr != nil
	if doesNotExist {
		return &api.OrgsIDPatchNotFound{
			Code:    404,
			Message: getErr.Error(),
		}, nil
	}

	var reqOrg org.Org
	reqOrg.ID = org_id.Wrap(params.ID)
	if req.Name.Set {
		reqOrg.Name = req.Name.Value
	}
	if req.Bio.Set {
		reqOrg.Bio = req.Bio.Value
	}
	reqOrg.Images = url_slice.Wrap(req.Images)
	if len(req.Tags) > 0 {
		reqOrg.OrgTags = []org.OrgTag{}
		for _, v := range req.Tags {
			reqOrg.OrgTags = append(reqOrg.OrgTags, org.OrgTag{ID: v})
		}
	}

	o, valErr, txErr := h.controller.UpdateOrg(reqOrg)
	if valErr != nil {
		return &api.OrgsIDPatchBadRequest{
			Code:    400,
			Message: valErr.Error(),
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to update organization")
	}
	res := api.OrgsIDPatchOK{
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

// Updates an organization by its ID.
// PUT /orgs/{id}
func (h Handler) OrgsIDPut(ctx context.Context, req *api.OrgsIDPutReq, params api.OrgsIDPutParams) (api.OrgsIDPutRes, error) {
	// TODO: Write tests
	h.logger.Info(fmt.Sprintf("PUT /orgs/%s", params.ID))

	_, getErr := h.controller.GetOrg(org_id.Wrap(params.ID))
	alreadyExists := getErr == nil

	var reqOrg org.Org
	reqOrg.ID = org_id.Wrap(params.ID)
	reqOrg.Name = req.Name
	reqOrg.Bio = req.Bio
	reqOrg.Images = url_slice.Wrap(req.Images)
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
		res := api.OrgsIDPutOK{
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

	res := api.OrgsIDPutCreated{
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
