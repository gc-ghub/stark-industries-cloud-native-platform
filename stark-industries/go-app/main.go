package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// Force v2 unless APP_VERSION is explicitly provided
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v2"
	}

	hostname, _ := os.Hostname()
	timestamp := time.Now().Format("Jan 2, 2006 15:04:05")

	// HTML for Canary v2
	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8" />
	<title>Stark Industries App — v2 Canary</title>
	<style>
		body {
			margin: 0;
			font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			background: linear-gradient(135deg, #1a2a3a, #25354a, #345b74);
			color: #fff;
			text-align: center;
			padding: 40px;
		}

		.card {
			background: rgba(255, 255, 255, 0.10);
			padding: 40px;
			margin: auto;
			max-width: 700px;
			border-radius: 18px;
			box-shadow: 0 8px 25px rgba(0,0,0,0.3);
			backdrop-filter: blur(10px);
			animation: fadeIn 1s ease-in-out;
		}

		@keyframes fadeIn {
			from { opacity: 0; transform: translateY(20px); }
			to { opacity: 1; transform: translateY(0); }
		}

		h1 {
			font-size: 44px;
			margin-bottom: 10px;
		}

		.version-badge {
			display: inline-block;
			background: #3399ff;
			padding: 6px 14px;
			border-radius: 12px;
			font-size: 16px;
			margin-bottom: 15px;
			font-weight: bold;
			animation: pulse 1.8s infinite;
		}

		@keyframes pulse {
			0% { transform: scale(1); }
			50% { transform: scale(1.08); }
			100% { transform: scale(1); }
		}

		.footer {
			margin-top: 30px;
			font-size: 14px;
			opacity: 0.8;
		}

	</style>
	</head>

	<body>
		<div class="card">
			<h1>🟦 Stark Industries — Canary Release</h1>

			<div class="version-badge">Version: %s</div>

			<p>This is <strong>v2 Canary</strong>, served using Istio traffic splitting.</p>

			<br />

			<h3>📦 Pod Information</h3>
			<p><strong>Pod Name:</strong> %s</p>

			<h3>⏱ Deployed At</h3>
			<p>%s</p>

			<div class="footer">
				Canary Deployment via Istio • ArgoCD • EKS • GitHub Actions
			</div>
		</div>
	</body>
	</html>
	`, version, hostname, timestamp)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
