package handler

import (
	"context"
	"couplet/internal/api"
	"fmt"

	ht "github.com/ogen-go/ogen/http"
)

// Create a new organization.
// POST /orgs
func (h Handler) CreateOrg(ctx context.Context, req *api.Org) (api.CreateOrgRes, error) {
	h.logger.Info("POST /orgs")
	return &api.Org{}, ht.ErrNotImplemented
}

// Delete an organization by its ID.
// DELETE /orgs/{orgId}
func (h Handler) DeleteOrgById(ctx context.Context, params api.DeleteOrgByIdParams) (api.DeleteOrgByIdRes, error) {
	h.logger.Info(fmt.Sprintf("DELETE /orgs/%s", params.OrgId))
	return &api.Org{}, ht.ErrNotImplemented
}

// Get an organization by its ID.
// GET /orgs/{orgId}
func (h Handler) GetOrgById(ctx context.Context, params api.GetOrgByIdParams) (api.GetOrgByIdRes, error) {
	h.logger.Info(fmt.Sprintf("GET /orgs/%s", params.OrgId))
	return &api.Org{}, ht.ErrNotImplemented
}

// Get multiple organizations.
// GET /orgs
func (h Handler) GetOrgs(ctx context.Context, params api.GetOrgsParams) ([]api.Org, error) {
	h.logger.Info("GET /orgs")
	return []api.Org{}, ht.ErrNotImplemented
}

// Partially update an organization by its ID.
// PATCH /orgs/{orgId}
func (h Handler) PartialUpdateOrgById(ctx context.Context, req *api.Org, params api.PartialUpdateOrgByIdParams) (api.PartialUpdateOrgByIdRes, error) {
	h.logger.Info(fmt.Sprintf("PATCH /orgs/%s", params.OrgId))
	return &api.Org{}, ht.ErrNotImplemented
}

// Update an organization by its ID.
// PUT /orgs/{orgId}
func (h Handler) UpdateOrgById(ctx context.Context, req *api.Org, params api.UpdateOrgByIdParams) (api.UpdateOrgByIdRes, error) {
	h.logger.Info(fmt.Sprintf("PUT /orgs/%s", params.OrgId))
	return &api.Org{}, ht.ErrNotImplemented
}
