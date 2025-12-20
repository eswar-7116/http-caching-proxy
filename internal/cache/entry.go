package cache

import (
	"net/http"
	"time"
)

type Entry struct {
	URL        string
	Headers    http.Header
	StatusCode int
	Response   []byte
	ExpiresAt  time.Time
}
