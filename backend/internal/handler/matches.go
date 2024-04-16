package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/database/user_id"
)

// MatchesIDGet implements api.Handler.
func (h Handler) MatchesIDGet(ctx context.Context, params api.MatchesIDGetParams) ([]api.MatchesIDGetOKItem, error) {
	// TODO: Write tests
	h.logger.Info("GET /matches/{id}")
	userMatches, err := h.controller.GetUserMatches(user_id.Wrap(params.ID))
	res := []api.MatchesIDGetOKItem{}
	for _, m := range userMatches {
		item := api.MatchesIDGetOKItem{
			// ID:        m.ID.Unwrap(),
			// CreatedAt: m.CreatedAt,
			// UpdatedAt: m.UpdatedAt,
			FirstName: api.NewOptString(m.FirstName),
			LastName:  api.NewOptString(m.LastName),
			Age:       api.NewOptUint8(m.Age),
			// Bio:       m.Bio,
			Images: m.Images,
			// Viewed:    m.Viewed,
		}
		res = append(res, item)
	}
	return res, err
}
