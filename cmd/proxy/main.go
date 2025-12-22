package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eswar-7116/http-caching-proxy/internal/cache"
	"github.com/eswar-7116/http-caching-proxy/internal/server"
)

const PORT = "8000"

func main() {
	srv := &http.Server{
		Addr: ":" + PORT,
	}

	handler := server.Server{
		Cache: cache.New(300 * time.Second),
	}

	http.HandleFunc("/", handler.Handler)

	go func() {
		fmt.Println("Server started at port", PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error: ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	fmt.Println("\nServer gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown error: ", err)
	}

	fmt.Println("Server shutdown successfully")
}
