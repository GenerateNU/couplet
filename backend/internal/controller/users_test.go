package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"

	// "fmt"
	// "regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	// "github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PASS
func TestGetUser(t *testing.T) {
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// Mock User Data
	newUserId := user_id.Wrap(uuid.New())
	user1 := user.User{
		ID:        newUserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: "UserFirstName",
		LastName:  "UserLastName",
		Age:       20,
	}

	// Set expectations to create user and fetch user by ID
	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "users"`).
		WithArgs(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT 1`).
		WithArgs(user1.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age))

	// Insert the user into the database
	err = db.Create(&user1).Error
	require.NoError(t, err)

	// Get User and Check Fields
	db_user, db_error := c.GetUser(user1.ID)
	require.NoError(t, db_error)

	require.Equal(t, user1.ID, db_user.ID)
	require.Equal(t, user1.CreatedAt, db_user.CreatedAt)
	require.Equal(t, user1.UpdatedAt, db_user.UpdatedAt)
	require.Equal(t, user1.FirstName, db_user.FirstName)
	require.Equal(t, user1.LastName, db_user.LastName)
	require.Equal(t, user1.Age, db_user.Age)
}

// PASS
func TestGetUsers(t *testing.T) {
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// Mock First User Data
	userId1 := user_id.Wrap(uuid.New())
	user1 := user.User{
		ID:        userId1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: "User1FirstName",
		LastName:  "User1LastName",
		Age:       23,
	}

	// Mock Second User Data
	userId2 := user_id.Wrap(uuid.New())
	user2 := user.User{
		ID:        userId2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: "User2FirstName",
		LastName:  "User2LastName",
		Age:       40,
	}

	// Set expectations to create both users and fetch both users
	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "users"`).
		WithArgs(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "users"`).
		WithArgs(user2.ID, user2.CreatedAt, user2.UpdatedAt, user2.FirstName, user2.LastName, user2.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(`^SELECT \* FROM "users" LIMIT 1 OFFSET 2`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age).
			AddRow(user2.ID, user2.CreatedAt, user2.UpdatedAt, user2.FirstName, user2.LastName, user2.Age))

	// Insert users into the database
	err = db.Create(&user1).Error
	require.NoError(t, err)

	err = db.Create(&user2).Error
	require.NoError(t, err)

	// Get Users and Check Fields
	users, err := c.GetUsers(1, 2)
	require.NoError(t, err)
	require.Len(t, users, 2)

	// Check first user
	require.Equal(t, user1.ID, users[0].ID)
	require.Equal(t, user1.CreatedAt, users[0].CreatedAt)
	require.Equal(t, user1.UpdatedAt, users[0].UpdatedAt)
	require.Equal(t, user1.FirstName, users[0].FirstName)
	require.Equal(t, user1.LastName, users[0].LastName)
	require.Equal(t, user1.Age, users[0].Age)

	// Check second user
	require.Equal(t, user2.ID, users[1].ID)
	require.Equal(t, user2.CreatedAt, users[1].CreatedAt)
	require.Equal(t, user2.UpdatedAt, users[1].UpdatedAt)
	require.Equal(t, user2.FirstName, users[1].FirstName)
	require.Equal(t, user2.LastName, users[1].LastName)
	require.Equal(t, user2.Age, users[1].Age)
}

