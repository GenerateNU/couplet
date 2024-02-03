// Handles API requests and translate between internal and external schema
package handler

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
)

// Handles incoming API requests
type Handler struct {
	controller controller.Controller // executes business logic
}

// Creates a new handler for all defined API endpoints
func NewHandler(controller controller.Controller) api.Handler {
	return Handler{
		controller,
	}
}

// Checks if the server is running and servicing requests.
// GET /health-check
func (h Handler) HealthCheck(ctx context.Context) error {
	return nil
}
