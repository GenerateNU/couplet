package url_slice_test

import (
	"couplet/internal/database/url_slice"
	"couplet/internal/util"
	"database/sql/driver"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapAndUnwrap(t *testing.T) {
	testCases := []struct {
		input []url.URL
	}{{nil},
		{[]url.URL{}},
		{[]url.URL{util.MustParseUrl("example.com")}},
		{[]url.URL{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}}}
	for _, v := range testCases {
		assert.Equal(t, v.input, url_slice.Wrap(v.input).Unwrap())
	}
}

func TestScan(t *testing.T) {
	validTestCases := []struct {
		input  any
		output url_slice.UrlSlice
	}{{nil, nil},
		{"", url_slice.UrlSlice{}},
		{[]byte(""), url_slice.UrlSlice{}},
		{"example.com", url_slice.UrlSlice{util.MustParseUrl("example.com")}},
		{[]byte("example.com"), url_slice.UrlSlice{util.MustParseUrl("example.com")}},
		{"example.com example.com", url_slice.UrlSlice{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}},
		{[]byte("example.com example.com"), url_slice.UrlSlice{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}}}
	invalidTestCases := []struct {
		input any
	}{{1}, {'t'}, {[]byte{0x7f}}, {util.MustParseUrl("example.com")}}

	for _, v := range validTestCases {
		var s url_slice.UrlSlice
		err := s.Scan(v.input)
		assert.Nil(t, err)
		assert.Equal(t, v.output, s)
	}

	for _, v := range invalidTestCases {
		var s url_slice.UrlSlice
		err := s.Scan(v.input)
		assert.NotNil(t, err)
	}
}

func TestValue(t *testing.T) {
	testCases := []struct {
		input  url_slice.UrlSlice
		output driver.Value
	}{{nil, nil},
		{url_slice.UrlSlice{}, ""},
		{url_slice.UrlSlice{util.MustParseUrl("example.com")}, "example.com"},
		{url_slice.UrlSlice{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}, "example.com example.com"}}

	for _, v := range testCases {
		val, err := v.input.Value()
		assert.Nil(t, err)
		assert.Equal(t, v.output, val)
		if v.input != nil {
			_, isString := val.(string)
			assert.True(t, isString)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	validTestCases := []struct {
		input  []byte
		output url_slice.UrlSlice
	}{{[]byte("null"), nil},
		{[]byte("[]"), url_slice.UrlSlice{}},
		{[]byte("[\"example.com\"]"), url_slice.UrlSlice{util.MustParseUrl("example.com")}},
		{[]byte("[\"example.com\", \"example.com\"]"), url_slice.UrlSlice{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}}}
	invalidTestCases := []struct {
		input []byte
	}{{[]byte("\"example.com\"")},
		{[]byte("1")},
		{[]byte("true")},
		{[]byte("[1]")},
		{[]byte("[true, true, false]")},
		{[]byte("[\"%ZZ\"]")}}

	for _, v := range validTestCases {
		var s url_slice.UrlSlice
		err := json.Unmarshal(v.input, &s)
		assert.Nil(t, err)
		assert.Equal(t, v.output, s)
	}

	for _, v := range invalidTestCases {
		var s url_slice.UrlSlice
		err := json.Unmarshal(v.input, &s)
		assert.NotNil(t, err)
	}
}

func TestMarshal(t *testing.T) {
	testCases := []struct {
		input  url_slice.UrlSlice
		output string
	}{{url_slice.UrlSlice{}, "[]"},
		{url_slice.UrlSlice{util.MustParseUrl("example.com")}, "[\"example.com\"]"},
		{url_slice.UrlSlice{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}, "[\"example.com\",\"example.com\"]"}}

	for _, v := range testCases {
		rawJson, err := json.Marshal(v.input)
		assert.Nil(t, err)
		assert.Equal(t, v.output, string(rawJson))
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		input  url_slice.UrlSlice
		output string
	}{{nil, ""},
		{url_slice.UrlSlice{}, ""},
		{url_slice.UrlSlice{util.MustParseUrl("example.com")}, "example.com"},
		{url_slice.UrlSlice{util.MustParseUrl("example.com"), util.MustParseUrl("example.com")}, "example.com example.com"}}

	for _, v := range testCases {
		s := v.input.String()
		assert.Equal(t, v.output, s)
	}
}
