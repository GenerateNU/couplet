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
		
	assert.NotEmpty(t, c)
	assert.Nil(t, err)
	assert.NotNil(t, mock)
}

func TestCreateUser(t *testing.T) {
	c, err := controller.NewController(nil)
	assert.Empty(t, c)
	assert.NotNil(t, err)

	// db, mock := database.NewMockDB()
	// c, err = controller.NewController(db)

	    // Create a request
    // req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(reqBytes))
    // if err != nil {
    //     t.Fatalf("Failed to create request: %v", err)
    // }

    // // Create a ResponseRecorder to record the response
    // rr := httptest.NewRecorder()

    // // Call the CreateUser handler
    // handler := http.HandlerFunc(c.CreateUser)
    // handler.ServeHTTP(rr, req)

    // // Check the status code
    // assert.Equal(t, http.StatusOK, rr.Code)

    // // Unmarshal the response body
    // var resBody map[string]interface{}
    // err = json.Unmarshal(rr.Body.Bytes(), &resBody)
    // if err != nil {
    //     t.Fatalf("Failed to unmarshal response body: %v", err)
    // }

    // // Check the response body
    // assert.Equal(t, "John", resBody["FirstName"])
    // assert.Equal(t, "Doe", resBody["LastName"])
    // assert.Equal(t, 25, resBody["Age"])

	// firstName, lastName, age := "John", "Doe", 25

	// mock.ExpectBegin()
	// mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO "users" ("id","created_at","updated_at","first_name","last_name","age") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
	// 	WithArgs(firstName, lastName, age)
	// mock.ExpectCommit()

	// assert.NotEmpty(t, c)
	// assert.Nil(t, err)

	// user, err := c.CreateUser(context.Background(), "John", "Doe", 25)
	
	// assert.NotEmpty(t, user)


	// assert.Equal(t, "John", (*user).FirstName)
	// assert.Equal(t, "Doe", (*user).LastName)
	// assert.Equal(t, uint8(25), (*user).FirstName)
}