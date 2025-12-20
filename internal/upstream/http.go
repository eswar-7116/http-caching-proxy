package upstream

import (
	"bytes"
	"io"
	"maps"
	"net/http"
	"time"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func Fetch(w http.ResponseWriter, url string) (*cache.Entry, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	maps.Copy(w.Header(), resp.Header)
	w.Header().Set("X-Cache", "MISS")
	w.WriteHeader(resp.StatusCode)

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)
	io.Copy(w, tee)

	return &cache.Entry{
		URL:        url,
		Headers:    resp.Header.Clone(),
		StatusCode: resp.StatusCode,
		Response:   buf.Bytes(),
	}, nil
}
