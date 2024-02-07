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

func TestPutUserById(t *testing.T) {
	// Database Setup
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}

	// Create new User 
	uuid1 := uuid.New()
	db_UUID := userId.UserID(uuid1)
	time1 := time.Now()

	user1 := user.User{
		ID:        db_UUID,
		CreatedAt: time1,
		UpdatedAt: time1,
		FirstName: "UserFirstName",
		LastName:  "UserLastName",
		Age:       25,
	}

	// Insert the user into the database
	err := db.Create(&user1).Error
	assert.NoError(t, tx)
	// Gets and Checks if the Test User is properly created
	var retrievedUser user.User;
	err = db.First(&retrievedUser, "id = ", db_UUID).Error
	assert.NoError(t, err)
	assert.Equal(t, user1.ID, retrievedUser.ID)
	assert.Equal(t, user1.CreatedAt, retrievedUser.CreatedAt)
	assert.Equal(t, user1.UpdatedAt, retrievedUser.UpdatedAt)
	assert.Equal(t, user1.FirstName, retrievedUser.FirstName)
	assert.Equal(t, user1.LastName, retrievedUser.LastName)
	assert.Equal(t, user1.Age, retrievedUser.Age)

	// Test the Put Request
	putRequestParams := api.PutUserByIdParams{
		UserId:    uuid1,
		FirstName: api.NewOptString("PutUserFirstName"),
		LastName:  api.NewOptString("PutUserLastName"),
		Age:       api.NewOptInt(99),
	}

	expectUserQuery(mock, user1)
	updatedUser, err := c.PutUserById(context.Background(), putRequestParams)
	assert.NoError(t, err)
	assert.Equal(t, db_UUID, updatedUser.ID)
	assert.Equal(t, "PutUserFirstName", updatedUser.FirstName)
	assert.Equal(t, "PutUserLastName", updatedUser.LastName)
	assert.Equal(t, uint8(99), updatedUser.Age)
}

func TestGetAllUsers(t *testing.T) {
	// Database Setup
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}

	// Create new User 1 
	uuid1 := uuid.New()
	db_UUID1 := userId.UserID(uuid1)
	time1 := time.Now()
	user1 := user.User{
		ID:        db_UUID1,
		CreatedAt: time1,
		UpdatedAt: time1,
		FirstName: "User1FirstName",
		LastName:  "User1LastName",
		Age:       20,
	}
	err := db.Create(&user1).Error
    assert.NoError(t, err, "Failed to create user 1")

	// Create new User 2
	uuid2 := uuid.New()
	db_UUID2 := userId.UserID(uuid2)
	time2 := time.Now()
	user2 := user.User{
		ID:        db_UUID2,
		CreatedAt: time2,
		UpdatedAt: time2,
		FirstName: "User2FirstName",
		LastName:  "User2LastName",
		Age:       40,
	}
	err := db.Create(&user2).Error
    assert.NoError(t, err, "Failed to create user 2")

    // Test the Get All Users Request
    getUsersRequestParams := api.GetAllUsersParams{
        Limit: 50,
		Offset: 0,
    }
	
    expectUserQuery(mock, user1)
	expectUserQuery(mock, user2)

	getUsers, err := c.GetAllUsers(context.Background(), getUsersRequestParams)

	// Check Both Users and their Fields 
	assert.Equal(t, 2, len(getUsers))
    
    assert.Equal(t, user1.ID, getUsers[0].ID)
    assert.Equal(t, user1.CreatedAt, getUsers[0].CreatedAt)
    assert.Equal(t, user1.UpdatedAt, getUsers[0].UpdatedAt)
    assert.Equal(t, user1.FirstName, getUsers[0].FirstName)
    assert.Equal(t, user1.LastName, getUsers[0].LastName)
    assert.Equal(t, user1.Age, getUsers[0].Age)

    assert.Equal(t, user2.ID, getUsers[1].ID)
    assert.Equal(t, user2.CreatedAt, getUsers[1].CreatedAt)
    assert.Equal(t, user2.UpdatedAt, getUsers[1].UpdatedAt)
    assert.Equal(t, user2.FirstName, getUsers[1].FirstName)
    assert.Equal(t, user2.LastName, getUsers[1].LastName)
    assert.Equal(t, user2.Age, getUsers[1].Age)
}

func expectUserQuery(mock sqlmock.Sqlmock, user user.User) {
	mock.ExpectQuery("^SELECT \\* FROM \"users\" WHERE id = \\$1 ORDER BY \"users\".\"id\" LIMIT 1").
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.FirstName, user.LastName, user.Age))
}
