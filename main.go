package main

import (
	"io"
	"log"
	"maps"
	"net/http"
	"time"
)

type Entry struct {
	URL        string
	Headers    http.Header
	StatusCode int
	Response   []byte
}

var entries = make(map[string]Entry)

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		log.Println("Got an empty URL parameter")
		w.Write([]byte("URL must not be empty!"))
		return
	}

	if entry, ok := entries[url]; ok {
		maps.Copy(w.Header(), entry.Headers)
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(entry.StatusCode)
		w.Write(append(entry.Response, []byte("FROM CACHE")...))
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("ERROR while invoking a GET request to "+url+"\n", err)
		w.Write([]byte("ERROR while invoking a GET request to " + url))
		return
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR reading response from "+url+"\n", err)
		w.Write([]byte("ERROR reading response from " + url))
		return
	}
	defer resp.Body.Close()

	entries[url] = Entry{
		URL:        url,
		Headers:    r.Header.Clone(),
		StatusCode: resp.StatusCode,
		Response:   respBytes,
	}

	maps.Copy(w.Header(), resp.Header)
	w.Header().Add("X-Cache", "MISS")
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/fetch", handler)
	http.ListenAndServe(":8000", nil)
}
