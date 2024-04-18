package user_test

import (
	"couplet/internal/database/url_slice"
	"couplet/internal/database/user"
	"couplet/internal/database/user_id"
	"couplet/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserBeforeCreate(t *testing.T) {
	noIdUser := user.User{
		ID:        user_id.UserID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		FirstName: "First",
		LastName:  "Last",
		Age:       21,
		Images:    url_slice.UrlSlice{util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png"), util.MustParseUrl("https://example.com/image.png")},
	}
	require.Nil(t, (&noIdUser).BeforeCreate(nil))
	assert.NotEmpty(t, noIdUser.ID)
	id := noIdUser.ID

	require.Nil(t, (&noIdUser).BeforeCreate(nil))
	assert.Equal(t, id, noIdUser.ID)
}
