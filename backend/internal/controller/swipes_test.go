package controller_test

// func TestCreateEventSwipe(t *testing.T) {
// 	// set up mock database
// 	db, mock := database.NewMockDB()
// 	c, err := controller.NewController(db, nil)
// 	assert.NotEmpty(t, c)
// 	assert.Nil(t, err)

// 	// set up example event data
// 	event_id := event_id.Wrap(uuid.New())
// 	user_id := user_id.Wrap(uuid.New())

// 	exampleEventSwipe := user.EventSwipe{
// 		UserID:  user_id,
// 		EventID: event_id,
// 		Liked:   true,
// 	}

// 	// expect the insert statement and create the event
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO "event_swipes" ("user_id","event_id","liked","created_at","updated_at") VALUES ($1,$2,$3,$4,$5)`)).
// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	insertedEventSwipe, valErr, txErr := c.CreateEventSwipe(exampleEventSwipe)
// 	assert.Nil(t, valErr)
// 	assert.Nil(t, txErr)

// 	// ensure that all fields were set properly on the Event object
// 	assert.Equal(t, insertedEventSwipe.UserID, exampleEventSwipe.UserID)
// 	assert.Equal(t, insertedEventSwipe.EventID, exampleEventSwipe.EventID)
// 	assert.Equal(t, insertedEventSwipe.Liked, exampleEventSwipe.Liked)

// 	/* TODO:
// 	- test that the same user can have multiple swipes for different events
// 	- test that the same event can have multiple swipes from different users */

// 	// ensure that all expectations are met in the mock
// 	errExpectations := mock.ExpectationsWereMet()
// 	assert.Nil(t, errExpectations)
// }

// func TestCreateUserSwipe(t *testing.T) {
// 	// set up mock database
// 	db, mock := database.NewMockDB()
// 	c, err := controller.NewController(db, nil)
// 	assert.NotEmpty(t, c)
// 	assert.Nil(t, err)

// 	// set up recorder to keep track of the auto-generated eventID
// 	rec := dbtesting.NewValueRecorder()

// 	// set up example event data
// 	other_user_id := user_id.Wrap(uuid.New())
// 	user_id := user_id.Wrap(uuid.New())

// 	// set up example event data
// 	exampleUserSwipe := user.UserSwipe{
// 		UserID:      user_id,
// 		OtherUserID: other_user_id,
// 		Liked:       true,
// 	}

// 	// expect the insert statement and create the event
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO "user_swipes" ("id","user_id","other_user_id","liked","created_at","updated_at")
// 		VALUES ($1,$2,$3,$4,$5,$6)`)).
// 		WithArgs(rec.Record("idOne"), exampleUserSwipe.UserID, exampleUserSwipe.OtherUserID, exampleUserSwipe.Liked, sqlmock.AnyArg(), sqlmock.AnyArg()).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	insertedUserSwipe, valErr, txErr := c.CreateUserSwipe(exampleUserSwipe)
// 	assert.Nil(t, valErr)
// 	assert.Nil(t, txErr)

// 	// ensure that all fields were set properly on the Event object
// 	assert.Equal(t, insertedUserSwipe.UserID, exampleUserSwipe.UserID)
// 	assert.Equal(t, insertedUserSwipe.OtherUserID, exampleUserSwipe.OtherUserID)
// 	assert.Equal(t, insertedUserSwipe.Liked, exampleUserSwipe.Liked)

// }
