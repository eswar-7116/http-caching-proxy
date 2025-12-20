package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
	"github.com/eswar-7116/http-caching-proxy/internal/server"
)

const PORT = "8000"

func main() {
	server := server.Server{
		Cache: cache.New(300 * time.Second),
	}

	http.HandleFunc("/", server.Handler)

	fmt.Println("Server started at port", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
