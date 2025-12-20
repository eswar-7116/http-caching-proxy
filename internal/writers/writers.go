package writers

import (
	"maps"
	"net/http"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
)

func WriteCachedResponse(w http.ResponseWriter, entry cache.Entry) {
	maps.Copy(w.Header(), entry.Headers)
	w.Header().Set("X-Cache", "HIT")
	w.WriteHeader(entry.StatusCode)
	w.Write(entry.Response)
}
