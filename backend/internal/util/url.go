package util

import "net/url"

// Parses a string as a URL. Panics if parsing fails
func MustParseUrl(rawURL string) url.URL {
	url, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return *url
}
