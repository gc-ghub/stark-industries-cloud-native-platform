# Stark Industries — Cloud Native Platform
Production-grade EKS + GitOps + Canary Deployments with Istio

[![Demo](https://img.shields.io/badge/demo-ready-blue)](#)
[![GitOps](https://img.shields.io/badge/GitOps-ArgoCD-orange)](#)
[![CI](https://img.shields.io/badge/CI-GitHub%20Actions-brightgreen)](#)

A full DevOps + Cloud-Native reference project demonstrating **CI/CD, GitOps, Kubernetes, Service Mesh, and Canary Deployments** using real AWS infrastructure.



## 📌 Tools Used

| Area | Tool(s) | Purpose / Notes |
|---|---|---|
| ☁️ Cloud | AWS (EKS, ECR, ELB) | Managed Kubernetes (EKS), container registry (ECR), and cloud load balancers (ELB). Chart values use region ap-southeast-1 as an example. |
| ⚙️ CI | GitHub Actions | Builds containers, runs tests, and updates manifests/images in the git repo. |
| 🔁 CD / GitOps | Argo CD | Watches Git repository and automatically syncs chart/manifests to clusters. |
| 📦 Containers | Docker | Multi-stage builds (golang builder + distroless runtime) are used for minimal, secure images. |
| 🧭 Orchestration | Kubernetes (EKS) | Deploys workloads, services, ingress and manages lifecycle. |
| 📦 Packaging | Helm (v3+) | Chart available at `stark-industries/helm/stark-industries-go` for packaging and releases. |
| 🌐 Service mesh | Istio | Traffic management and advanced routing (used for canary rollouts). |
| 🔍 Observability | Kiali | Visualize service mesh topology, traffic flows and metrics (screenshots provided). |
| 🧩 Languages | Go (golang:1.22 in Dockerfile) | App code is written in Go; Dockerfile builds using golang:1.22-alpine. |
| 🎯 Canary rollouts | Istio routes (VirtualService + DestinationRule) | Weighted routing for progressive rollouts and verification of canary behavior. |
| 🧪 Local testing | kubectl, helm, minikube/kind | Useful for verifying manifests and testing locally before pushing to cloud. |


---

## 🧭 Table of contents

1. 🔎 [Quick overview](#quick-overview)
2. 📦 [What's included](#whats-included)
3. 🧑‍💻 [Quick start (local development)](#quick-start-local-development)
4. 🛠️ [Building & publishing containers](#building-and-publishing-containers)
5. 📜 [Helm chart usage](#helm-chart-usage)
6. ☸️ [Kubernetes manifests (direct deploy)](#kubernetes-manifests-direct-deploy)
7. 🎯 [Canary routing with Istio](#canary-routing-with-istio)
8. 🔁 [CI / GitOps flow (overview)](#ci--gitops-flow-overview)
9. 🗂️ [Project structure](#project-structure)
10. 🧰 [Troubleshooting & common editor lint issues](#troubleshooting--editor-lint-issues)
11. 🧹 [Cleanup](#cleanup)
12. 🤝 [Contributing](#contributing)

---

## Quick overview

This repository demonstrates:
- Two small Go web apps (v1 = Stark, v2 = Doom) that render a simple web page showing version & hostname.
- Docker multi-stage builds (secure distroless runtime image)
- Helm chart for packaging and deploying the app(s)
- Kubernetes manifests for quick testing (deployments, service, ingress, Istio virtual-service/destination rule)
- Example GitOps flow using GitHub Actions + ArgoCD

The sample setup showcases a canary deployment where traffic can be split between v1 and v2 using Istio.

---
## What's included


- `stark-industries/go-app/` — v1 Stark web app
- `stark-industries/go-app-v2/` — v2 Doom web app
- `stark-industries/helm/stark-industries-go/` — Helm chart (templates, values)
- `stark-industries/k8s/manifests/` — Kubernetes manifest examples
- `stark-industries/k8s/argocd/` — example ArgoCD Application manifest

---

## Quick start (local development)

Prerequisites
- Go 1.20+
- Docker (to build/run images)
- Optional: Helm, kubectl, a local K8s cluster (minikube or kind)

Run the Go services locally (no container):

```powershell
cd stark-industries/go-app
go run main.go

# in another shell
cd stark-industries/go-app-v2
go run main.go
```

Both apps listen on :8080. Open http://localhost:8080 and refresh — the page will display the app version and host information.

---

## Building and publishing containers

Build locally:

```bash
cd stark-industries/go-app
docker build -t stark-industries-go:local .

# run
docker run --rm -p 8080:8080 -e APP_VERSION=v1 stark-industries-go:local
```

Push to AWS ECR (example):

```bash
# create repo (if missing)
aws ecr create-repository --repository-name stark-industries-go-web-app --region ap-southeast-1

# login
aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin <ACCOUNT>.dkr.ecr.ap-southeast-1.amazonaws.com

# tag & push
docker tag stark-industries-go:local <ACCOUNT>.dkr.ecr.ap-southeast-1.amazonaws.com/stark-industries-go-web-app:1.0.0
docker push <ACCOUNT>.dkr.ecr.ap-southeast-1.amazonaws.com/stark-industries-go-web-app:1.0.0
```

Replace `<ACCOUNT>` with your AWS account id.

---

## Helm chart usage

Chart path: `stark-industries/helm/stark-industries-go`

Install into cluster:

```bash
helm install my-stark ./stark-industries/helm/stark-industries-go \
	--set image.repository=<ACCOUNT>.dkr.ecr.ap-southeast-1.amazonaws.com/stark-industries-go-web-app \
	--set image.tag=1.0.0

# upgrade
helm upgrade my-stark ./stark-industries/helm/stark-industries-go --reuse-values

# package chart
helm package ./stark-industries/helm/stark-industries-go
```

Notes
- Replace the image and tag with the images you pushed to your container registry.
- Avoid using `latest` in production — prefer immutable tags or image digests.

---

## Kubernetes manifests (direct deploy)

Quick apply the example manifests for v1/v2 and service:

```bash
kubectl apply -f stark-industries/k8s/manifests/deployment-v1.yaml
kubectl apply -f stark-industries/k8s/manifests/deployment-v2.yaml
kubectl apply -f stark-industries/k8s/manifests/service.yaml
```

Istio-related examples are also included (virtual-service.yaml, destination-rule.yaml) to run canary routing experiments.

---

## Canary routing with Istio

This repository demonstrates how to split traffic between `v1` and `v2` by:

- Defining `subsets` in a DestinationRule (labels: `version: v1` / `version: v2`).
- Setting weights in VirtualService to control traffic routing (0–100).

Example route snippet (50/50):

```yaml
http:
  - route:
    - destination:
        host: my-service
        subset: v1
      weight: 50
    - destination:
        host: my-service
        subset: v2
      weight: 50
```

 
---
## Screenshots & analysis

Below are screenshots included in the `pictures/` directory with short descriptions so you (and others) can quickly understand what each image shows. All images are stored at `pictures/<filename>` and are ready to view in this repo.

### ArgoCD / GitOps

The screenshots below show ArgoCD in action — application list, diffs, resource tree, sync history and logs.

![ArgoCD — Applications view](pictures/argocd-1.png)
*Applications list & sync status.*

![ArgoCD — App diffs & recent syncs](pictures/argocd-2.png)
*App diffs & recent syncs.*

![ArgoCD — Resource tree & health](pictures/argocd-3.png)
*Resource tree & health.*

![ArgoCD — Sync operations](pictures/argocd-4.png)
*Sync operations & rollout history.*

![ArgoCD — Manifests view](pictures/argocd-5.png)
*Manifests comparison view.*

![ArgoCD — Resource summary](pictures/argocd-6.png)
*Resource summary snapshot.*

![ArgoCD — Events timeline](pictures/argocd-7.png)
*Events / sync timeline.*

![ArgoCD — Logs / resource detail](pictures/argocd-8.png)
*Logs and resource inspection view.*

### Application UI / Canary examples

![Stark v1 app UI](pictures/go-app.png)
*Stark (v1) live UI snapshot.*

![Stark v1 alternate snapshot](pictures/go-app-v1.png)
*Alternate Stark (v1) snapshot.*

![Canary test — v1 response](pictures/canary-test-go-app-v1.png)
*Example response served by v1 during a canary test.*

![Canary test — v2 response](pictures/canary-test-go-app-v2.png)
*Example response served by v2 (Doom) during a canary test.*

### Istio / Kiali (observability & traffic)

![Istio install / control plane](pictures/istio-install.png)
*Istio install / control-plane snapshot.*

![Kiali — workloads view](pictures/kiali-workloads-1.png)
*Kiali workloads & versioned services.*

![Kiali — traffic graph](pictures/kiali-traffic-graph-1.png)
*Traffic graph — shows live traffic between services.*

![Kiali — canary traffic graph](pictures/kiali-traffic-graph-canary.png)
*Canary split visualization.*

![Kiali — canary traffic graph (alt)](pictures/kiali-traffic-graph-canary-2.png)
*Alternate canary traffic graph.*

![Kiali — metrics snapshot](pictures/kiali-metrics-1.png)
*Metrics panel (latency / error rates / request rate).* 

![Kiali — metrics alt](pictures/kiali-metric-2.png)
*Alternate metrics view.*

![Kiali — mesh topology](pictures/kiali-mesh-1.png)
*Mesh topology overview.*

![Kiali — mesh topology (alt)](pictures/kiali-mesh-2.png)
*Mesh topology with metrics.*

![Kiali — mesh topology (extra)](pictures/kiali-mesh-3.png)
*Additional mesh snapshot (kiali-mesh-3.png).* 

### Misc / Misc screenshots

![Flow diagram / overview](pictures/flow_diagram.png)
*flow_diagram.png — high-level pipeline illustration.*

![Misc local screenshot](pictures/Screenshot 2025-11-29 042602.png)
*Screenshot 2025-11-29 042602.png — misc timestamped capture.*

---

## Cleanup

Example commands to remove demo resources (replace placeholders and confirm before running):

```bash
eksctl delete cluster --name stark-eks --region ap-southeast-1
aws ecr delete-repository --repository-name stark-industries-go-web-app --region ap-southeast-1 --force
aws ecr delete-repository --repository-name stark-industries-go-web-app-v2 --region ap-southeast-1 --force
```

---

## Contributing

Contributions welcome — please keep changes small and provide tests or manual verification steps. Update README/docs if you add or change charts/manifests.

Suggested PR checklist
- Build images successfully
- `helm template` renders without errors
- `kubectl apply --dry-run=client` passes for example manifests

---


<!-- screenshots are embedded in-section (ArgoCD, App UI and Istio) to avoid duplication -->
