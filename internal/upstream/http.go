package upstream

import (
	"bytes"
	"io"
	"maps"
	"net/http"
	"strings"
	"time"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func Fetch(w http.ResponseWriter, url string, headers http.Header) (*cache.Entry, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for _, h := range []string{
		"User-Agent",
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		"Authorization",
		"Cookie",
	} {
		if v := headers.Get(h); v != "" {
			req.Header.Set(h, v)
		}
	}

	for k, vv := range headers {
		if strings.HasPrefix(http.CanonicalHeaderKey(k), "X-") {
			for _, v := range vv {
				req.Header.Add(k, v)
			}
		}
	}

	resp, err := client.Do(req)
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
