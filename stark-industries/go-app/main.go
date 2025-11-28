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
	<title>Stark Industries App</title>
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
			background: rgba(255, 255, 255, 0.08);
			padding: 40px;
			margin: auto;
			max-width: 700px;
			border-radius: 18px;
			box-shadow: 0 8px 25px rgba(0,0,0,0.3);
			backdrop-filter: blur(10px);
			transition: transform 0.3s ease;
		}

		.card:hover {
			transform: scale(1.02);
		}

		h1 {
			font-size: 44px;
			margin-bottom: 10px;
		}

		.version-badge {
			display: inline-block;
			background: #ff6b6b;
			padding: 6px 14px;
			border-radius: 12px;
			font-size: 16px;
			margin-bottom: 15px;
			font-weight: bold;
		}

		.footer {
			margin-top: 30px;
			font-size: 14px;
			opacity: 0.8;
		}

		a {
			color: #00eaff;
			text-decoration: none;
		}

		a:hover {
			text-decoration: underline;
		}
	</style>
	</head>

	<body>
		<div class="card">
			<h1>🚀 Stark Industries Dashboard</h1>

			<div class="version-badge">Version: %s</div>

			<p>This app is auto-deployed using <strong>GitHub Actions → ECR → ArgoCD → EKS</strong></p>

			<br />

			<h3>📦 Pod Information</h3>
			<p><strong>Pod Name:</strong> %s</p>

			<h3>⏱ Deployed At</h3>
			<p>%s</p>

			<div class="footer">
				We build and ship War Machines 🤖 and Tony Stark approves this deployment! ⚡
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
