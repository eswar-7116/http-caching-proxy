package cache

import "net/http"

type Entry struct {
	URL        string
	Headers    http.Header
	StatusCode int
	Response   []byte
}
