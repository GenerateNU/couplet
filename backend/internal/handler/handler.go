// Handles API requests and translate between internal and external schema
package handler

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target ../api --clean ../../../openapi.yaml

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	"log/slog"
)

// Handles incoming API requests
type Handler struct {
	controller controller.Controller // executes business logic
	logger     *slog.Logger          // event logger
}

// Creates a new handler for all defined API endpoints
func NewHandler(controller controller.Controller, logger *slog.Logger) api.Handler {
	return Handler{
		controller,
		logger,
	}
}

// Checks if the server is running and servicing requests.
// GET /health-check
func (h Handler) HealthCheckGet(ctx context.Context) error {
	h.logger.Info("GET /health-check")
	return nil
}
