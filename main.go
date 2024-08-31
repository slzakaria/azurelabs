package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

// Middleware for logging requests and status codes
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &logResponseWriter{w, http.StatusOK}

		next.ServeHTTP(lrw, r)

		log.Printf("Method: %s | Route: %s | Status: %d | Duration: %s",
			r.Method, r.URL.Path, lrw.statusCode, time.Since(start))
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func main() {
	http.HandleFunc("/photos", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/photos")
		if err != nil {
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			log.Printf("Error: %v", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response", http.StatusInternalServerError)
			log.Printf("Error: %v", err)
			return
		}

		var photos []Photo
		if err := json.Unmarshal(body, &photos); err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
			log.Printf("Error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photos)
	})

	loggedRouter := loggingMiddleware(http.DefaultServeMux)

	log.Println("Server is running on port 3000...")
	if err := http.ListenAndServe(":3000", loggedRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
