package controller_test

import (
	"context"
	"couplet/internal/api"
	"couplet/internal/controller"
	database "couplet/internal/database"
	user "couplet/internal/database/user"
	userId "couplet/internal/database/user/id"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	// Create New User Using PUT
	putRequestBody := api.User{}
	putRequestBody.SetFirstName("UserFirstName")
	putRequestBody.SetLastName("UserLastName")
	putRequestBody.SetAge(25)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
        INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
        VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), putRequestBody.FirstName, putRequestBody.LastName, uint8(putRequestBody.Age)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	dummyID := uuid.New()

	// Insert User into database
	createUser, err := c.SaveUserById(context.Background(), &putRequestBody, dummyID.String())
	require.NoError(t, err)

	require.Equal(t, "UserFirstName", createUser.FirstName)
	require.Equal(t, "UserLastName", createUser.LastName)
	require.Equal(t, uint8(25), createUser.Age)

	// Get User ID to Update User
	newUserID := uuid.UUID(createUser.ID).String()

	putRequestBody2 := api.User{}
	putRequestBody2.SetFirstName("UpdatedFirstName")
	putRequestBody2.SetLastName("UpdatedLastName")
	putRequestBody2.SetAge(99)

	// Retrieve the User and Update the User
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT \\* FROM \"users\" WHERE id = \\$1 ORDER BY \"users\".\"id\" LIMIT 1").
		WithArgs(newUserID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(newUserID, createUser.CreatedAt, createUser.UpdatedAt, createUser.FirstName, createUser.LastName, createUser.Age))
	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE "users"
		SET "first_name" = $1, "last_name" = $2, "age" = $3, "updated_at" = $4
		WHERE "id" = $5`)).
		WithArgs("UpdatedFirstName", "UpdatedLastName", uint8(99), sqlmock.AnyArg(), newUserID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	putUser, err := c.SaveUserById(context.Background(), &putRequestBody2, newUserID)
	require.NoError(t, err)

	require.Equal(t, "UpdatedFirstName", putUser.FirstName)
	require.Equal(t, "UpdatedLastName", putUser.LastName)
	require.Equal(t, uint8(99), putUser.Age)
	require.Equal(t, createUser.CreatedAt, putUser.CreatedAt)
	require.True(t, putUser.UpdatedAt.After(createUser.UpdatedAt))
	require.Equal(t, createUser.ID, putUser.ID)
}

func TestGetAllUsers(t *testing.T) {
	// Database Setup
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db)
	require.NotEmpty(t, c)
	require.NoError(t, err)

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

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, uint8(user1.Age)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = db.Create(&user1).Error
	require.NoError(t, err, "Failed to create user 1")

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

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
        INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
        VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(user2.ID, user2.CreatedAt, user2.UpdatedAt, user2.FirstName, user2.LastName, uint8(user2.Age)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = db.Create(&user2).Error
	require.NoError(t, err, "Failed to create user 2")

	// Test the Get All Users Request
	var limit uint8 = 50
	var offset uint32 = 0

	getUsers, err := c.GetAllUsers(limit, offset)
	require.NoError(t, err, "Failed to get all users")

	// Check Both Users and their Fields
	require.Equal(t, 2, len(getUsers))

	require.Equal(t, user1.ID, getUsers[0].ID)
	require.Equal(t, user1.CreatedAt, getUsers[0].CreatedAt)
	require.Equal(t, user1.UpdatedAt, getUsers[0].UpdatedAt)
	require.Equal(t, user1.FirstName, getUsers[0].FirstName)
	require.Equal(t, user1.LastName, getUsers[0].LastName)
	require.Equal(t, user1.Age, getUsers[0].Age)

	require.Equal(t, user2.ID, getUsers[1].ID)
	require.Equal(t, user2.CreatedAt, getUsers[1].CreatedAt)
	require.Equal(t, user2.UpdatedAt, getUsers[1].UpdatedAt)
	require.Equal(t, user2.FirstName, getUsers[1].FirstName)
	require.Equal(t, user2.LastName, getUsers[1].LastName)
	require.Equal(t, user2.Age, getUsers[1].Age)
}

func expectUserQuery(mock sqlmock.Sqlmock, user user.User) {
	mock.ExpectQuery("^SELECT \\* FROM \"users\" WHERE id = \\$1 ORDER BY \"users\".\"id\" LIMIT 1").
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.FirstName, user.LastName, user.Age))
}
