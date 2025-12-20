package upstream

import (
	"io"
	"net/http"
	"time"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func Fetch(url string) (*cache.Entry, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &cache.Entry{
		URL:        url,
		Headers:    resp.Header.Clone(),
		StatusCode: resp.StatusCode,
		Response:   body,
	}, nil
}
