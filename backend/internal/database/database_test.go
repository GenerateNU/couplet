package database_test

import (
	"couplet/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDBwithNoDatabaseRunning(t *testing.T) {
	db, err := database.NewDB("", 0, "", "", "", nil)
	assert.Empty(t, db)
	assert.NotNil(t, err)
}

func TestNewMockDB(t *testing.T) {
	assert.NotPanics(t, func() { database.NewMockDB() })
	db, mock := database.NewMockDB()
	assert.NotEmpty(t, db)
	assert.NotEmpty(t, mock)
}

func TestEnableConnPoolingOnNilDB(t *testing.T) {
	assert.NotEmpty(t, database.EnableConnPooling(nil))
}

func TestMigrateOnNilDB(t *testing.T) {
	assert.NotEmpty(t, database.Migrate(nil))
}
