package server

import (
	"log"
	"net/http"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
	"github.com/eswar-7116/http-caching-proxy/internal/upstream"
	"github.com/eswar-7116/http-caching-proxy/internal/util"
	"github.com/eswar-7116/http-caching-proxy/internal/writers"
)

type Server struct {
	Cache *cache.Cache
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" || !util.IsValidHTTPURL(url) {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	if entry, ok := s.Cache.Get(url); ok {
		writers.WriteCachedResponse(w, entry)
		return
	}

	entry, err := upstream.Fetch(w, url)
	if err != nil {
		log.Printf("ERROR while fetching upstream for '%s': %s", url, err.Error())
		http.Error(w, "upstream fetch failed", http.StatusBadGateway)
		return
	}

	s.Cache.Set(url, *entry)
}