func TestSaveUser(t *testing.T) {
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	require.NotEmpty(t, c)
	require.Nil(t, err)

	// Mock First User Data
	userId1 := user_id.Wrap(uuid.New())
	user1 := user.User{
		ID:        userId1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: "User1FirstName",
		LastName:  "User1LastName",
		Age:       29,
	}

	// Set Expectation to Create First User
	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "users"`).
		WithArgs(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Insert the user into the database
	err = db.Create(&user1).Error
	require.NoError(t, err)

	// Change User Information Using Put
	putRequestBody := user.User{
		FirstName: "PutFirstName",
		LastName:  "PutLastName",
		Age:       20,
	}
	
	// Expect the Select Statement and Update the User 
	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE id = \$1 ORDER BY "users"\."id" LIMIT 1`).
		WithArgs(userId1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age))
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "created_at"=\$1,"updated_at"=\$2,"first_name"=\$3,"last_name"=\$4,"age"=\$5 WHERE "id" = \$6`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), putRequestBody.FirstName, putRequestBody.LastName, putRequestBody.Age, user1.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Change User Information in database with PUT Endpoint
	updatedUserResult, err := c.SaveUser(putRequestBody, user1.ID)
	require.NoError(t, err)
	require.NotNil(t, updatedUserResult)

	require.Equal(t, user1.ID, updatedUserResult.ID)
	require.Equal(t, user1.CreatedAt, updatedUserResult.CreatedAt)
	require.NotEqual(t, user1.UpdatedAt, updatedUserResult.UpdatedAt)
	require.Equal(t, putRequestBody.FirstName, updatedUserResult.FirstName)
	require.Equal(t, putRequestBody.LastName, updatedUserResult.LastName)
	require.Equal(t, putRequestBody.Age, updatedUserResult.Age)
}

// func TestPartialUpdateUser(t *testing.T) {
// 	db, _ := database.NewMockDB()
// 	c, err := controller.NewController(db, nil)

// 	uuid1 := uuid.New()
// 	db_UUID := user_id.UserID(uuid1)
// 	time1 := time.Now()

// 	user1 := user.User{
// 		ID:        db_UUID,
// 		CreatedAt: time1,
// 		UpdatedAt: time1,
// 		FirstName: "Stone",
// 		LastName:  "Liu",
// 		Age:       20,
// 	}
// 	//Insert the user into the database
// 	tx := db.Create(&user1)

// 	//Gets the Stone user from the database
// 	requestUser := user.User{
// 		ID:        db_UUID,
// 		FirstName: "Rock",
// 		LastName:  "Johnson",
// 		Age:       uint8(99),
// 	}
// 	databaseUser, _, _ := c.UpdateUser(user1)
// 	databaseUser1, _, _ := c.UpdateUser(requestUser)

// 	if tx.Error != nil && err != nil {
// 		fmt.Println("Error Has Occured")
// 	} else {
// 		//Nothing should have changed
// 		assert.Equal(t, "Stone", databaseUser.FirstName)
// 		assert.Equal(t, "Liu", databaseUser.LastName)
// 		assert.Equal(t, time1, databaseUser.CreatedAt)
// 		assert.Equal(t, time1, databaseUser.UpdatedAt)
// 		assert.Equal(t, db_UUID, databaseUser.ID)
// 		assert.Equal(t, uint8(20), databaseUser.Age)
// 		//Update the First and Last Name
// 		assert.Equal(t, "Rock", databaseUser1.FirstName)
// 		assert.Equal(t, "Johnson", databaseUser1.LastName)
// 		assert.Equal(t, db_UUID, databaseUser1.ID)
// 		assert.Equal(t, uint8(99), databaseUser1.Age)
// 	}
// }

// func TestCreateUser(t *testing.T) {
// 	// set up mock database
// 	db, mock := database.NewMockDB()
// 	// logger := slog.New(pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug)))
// 	c, err := controller.NewController(db, nil)
// 	require.NotEmpty(t, c)
// 	require.Nil(t, err)

// 	// set up recorder to keep track of the auto-generated userID
// 	rec := dbtesting.NewValueRecorder()

// 	// set up user data
// 	firstName := "John"
// 	lastName := "Smith"
// 	var age uint8 = 20

// 	// expect the insert statement and create the user
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
// 		VALUES ($1,$2,$3,$4,$5,$6)`)).
// 		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	user, err := c.CreateUser(firstName, lastName, age)
// 	require.Nil(t, err)

// 	// ensure that all fields were set properly on the User object
// 	require.Equal(t, user.Age, age)
// 	require.Equal(t, user.FirstName, firstName)
// 	require.Equal(t, user.LastName, lastName)

// 	// create a second user with the same data to show that repeated POST calls always creates new users
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
// 		VALUES ($1,$2,$3,$4,$5,$6)`)).
// 		WithArgs(rec.Record("newUserId"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	newUser, err := c.CreateUser(firstName, lastName, age)
// 	require.Nil(t, err)

// 	require.Equal(t, newUser.Age, age)
// 	require.Equal(t, newUser.FirstName, firstName)
// 	require.Equal(t, newUser.LastName, lastName)

// 	// IMPORTANT! assert that internally, the second user is not the same as the first user
// 	require.NotEqual(t, newUser.ID, user.ID)

// 	// ensure that all expectations are met in the mock
// 	errExpectations := mock.ExpectationsWereMet()
// 	require.Nil(t, errExpectations)
// }

// func TestDeleteUser(t *testing.T) {
// 	// set up mock database
// 	db, mock := database.NewMockDB()

// 	c, err := controller.NewController(db, nil)
// 	require.NotEmpty(t, c)
// 	require.Nil(t, err)

// 	// set up recorder to keep track of the auto-generated userID and created/updated times
// 	rec := dbtesting.NewValueRecorder()

// 	// set up user data
// 	firstName := "firstName"
// 	lastName := "lastName"
// 	var age uint8 = 20

// 	// expect the insert statement and create the user
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
// 		VALUES ($1,$2,$3,$4,$5,$6)`)).
// 		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), "firstName", "lastName", age).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectCommit()

// 	_, err = c.CreateUser("firstName", "lastName", age)
// 	require.Nil(t, err)

// 	// retrieve the user's ID
// 	userId := rec.Value("id").(string)

// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
// 		WithArgs(userId).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
// 			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), "firstName", "lastName", 20))

// 	mock.ExpectBegin()

// 	// expect the delete statement and delete the user
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		DELETE FROM "users" WHERE "users"."id" = $1`)).
// 		WithArgs(userId).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	uuidUser, err := uuid.Parse(userId)
// 	require.Nil(t, err)

// 	deletedUser, err := c.DeleteUser(user_id.Wrap(uuidUser))
// 	require.Nil(t, err)

// 	// ensure that the deleted user is returned and matches the info of the user that was created
// 	require.Equal(t, deletedUser.Age, age)
// 	require.Equal(t, deletedUser.FirstName, firstName)
// 	require.Equal(t, deletedUser.LastName, lastName)

// 	// try deleting a fake user
// 	badId := uuid.New()
// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
// 		WithArgs(badId).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"})) // no rows added

// 	deletedUser, err = c.DeleteUser(user_id.Wrap(badId))
// 	require.Error(t, err)

// 	// ensure that all expectations are met in the mock
// 	errExpectations := mock.ExpectationsWereMet()
// 	require.Nil(t, errExpectations)
// }
