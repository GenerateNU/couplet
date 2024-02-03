package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	c, err := controller.NewController(nil)
	assert.Empty(t, c)
	assert.NotNil(t, err)

	db, mock := database.NewMockDB()

	c, err = controller.NewController(db)
	assert.NotEmpty(t, mock)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)
}
