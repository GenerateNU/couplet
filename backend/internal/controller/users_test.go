package controller_test

import (
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arsham/dbtools/dbtesting"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	db, mock := database.NewMockDB()
	c, _ := controller.NewController(db, nil)
	uuid1 := uuid.New()
	db_UUID := user_id.Wrap(uuid1)
	time1 := time.Now()

	user1 := user.User{
		ID:        db_UUID,
		CreatedAt: time1,
		UpdatedAt: time1,
		FirstName: "Stone",
		LastName:  "Liu",
		Age:       20,
		Images:    []user.UserImage{{Url: "https://example.com/image.png"}},
	}
	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "users"`).
		WithArgs(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT 1`).
		WithArgs(db_UUID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age))
	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT 1`).
		WithArgs(db_UUID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(user1.ID, user1.CreatedAt, user1.UpdatedAt, user1.FirstName, user1.LastName, user1.Age))
	mock.ExpectRollback()
	//Insert the user into the database
	tx := db.Create(&user1)

	db_user, db_error := c.GetUser(db_UUID)
	db_user1, _ := c.GetUser(db_UUID)

	if tx.Error != nil && db_error != nil {
		fmt.Println("Error Hit")
		fmt.Println(tx, mock, c)
	} else {
		assert.Equal(t, "Stone", db_user.FirstName)
		assert.Equal(t, "Liu", db_user.LastName)
		assert.Equal(t, time1, db_user.CreatedAt)
		assert.Equal(t, time1, db_user.UpdatedAt)
		assert.Equal(t, db_UUID, db_user.ID)

		assert.Equal(t, "Stone", db_user1.FirstName)
		assert.Equal(t, "Liu", db_user1.LastName)
		assert.Equal(t, time1, db_user1.CreatedAt)
		assert.Equal(t, time1, db_user1.UpdatedAt)
		assert.Equal(t, db_UUID, db_user1.ID)
	}
}

func TestPartialUpdateUser(t *testing.T) {
	db, _ := database.NewMockDB()
	c, err := controller.NewController(db, nil)

	uuid1 := uuid.New()
	db_UUID := user_id.Wrap(uuid1)
	time1 := time.Now()

	user1 := user.User{
		ID:        db_UUID,
		CreatedAt: time1,
		UpdatedAt: time1,
		FirstName: "Stone",
		LastName:  "Liu",
		Age:       20,
		Images:    []user.UserImage{{Url: "https://example.com/image.png"}},
	}
	//Insert the user into the database
	tx := db.Create(&user1)

	//Gets the Stone user from the database
	requestUser := user.User{
		ID:        db_UUID,
		FirstName: "Rock",
		LastName:  "Johnson",
		Age:       uint8(99),
		Images:    []user.UserImage{{Url: "https://example.com/image.png"}},
	}
	databaseUser, _, _ := c.UpdateUser(user1)
	databaseUser1, _, _ := c.UpdateUser(requestUser)

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

func TestPutUser(t *testing.T) {
	// Database Setup
	db, mock := database.NewMockDB()
	c, err := controller.NewController(db, nil)
	assert.Nil(t, err)

	// Create New User Using POST
	firstName, lastName := "John", "Doe"
	var age uint8 = 20
	images := []user.UserImage{{Url: "https://example.com/image.png"}}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
        INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
        VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "https://example.com/image.png", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// Insert User into database
	createUser, err := c.CreateUser(firstName, lastName, age, images)
	assert.Nil(t, err)

	assert.Equal(t, firstName, createUser.FirstName)
	assert.Equal(t, lastName, createUser.LastName)
	assert.Equal(t, age, createUser.Age)

	// // Get User ID to Update User
	newUserID := createUser.ID

	updatedFirstName, updatedLastName := "Jane", "Smith"
	var updatedAge uint8 = 20

	putRequestBody := user.User{
		FirstName: updatedFirstName,
		LastName:  updatedLastName,
		Age:       updatedAge,
		Images:    images,
	}

	// Retrieve the User and Update the User
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT \\* FROM \"users\" WHERE id = \\$1 ORDER BY \"users\".\"id\" LIMIT 1").
		WithArgs(newUserID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(newUserID, createUser.CreatedAt, createUser.UpdatedAt, createUser.FirstName, createUser.LastName, createUser.Age))
	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE "users"
		SET "age"=$1,"first_name"=$2,"last_name"=$3,"updated_at"=$4 WHERE "id" = $5`)).
		WithArgs(updatedAge, updatedFirstName, updatedLastName, sqlmock.AnyArg(), newUserID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id","id") VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "https://example.com/image.png", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	putUser, err := c.SaveUser(putRequestBody, newUserID)
	assert.Nil(t, err)

	assert.Equal(t, updatedFirstName, putUser.FirstName)
	assert.Equal(t, updatedLastName, putUser.LastName)
	assert.Equal(t, updatedAge, putUser.Age)
	assert.Equal(t, createUser.CreatedAt, putUser.CreatedAt)
	assert.True(t, putUser.UpdatedAt.After(createUser.UpdatedAt))
	assert.Equal(t, createUser.ID, putUser.ID)
}

func TestCreateUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()
	// logger := slog.New(pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug)))
	c, err := controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID
	rec := dbtesting.NewValueRecorder()

	// set up user data
	firstName := "John"
	lastName := "Smith"
	var age uint8 = 20

	images := []user.UserImage{{Url: "https://example.com/image.png"}}

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "https://example.com/image.png", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	newUser1, err := c.CreateUser(firstName, lastName, age, images)
	assert.Nil(t, err)

	// ensure that all fields were set properly on the User object
	assert.Equal(t, newUser1.Age, age)
	assert.Equal(t, newUser1.FirstName, firstName)
	assert.Equal(t, newUser1.LastName, lastName)

	// create a second user with the same data to show that repeated POST calls always creates new users
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("newUserId"), sqlmock.AnyArg(), sqlmock.AnyArg(), firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id","id") VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "https://example.com/image.png", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	newUser2, err := c.CreateUser(firstName, lastName, age, images)
	assert.Nil(t, err)

	assert.Equal(t, newUser2.Age, age)
	assert.Equal(t, newUser2.FirstName, firstName)
	assert.Equal(t, newUser2.LastName, lastName)

	// IMPORTANT! assert that internally, the second user is not the same as the first user
	assert.NotEqual(t, newUser2.ID, newUser1.ID)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}

func TestDeleteUser(t *testing.T) {
	// set up mock database
	db, mock := database.NewMockDB()

	c, err := controller.NewController(db, nil)
	assert.NotEmpty(t, c)
	assert.Nil(t, err)

	// set up recorder to keep track of the auto-generated userID and created/updated times
	rec := dbtesting.NewValueRecorder()

	// set up user data
	firstName := "firstName"
	lastName := "lastName"
	var age uint8 = 20
	images := []user.UserImage{{Url: "https://example.com/image.png"}}

	// expect the insert statement and create the user
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age")
		VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(rec.Record("id"), rec.Record("createdTime"), rec.Record("updatedTime"), "firstName", "lastName", age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO "user_images" ("created_at","updated_at","url","user_id") VALUES ($1,$2,$3,$4)
		ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "https://example.com/image.png", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectCommit()

	_, err = c.CreateUser("firstName", "lastName", age, images)
	assert.Nil(t, err)

	// retrieve the user's ID
	userId := rec.Value("id").(string)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"}).
			AddRow(userId, rec.Value("createdTime").(time.Time), rec.Value("updatedTime").(time.Time), "firstName", "lastName", 20))

	// expect the delete statement and delete the user
	mock.ExpectExec(regexp.QuoteMeta(`
		DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	uuidUser, err := uuid.Parse(userId)
	assert.Nil(t, err)

	deletedUser, err := c.DeleteUser(user_id.Wrap(uuidUser))
	assert.Nil(t, err)

	// ensure that the deleted user is returned and matches the info of the user that was created
	assert.Equal(t, deletedUser.Age, age)
	assert.Equal(t, deletedUser.FirstName, firstName)
	assert.Equal(t, deletedUser.LastName, lastName)

	// try deleting a fake user
	badId := uuid.New()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(badId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "first_name", "last_name", "age"})) // no rows added

	deletedUser, err = c.DeleteUser(user_id.Wrap(badId))
	assert.Error(t, err)

	// ensure that all expectations are met in the mock
	errExpectations := mock.ExpectationsWereMet()
	assert.Nil(t, errExpectations)
}
