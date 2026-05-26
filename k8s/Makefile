.PHONY: help build docker-build deploy deploy-all delete delete-all logs status clean setup-minikube start apply

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Prerequisites:"
	@echo "  - minikube: brew install minikube (macOS) or see https://minikube.sigs.k8s.io/docs/start/"
	@echo "  - kubectl:  brew install kubectl (macOS) or see https://kubernetes.io/docs/tasks/tools/"
	@echo "  - Docker:   Docker Desktop or docker CLI"

start: setup-minikube ## Start minikube and setup environment (alias for setup-minikube)

setup-minikube: ## Setup minikube cluster (if not already running)
	@echo "Checking minikube status..."
	@minikube status || minikube start --driver=docker --wait=all --wait-timeout=5m
	@eval $$(minikube docker-env)
	@echo "Minikube is ready!"

build: ## Build the Go application locally
	@echo "Building Go application..."
	@go mod download
	@go build -o k8s-demo .
	@echo "Build complete!"

docker-build: ## Build Docker image (uses minikube's Docker daemon)
	@echo "Building Docker image..."
	@eval $$(minikube docker-env) && docker build -t k8s-demo:latest .
	@echo "Docker image built successfully!"

apply: ## Apply all Kubernetes manifests (without rebuilding)
	@echo "Applying Kubernetes manifests..."
	@kubectl apply -f k8s/namespace.yaml --validate=false
	@kubectl apply -f k8s/serviceaccount.yaml
	@kubectl apply -f k8s/clusterrole.yaml
	@kubectl apply -f k8s/clusterrolebinding.yaml
	@kubectl apply -f k8s/db-credentials.yaml
	@kubectl apply -f k8s/configmap.yaml
	@kubectl apply -f k8s/postgres-pvc.yaml
	@kubectl apply -f k8s/postgres-deployment.yaml
	@kubectl apply -f k8s/redis-deployment.yaml
	@kubectl apply -f k8s/deployment.yaml
	@kubectl apply -f k8s/service.yaml
	@echo "Resources applied!"

deploy: docker-build apply ## Build Docker image and deploy application to Kubernetes
	@echo "Waiting for deployments to be ready..."
	@kubectl wait --for=condition=available --timeout=120s deployment/k8s-demo -n k8s-demo || true
	@kubectl wait --for=condition=available --timeout=60s deployment/redis -n k8s-demo || true
	@echo "Deployment complete!"
	@echo ""
	@echo "Access the application at: http://$$(minikube ip):30080"
	@echo "Or run: minikube service k8s-demo -n k8s-demo"

deploy-all: start deploy ## Start minikube, build and deploy everything

status: ## Show deployment status
	@echo "=== Namespace ==="
	@kubectl get namespace k8s-demo 2>/dev/null || echo "Namespace not found"
	@echo ""
	@echo "=== Deployments & StatefulSets ==="
	@kubectl get deployments,statefulsets -n k8s-demo 2>/dev/null || echo "No deployments found"
	@echo ""
	@echo "=== Pods ==="
	@kubectl get pods -n k8s-demo -o wide 2>/dev/null || echo "No pods found"
	@echo ""
	@echo "=== Services ==="
	@kubectl get services -n k8s-demo 2>/dev/null || echo "No services found"
	@echo ""
	@echo "=== ConfigMaps & Secrets ==="
	@kubectl get configmaps,secrets -n k8s-demo 2>/dev/null || echo "No configmaps found"
	@echo ""
	@echo "=== PersistentVolumeClaims ==="
	@kubectl get pvc -n k8s-demo 2>/dev/null || echo "No PVCs found"

logs: ## View application logs
	@echo "Fetching logs from k8s-demo pods..."
	@kubectl logs -n k8s-demo -l app=k8s-demo --tail=100 -f

delete: ## Delete all Kubernetes resources
	@echo "Deleting Kubernetes resources..."
	@kubectl delete -f k8s/service.yaml --ignore-not-found=true
	@kubectl delete -f k8s/deployment.yaml --ignore-not-found=true
	@kubectl delete -f k8s/redis-deployment.yaml --ignore-not-found=true
	@kubectl delete -f k8s/postgres-deployment.yaml --ignore-not-found=true
	@kubectl delete -f k8s/postgres-pvc.yaml --ignore-not-found=true
	@kubectl delete -f k8s/configmap.yaml --ignore-not-found=true
	@kubectl delete -f k8s/db-credentials.yaml --ignore-not-found=true
	@kubectl delete -f k8s/clusterrolebinding.yaml --ignore-not-found=true
	@kubectl delete -f k8s/clusterrole.yaml --ignore-not-found=true
	@kubectl delete -f k8s/serviceaccount.yaml --ignore-not-found=true
	@kubectl delete -f k8s/namespace.yaml --ignore-not-found=true
	@echo "All resources deleted!"

delete-all: delete ## Delete all resources and stop minikube
	@echo "Stopping minikube..."
	@minikube stop
	@echo "Minikube stopped!"

clean: delete ## Clean up everything including local binaries
	@echo "Cleaning up local files..."
	@rm -f k8s-demo
	@echo "Clean complete!"

run-local: build ## Run the application locally (requires kubeconfig)
	@echo "Running application locally..."
	@./k8s-demo

test: ## Run Go tests
	@echo "Running tests..."
	@go test -v ./...

fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

shell: ## Open a shell in one of the pods
	@kubectl exec -it -n k8s-demo $$(kubectl get pods -n k8s-demo -l app=k8s-demo -o jsonpath='{.items[0].metadata.name}') -- /bin/sh

port-forward: ## Forward port 8080 to localhost
	@echo "Forwarding port 8080..."
	@echo "Access the application at: http://localhost:8080"
	@kubectl port-forward -n k8s-demo service/k8s-demo 8080:8080

port-forward-postgres: ## Forward PostgreSQL port 5432 to localhost
	@echo "Forwarding PostgreSQL port 5432..."
	@echo "Connect to: postgresql://guest:guest@localhost:5432/demoapp"
	@kubectl port-forward -n k8s-demo service/postgres 5432:5432

port-forward-redis: ## Forward Redis port 6379 to localhost
	@echo "Forwarding Redis port 6379..."
	@echo "Connect with: redis-cli -h localhost -p 6379"
	@kubectl port-forward -n k8s-demo service/redis 6379:6379

