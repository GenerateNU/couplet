package controller_test

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	"couplet/internal/database"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestGetUserById(t *testing.T) {

	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)

	uuid1 := uuid.New()
	time1 := time.Now()

	user1 := database.User{
		ID:        uuid1,
		CreatedAt: time1,
		UpdatedAt: time1,
		FirstName: "Stone",
		LastName:  "Liu",
		Age:       20,
	}
	//Insert the user into the database
	tx := db.Create(&user1)
	//Gets the Stone user from the database
	expectUserQuery(mock, user1)
	databaseUser, _ := c.GetUserById(context.Background(), api.GetUserByIdParams{UserId: uuid1})
	expectUserQuery(mock, user1)
	databaseUserAgain, _ := c.GetUserById(context.Background(), api.GetUserByIdParams{UserId: uuid1})

	if tx.Error != nil && err != nil {
		fmt.Println(tx.Error, err)
	} else {
		assert.Equal(t, databaseUser.FirstName, "Stone")
		assert.Equal(t, databaseUser.LastName, "Liu")
		assert.Equal(t, databaseUser.CreatedAt, time1)
		assert.Equal(t, databaseUser.UpdatedAt, time1)
		assert.Equal(t, databaseUser.ID, uuid1)

		//Ensure that multiple calls to the get method returns the same result
		assert.Equal(t, databaseUserAgain.FirstName, "Stone")
		assert.Equal(t, databaseUserAgain.LastName, "Liu")
		assert.Equal(t, databaseUserAgain.CreatedAt, time1)
		assert.Equal(t, databaseUserAgain.UpdatedAt, time1)
		assert.Equal(t, databaseUserAgain.ID, uuid1)
	}
}

func expectUserQuery(mock sqlmock.Sqlmock, user database.User) {
	mock.ExpectQuery("^SELECT \\* FROM \"users\" WHERE id = \\$1 ORDER BY \"users\".\"id\" LIMIT 1").
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.FirstName, user.LastName, user.Age))
}
