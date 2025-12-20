package util

import "net/url"

func IsValidHTTPURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}
