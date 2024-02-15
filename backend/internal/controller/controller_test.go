package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"log/slog"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	c, err := controller.NewController(nil, nil)
	assert.Empty(t, c)
	assert.NotNil(t, err)

	db, _ := database.NewMockDB()
	c, err = controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	c, err = controller.NewController(db, slog.Default())
	assert.NotEmpty(t, c)
	assert.Nil(t, err)
}
