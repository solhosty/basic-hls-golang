package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// so far allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// so far allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")
		// so far allow specific methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		// handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	hlsDir := "./videos"
	if _, err := os.Stat(hlsDir); os.IsNotExist(err) {
		log.Fatalf("HLS directory %s does not exist", hlsDir)
	}

	fs := http.FileServer(http.Dir(hlsDir))
	// wrap in CORS
	corsFs := enableCors(http.StripPrefix("/hls/", fs))
	http.Handle("/hls/", corsFs)

	port := 8080
	fmt.Printf("Starting HLS server on http://localhost:%d/hls/\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// command to run //
// ffmpeg -i ./videos/test.mp4 -bsf:v h264_mp4toannexb -codec copy -hls_list_size 0 videos/test.m3u8
