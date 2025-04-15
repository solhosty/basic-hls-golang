package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hlsDir := "./videos"
	if _, err := os.Stat(hlsDir); os.IsNotExist(err) {
		log.Fatalf("HLS directory %s does not exist", hlsDir)
	}
	fs := http.FileServer(http.Dir(hlsDir))
	http.Handle("/p/", http.StripPrefix("/p/", fs))

	port := 8080
	fmt.Printf("Starting HLS server on http://localhost:%d/p/\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// command to run //
// ffmpeg -i ./videos/test.mp4 -bsf:v h264_mp4toannexb -codec copy -hls_list_size 0 output.m3u8
