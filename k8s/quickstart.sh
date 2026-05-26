#!/bin/bash

set -e

echo "🚀 Kubernetes Go Demo - Quick Start Script"
echo "=========================================="
echo ""

# Check if minikube is installed
if ! command -v minikube &> /dev/null; then
    echo "❌ Minikube is not installed. Please install it from: https://minikube.sigs.k8s.io/docs/start/"
    exit 1
fi

# Check if kubectl is installed
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl is not installed. Please install it from: https://kubernetes.io/docs/tasks/tools/"
    exit 1
fi

# Check if docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install it from: https://www.docker.com/products/docker-desktop"
    exit 1
fi

# Check if go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install it from: https://golang.org/dl/"
    exit 1
fi

echo "✅ All prerequisites are installed"
echo ""

# Start minikube if not running
echo "🔄 Checking Minikube status..."
if ! minikube status &> /dev/null; then
    echo "Starting Minikube..."
    minikube start --driver=docker
else
    echo "✅ Minikube is already running"
fi
echo ""

# Configure Docker to use Minikube's Docker daemon
echo "🔄 Configuring Docker to use Minikube's daemon..."
eval $(minikube docker-env)
echo "✅ Docker configured"
echo ""

# Download Go dependencies
echo "🔄 Downloading Go dependencies..."
go mod download
echo "✅ Dependencies downloaded"
echo ""

# Build Docker image
echo "🔄 Building Docker image..."
docker build -t k8s-demo:latest .
echo "✅ Docker image built"
echo ""

# Deploy to Kubernetes
echo "🔄 Deploying to Kubernetes..."
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/serviceaccount.yaml
kubectl apply -f k8s/clusterrole.yaml
kubectl apply -f k8s/clusterrolebinding.yaml
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
echo "✅ Deployed to Kubernetes"
echo ""

# Wait for deployment to be ready
echo "⏳ Waiting for deployment to be ready..."
kubectl wait --for=condition=available --timeout=60s deployment/k8s-demo -n k8s-demo
echo "✅ Deployment is ready!"
echo ""

# Show status
echo "📊 Deployment Status:"
echo "===================="
kubectl get all -n k8s-demo
echo ""

# Get the URL
MINIKUBE_IP=$(minikube ip)
echo "✅ Application deployed successfully!"
echo ""
echo "🌐 Access the application at:"
echo "   http://${MINIKUBE_IP}:30080"
echo ""
echo "📝 Useful commands:"
echo "   View logs:           make logs"
echo "   Check status:        make status"
echo "   Port forward:        make port-forward"
echo "   Delete deployment:   make delete"
echo ""
echo "   Or open in browser:  minikube service k8s-demo -n k8s-demo"
echo ""

