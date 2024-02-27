package handler_test

import (
	"couplet/internal/controller"
	"couplet/internal/handler"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(controller.Controller{}, nil)
	assert.Empty(t, h)

	h = handler.NewHandler(controller.Controller{}, slog.Default())
	assert.NotEmpty(t, h)
}
