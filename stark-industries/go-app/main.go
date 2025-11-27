package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1"
	}

	fmt.Fprintf(w, `
        <h1 style="color:#c0392b; font-family:Arial; text-shadow:1px 1px #000;">
            🚀 Stark Industries App (%s)
        </h1>
        <p style="font-size:20px;">This is a CI/CD + GitOps auto-deployed version.</p>
        <p>Deployed via GitHub Actions → ECR → ArgoCD → EKS</p>
    `, version)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthHandler)
	http.ListenAndServe(":8080", nil)
}
