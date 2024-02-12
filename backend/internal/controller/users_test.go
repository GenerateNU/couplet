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

func TestCreateUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	// logger := slog.New(pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug)))
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID
	rec := dbtesting.NewValueRecorder()

	// set up user data
	firstName := "John"
	lastName := "Smith"
	var age uint8 = 20

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user, err := c.CreateUser(firstName, lastName, age)
	require.Nil(t, err)

	// ensure that all fields were set properly on the User object
	require.Equal(t, user.Age, age)
	require.Equal(t, user.FirstName, firstName)
	require.Equal(t, user.LastName, lastName)

	// create a second user with the same data to show that repeated POST calls always creates new users
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("newUserId"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newUser, err := c.CreateUser(firstName, lastName, age)
	require.Nil(t, err)

	require.Equal(t, newUser.Age, age)
	require.Equal(t, newUser.FirstName, firstName)
	require.Equal(t, newUser.LastName, lastName)

	// IMPORTANT! assert that internally, the second user is not the same as the first user
	require.NotEqual(t, newUser.ID, user.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}

func TestDeleteUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()

	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID and created/updated times
	rec := dbtesting.NewValueRecorder()

	// set up user data
	firstName := "firstName"
	lastName := "lastName"
	var age uint8 = 20

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), "firstName", "lastName", age).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	_, err = c.CreateUser("firstName", "lastName", age)
	require.Nil(t, err)

	// retrieve the user's ID
	userId := rec.Value("id").(string)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), "firstName", "lastName", 20))

	mock.ExpectBegin()

	// expect the delete statement and delete the user
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	uuidUser, err := uuid.Parse(userId)
	require.Nil(t, err)

	deletedUser, err := c.DeleteUserById(uuidUser)
	require.Nil(t, err)

	// ensure that the deleted user is returned and matches the info of the user that was created
	require.Equal(t, deletedUser.Age, age)
	require.Equal(t, deletedUser.FirstName, firstName)
	require.Equal(t, deletedUser.LastName, lastName)

	// try deleting a fake user
	badId := uuid.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(badId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"})) // no rows added

	deletedUser, err = c.DeleteUserById(badId)
	require.Error(t, err)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	require.Nil(t, errExpectations)
}
