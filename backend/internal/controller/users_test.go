package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"

	"testing"

	"context"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	firstName := "firstName"
	lastName := "lastName"
	age := 20

	mock.ExpectBegin()
	rec := dbtesting.NewValueRecorder()

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), "firstName", "lastName", age).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// this expects the default transaction to commit
	user, err := c.CreateUser(context.Background(), "firstName", "lastName", 20)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the User object
	assert.Equal(t, user.Age, uint8(age))
	assert.Equal(t, user.FirstName, firstName)
	assert.Equal(t, user.LastName, lastName)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}
