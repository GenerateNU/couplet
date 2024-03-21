package url_slice

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// A slice of URLs that can be easily stored in a database
type UrlSlice []url.URL

// Wraps a slice of URLs in a UrlSlice to provide additional functionality
func Wrap(s []url.URL) UrlSlice {
	return UrlSlice(s)
}

// Extracts the base slice of URLs from a UrlSlice
func (s UrlSlice) Unwrap() []url.URL {
	return []url.URL(s)
}

// Implements sql.Scanner to read from SQL databases
func (s *UrlSlice) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		return nil
	case []byte:
		return s.Scan(string(src))
	case string:
		*s = UrlSlice{}
		strings := strings.Split(src, " ")
		for _, v := range strings {
			if v != "" {
				u, err := url.Parse(v)
				if err != nil {
					return err
				}
				*s = append(*s, *u)
			}
		}
		return nil
	}
	return fmt.Errorf("unable to scan type %T into UrlSlice", src)
}

// Implements driver.Valuer to write to SQL databases
func (s UrlSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return s.String(), nil
}

// Implements json.Unmarshaler to read from JSON
func (s *UrlSlice) UnmarshalJSON(data []byte) error {
	var strSlice []string
	err := json.Unmarshal(data, &strSlice)
	if err != nil {
		return err
	}
	if strSlice == nil {
		*s = nil
		return nil
	}
	*s = UrlSlice{}
	for _, v := range strSlice {
		if v != "" {
			u, err := url.Parse(v)
			if err != nil {
				return err
			}
			*s = append(*s, *u)
		}
	}
	return nil
}

// Implements json.Marshaler to write to JSON
func (s UrlSlice) MarshalJSON() ([]byte, error) {
	strSlice := []string{}
	for _, v := range s {
		strSlice = append(strSlice, v.String())
	}
	return json.Marshal(strSlice)
}

// Returns the slice as a space-separated list of raw URLs
func (s UrlSlice) String() string {
	b := strings.Builder{}
	for i, v := range s {
		if i != 0 {
			b.WriteString(" ")
		}
		b.WriteString(v.String())
	}
	return b.String()
}
