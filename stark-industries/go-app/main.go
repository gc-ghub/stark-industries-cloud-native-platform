package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Canary version
	version := "v2-canary"

	hostname, _ := os.Hostname()
	timestamp := time.Now().Format("Jan 2, 2006 15:04:05")

	// *** IMPORTANT *** → Exactly 3 %s placeholders
	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8" />
	<title>Stark Industries — Canary</title>
	<style>
		body {
			margin: 0;
			font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			background: linear-gradient(135deg, #0f2027, #203a43, #2c5364);
			color: #fff;
			text-align: center;
			padding: 40px;
		}
		.card {
			background: rgba(255,255,255,0.08);
			padding: 40px;
			margin: auto;
			max-width: 700px;
			border-radius: 18px;
			box-shadow: 0 8px 25px rgba(0,0,0,0.3);
			backdrop-filter: blur(10px);
		}
		h1 {
			font-size: 42px;
			margin-bottom: 15px;
		}
		.version-badge {
			display: inline-block;
			background: #3399ff;
			padding: 8px 18px;
			border-radius: 12px;
			font-size: 18px;
			font-weight: bold;
			margin-bottom: 22px;
		}
		.footer {
			margin-top: 30px;
			font-size: 14px;
			opacity: 0.85;
		}
	</style>
	</head>

	<body>
		<div class="card">
			<h1>🔷 Stark Industries — Canary Release</h1>

			<div class="version-badge">Version: %s</div>

			<p>This is <strong>v2 Canary</strong>, deployed via Istio traffic splitting.</p>

			<h3>📦 Pod Information</h3>
			<p><strong>Pod Name:</strong> %s</p>

			<h3>⏱ Deployed At</h3>
			<p>%s</p>

			<div class="footer">
				Canary Deployment • Istio • ArgoCD • GitHub Actions • EKS
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
