package controller_test

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	database "couplet/internal/database"
	user "couplet/internal/database/user"
	userId "couplet/internal/database/user/id"
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
	db_UUID := userId.UserID(uuid1)
	time1 := time.Now()

	user1 := user.User{
		ID:        db_UUID,
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
	databaseUser, e := c.GetUserById(context.Background(), api.GetUserByIdParams{UserId: uuid1})
	expectUserQuery(mock, user1)
	databaseUserAgain, _ := c.GetUserById(context.Background(), api.GetUserByIdParams{UserId: uuid1})

	if tx.Error != nil && err != nil && e != nil {
		fmt.Println("Error Hit")
	} else {
		assert.Equal(t, "Stone", databaseUser.FirstName)
		assert.Equal(t, "Liu", databaseUser.LastName)
		assert.Equal(t, time1, databaseUser.CreatedAt)
		assert.Equal(t, time1, databaseUser.UpdatedAt)
		assert.Equal(t, db_UUID, databaseUser.ID)

		//Ensure that multiple calls to the get method returns the same result
		assert.Equal(t, "Stone", databaseUserAgain.FirstName)
		assert.Equal(t, "Liu", databaseUserAgain.LastName)
		assert.Equal(t, time1, databaseUserAgain.CreatedAt)
		assert.Equal(t, time1, databaseUserAgain.UpdatedAt)
		assert.Equal(t, db_UUID, databaseUserAgain.ID)

	}
}

func TestPartialUpdateUserById(t *testing.T) {
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)

	uuid1 := uuid.New()
	db_UUID := userId.UserID(uuid1)
	time1 := time.Now()

	user1 := user.User{
		ID:        db_UUID,
		CreatedAt: time1,
		UpdatedAt: time1,
		FirstName: "Stone",
		LastName:  "Liu",
		Age:       20,
	}
	//Insert the user into the database
	tx := db.Create(&user1)
	//Gets the Stone user from the database
	requestUUID := api.PartialUpdateUserByIdParams{
		UserId: uuid1,
	}
	requestName := api.PartialUpdateUserByIdParams{
		UserId:    uuid1,
		FirstName: api.NewOptString("Rock"),
		LastName:  api.NewOptString("Johnson"),
		Age:       api.NewOptInt(99),
	}
	expectUserQuery(mock, user1)
	databaseUser, _ := c.PartialUpdateUserById(context.Background(), requestUUID)
	expectUserQuery(mock, user1)
	databaseUser1, _ := c.PartialUpdateUserById(context.Background(), requestName)

	if tx.Error != nil && err != nil {
		fmt.Println("Error Has Occured")
	} else {
		//Nothing should have changed
		assert.Equal(t, "Stone", databaseUser.FirstName)
		assert.Equal(t, "Liu", databaseUser.LastName)
		assert.Equal(t, time1, databaseUser.CreatedAt)
		assert.Equal(t, time1, databaseUser.UpdatedAt)
		assert.Equal(t, db_UUID, databaseUser.ID)
		assert.Equal(t, uint8(20), databaseUser.Age)
		//Update the First and Last Name
		assert.Equal(t, "Rock", databaseUser1.FirstName)
		assert.Equal(t, "Johnson", databaseUser1.LastName)
		assert.Equal(t, db_UUID, databaseUser1.ID)
		assert.Equal(t, uint8(99), databaseUser1.Age)

	}
}

func expectUserQuery(mock sqlmock.Sqlmock, user user.User) {
	mock.ExpectQuery("^SELECT \\* FROM \"users\" WHERE id = \\$1 ORDER BY \"users\".\"id\" LIMIT 1").
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.FirstName, user.LastName, user.Age))
}
