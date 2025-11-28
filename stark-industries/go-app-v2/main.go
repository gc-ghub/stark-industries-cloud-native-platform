package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1"
	}

	hostname, _ := os.Hostname()
	timestamp := time.Now().Format("Jan 2, 2006 15:04:05")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8" />
	<title>⚙️ Doom Industries Control Panel</title>
	<style>
		body {
			margin: 0;
			font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			background: linear-gradient(135deg, #000000, #1b1b1b, #2e2e2e);
			color: #e0e0e0;
			text-align: center;
			padding: 40px;
		}

		.card {
			background: rgba(15, 255, 15, 0.08);
			padding: 40px;
			margin: auto;
			max-width: 700px;
			border-radius: 18px;
			box-shadow: 0 8px 25px rgba(0,255,0,0.2);
			backdrop-filter: blur(10px);
			transition: transform 0.3s ease;
			border: 1px solid #00ff00;
		}

		.card:hover {
			transform: scale(1.02);
		}

		h1 {
			font-size: 44px;
			margin-bottom: 10px;
			color: #00ff00;
			text-shadow: 0 0 10px #00ff00;
		}

		.version-badge {
			display: inline-block;
			background: #005500;
			padding: 6px 14px;
			border-radius: 12px;
			font-size: 16px;
			margin-bottom: 15px;
			font-weight: bold;
			color: #00ff00;
			border: 1px solid #00ff00;
		}

		.footer {
			margin-top: 30px;
			font-size: 14px;
			opacity: 0.8;
			color: #9aff9a;
		}

		a {
			color: #00ff00;
			text-decoration: none;
		}

		a:hover {
			text-decoration: underline;
		}
	</style>
	</head>

	<body>
		<div class="card">
			<h1>💀 Doom Industries Command Console</h1>

			<div class="version-badge">Version: %s</div>

			<p>This system is deployed through <strong>GitHub Actions → ECR → ArgoCD → EKS</strong>,  
			under the supreme authority of <b>Victor Von Doom</b>.</p>

			<br />

			<h3>🛰 Pod Telemetry</h3>
			<p><strong>Unit Identifier:</strong> %s</p>

			<h3>⏱ Activation Timestamp</h3>
			<p>%s</p>

			<div class="footer">
				“Doom needs no permission. Doom only needs power.” ⚡
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
