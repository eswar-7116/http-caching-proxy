package server

import (
	"log"
	"net/http"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
	"github.com/eswar-7116/http-caching-proxy/internal/upstream"
	"github.com/eswar-7116/http-caching-proxy/internal/writers"
)

type Server struct {
	Cache *cache.Cache
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		log.Println("Got an empty URL parameter")
		w.Write([]byte("URL must not be empty!"))
		return
	}

	if entry, ok := s.Cache.Get(url); ok {
		writers.WriteCachedResponse(w, entry)
		return
	}

	entry, err := upstream.Fetch(url)
	if err != nil {
		log.Printf("ERROR while fetching upstream for '%s': %s", url, err.Error())
		w.Write([]byte("ERROR while fetching upstream for " + url))
		return
	}

	s.Cache.Set(url, *entry)
	writers.WriteUpstreamResponse(w, entry)
}
