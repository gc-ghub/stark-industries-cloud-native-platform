package main

import (
	"fmt"
	"net/http"
	"os"
)

// Main application handler for /
func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1"
	}
	fmt.Fprintf(w, "<h1>Stark Industries App (%s)</h1><p>Serving from Go app</p>", version)
}

// NEW — simple health endpoint for probes
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	// Main route
	http.HandleFunc("/", handler)

	// Health route — used by K8s probes
	http.HandleFunc("/healthz", healthHandler)

	// Start HTTP server
	http.ListenAndServe(":8080", nil)
}
