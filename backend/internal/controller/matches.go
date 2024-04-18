package controller

import (
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"couplet/internal/database/user_match"
	"fmt"
	"time"
)

type MatchData struct {
	ID        user_id.UserID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Age       uint8
	Bio       string
	Images    url_slice.UrlSlice
	Viewed    bool
}

// Gets a specified user's matches from the database
func (c Controller) GetUserMatches(id user_id.UserID) (matchData []MatchData, txErr error) {
	var user user.User
	txErr = c.database.First(&user, id).Error
	if txErr != nil {
		return
	}

	var matches []user_match.UserMatch
	txErr = c.database.Model(&user).Association("Matches").Find(&matches)
	if txErr != nil {
		return
	}

	fmt.Println(matches)

	for _, match := range matches {
		matchUser, err := c.GetUser(match.MatchID)
		if err != nil {
			txErr = err
			return
		}
		fmt.Println(match)

		matchData = append(matchData, MatchData{
			ID:        match.MatchID,
			CreatedAt: match.CreatedAt,
			UpdatedAt: match.UpdatedAt,
			FirstName: matchUser.FirstName,
			LastName:  matchUser.LastName,
			Age:       matchUser.Age,
			Bio:       matchUser.Bio,
			Images:    matchUser.Images,
			Viewed:    match.Viewed,
		})
	}
	return
}
