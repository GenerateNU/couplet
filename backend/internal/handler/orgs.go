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
	if h.logger != nil {
		h.logger.Info("POST /orgs")
	}

	orgTags := []org.OrgTag{}
	for _, v := range req.Tags {
		orgTags = append(orgTags, org.OrgTag{ID: v})
	}

	o, valErr, txErr := h.controller.CreateOrg(org.Org{
		Name:    req.Name,
		Bio:     req.Bio,
		Images:  url_slice.Wrap(req.Images),
		OrgTags: orgTags,
	})
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: "failed to validate organization",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to create organization")
	}

	tags := []string{}
	for _, orgTag := range o.OrgTags {
		tags = append(tags, orgTag.ID)
	}
	res := api.OrgsPostCreated{
		ID:     o.ID.Unwrap(),
		Name:   o.Name,
		Bio:    o.Bio,
		Images: o.Images.Unwrap(),
		Tags:   tags,
	}
	return &res, nil
}

// Deletes an organization by its ID.
// DELETE /orgs/{id}
func (h Handler) OrgsIDDelete(ctx context.Context, params api.OrgsIDDeleteParams) (api.OrgsIDDeleteRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("DELETE /orgs/%s", params.ID))
	}

	o, txErr := h.controller.DeleteOrg(org_id.Wrap(params.ID))
	if txErr != nil {
		return &api.Error{
			Code:    404,
			Message: "organization not found",
		}, nil
	}

	tags := []string{}
	for _, orgTag := range o.OrgTags {
		tags = append(tags, orgTag.ID)
	}
	res := api.OrgsIDDeleteOK{
		ID:     o.ID.Unwrap(),
		Name:   o.Name,
		Bio:    o.Bio,
		Images: o.Images.Unwrap(),
		Tags:   tags,
	}
	return &res, nil
}

// Gets an organization by its ID.
// GET /orgs/{id}
func (h Handler) OrgsIDGet(ctx context.Context, params api.OrgsIDGetParams) (api.OrgsIDGetRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("GET /orgs/%s", params.ID))
	}

	o, txErr := h.controller.GetOrg(org_id.Wrap(params.ID))
	if txErr != nil {
		return &api.Error{
			Code:    404,
			Message: "organization not found",
		}, nil
	}

	tags := []string{}
	for _, orgTag := range o.OrgTags {
		tags = append(tags, orgTag.ID)
	}
	res := api.OrgsIDGetOK{
		ID:     o.ID.Unwrap(),
		Name:   o.Name,
		Bio:    o.Bio,
		Images: o.Images.Unwrap(),
		Tags:   tags,
	}
	return &res, nil
}

// Gets multiple organizations.
// GET /orgs
func (h Handler) OrgsGet(ctx context.Context, params api.OrgsGetParams) ([]api.OrgsGetOKItem, error) {
	if h.logger != nil {
		h.logger.Info("GET /orgs")
	}

	limit := params.Limit.Value   // default value makes this safe
	offset := params.Offset.Value // default value makes this safe
	orgs, txErr := h.controller.GetOrgs(limit, offset)
	if txErr != nil {
		return nil, errors.New("failed to get organizations")
	}
	res := []api.OrgsGetOKItem{}
	for _, o := range orgs {
		tags := []string{}
		for _, orgTag := range o.OrgTags {
			tags = append(tags, orgTag.ID)
		}
		item := api.OrgsGetOKItem{
			ID:     o.ID.Unwrap(),
			Name:   o.Name,
			Bio:    o.Bio,
			Images: o.Images.Unwrap(),
			Tags:   tags,
		}
		res = append(res, item)
	}
	return res, nil
}

// Partially updates an organization by its ID.
// PATCH /orgs/{id}
func (h Handler) OrgsIDPatch(ctx context.Context, req *api.Org, params api.OrgsIDPatchParams) (api.OrgsIDPatchRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PATCH /orgs/%s", params.ID))
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
	reqOrg.OrgTags = []org.OrgTag{}
	for _, v := range req.Tags {
		reqOrg.OrgTags = append(reqOrg.OrgTags, org.OrgTag{ID: v})
	}

	o, valErr, txErr := h.controller.UpdateOrg(reqOrg)
	if valErr != nil {
		return &api.OrgsIDPatchBadRequest{
			Code:    400,
			Message: "failed to validate organization",
		}, nil
	}
	if txErr != nil {
		return &api.OrgsIDPatchNotFound{
			Code:    404,
			Message: "organization not found",
		}, nil
	}

	tags := []string{}
	for _, orgTag := range o.OrgTags {
		tags = append(tags, orgTag.ID)
	}
	res := api.OrgsIDPatchOK{
		ID:     o.ID.Unwrap(),
		Name:   o.Name,
		Bio:    o.Bio,
		Images: o.Images.Unwrap(),
		Tags:   tags,
	}
	return &res, nil
}

// Updates an organization by its ID.
// PUT /orgs/{id}
func (h Handler) OrgsIDPut(ctx context.Context, req *api.OrgsIDPutReq, params api.OrgsIDPutParams) (api.OrgsIDPutRes, error) {
	if h.logger != nil {
		h.logger.Info(fmt.Sprintf("PUT /orgs/%s", params.ID))
	}

	orgTags := []org.OrgTag{}
	for _, v := range req.Tags {
		orgTags = append(orgTags, org.OrgTag{ID: v})
	}

	o, valErr, txErr := h.controller.SaveOrg(org.Org{
		Name:    req.Name,
		Bio:     req.Bio,
		Images:  url_slice.Wrap(req.Images),
		OrgTags: orgTags,
	})
	if valErr != nil {
		return &api.Error{
			Code:    400,
			Message: "failed to validate user",
		}, nil
	}
	if txErr != nil {
		return nil, errors.New("failed to save user")
	}

	tags := []string{}
	for _, orgTag := range o.OrgTags {
		tags = append(tags, orgTag.ID)
	}
	res := api.OrgsIDPutOK{
		ID:     o.ID.Unwrap(),
		Name:   o.Name,
		Bio:    o.Bio,
		Images: o.Images.Unwrap(),
		Tags:   tags,
	}
	return &res, nil
}
