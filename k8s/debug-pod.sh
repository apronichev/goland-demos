#!/bin/bash

# Debug script for troubleshooting pod restart loops

NAMESPACE="k8s-demo"
APP_LABEL="app=k8s-demo"

echo "🔍 Kubernetes Pod Debug Script"
echo "================================"
echo ""

# Check if namespace exists
echo "📦 Checking namespace..."
if ! kubectl get namespace $NAMESPACE &> /dev/null; then
    echo "❌ Namespace '$NAMESPACE' does not exist"
    exit 1
fi
echo "✅ Namespace exists"
echo ""

# Get pod status
echo "📊 Pod Status:"
echo "--------------"
kubectl get pods -n $NAMESPACE -l $APP_LABEL
echo ""

# Get pod events
echo "📋 Recent Events:"
echo "-----------------"
kubectl get events -n $NAMESPACE --sort-by='.lastTimestamp' | tail -20
echo ""

# Get pod description
echo "📝 Pod Description:"
echo "-------------------"
POD_NAME=$(kubectl get pods -n $NAMESPACE -l $APP_LABEL -o jsonpath='{.items[0].metadata.name}' 2>/dev/null)
if [ -n "$POD_NAME" ]; then
    kubectl describe pod $POD_NAME -n $NAMESPACE | tail -50
    echo ""
    
    # Get logs from current container
    echo "📄 Current Container Logs:"
    echo "--------------------------"
    kubectl logs $POD_NAME -n $NAMESPACE --tail=50 2>/dev/null || echo "No logs available"
    echo ""
    
    # Get logs from previous container (if it crashed)
    echo "📄 Previous Container Logs (if crashed):"
    echo "-----------------------------------------"
    kubectl logs $POD_NAME -n $NAMESPACE --previous --tail=50 2>/dev/null || echo "No previous logs available (pod may not have restarted yet)"
    echo ""
else
    echo "❌ No pods found with label $APP_LABEL"
fi

# Check deployment status
echo "🚀 Deployment Status:"
echo "--------------------"
kubectl get deployment k8s-demo -n $NAMESPACE -o wide 2>/dev/null || echo "Deployment not found"
echo ""

# Check service account and RBAC
echo "🔐 RBAC Configuration:"
echo "---------------------"
echo "ServiceAccount:"
kubectl get serviceaccount k8s-demo -n $NAMESPACE 2>/dev/null || echo "ServiceAccount not found"
echo ""
echo "ClusterRole:"
kubectl get clusterrole k8s-demo-reader 2>/dev/null || echo "ClusterRole not found"
echo ""
echo "ClusterRoleBinding:"
kubectl get clusterrolebinding k8s-demo-reader-binding 2>/dev/null || echo "ClusterRoleBinding not found"
echo ""

# Check if the image exists
echo "🐳 Docker Image Status:"
echo "----------------------"
eval $(minikube docker-env 2>/dev/null)
docker images | grep k8s-demo || echo "Image not found in minikube's docker daemon"
echo ""

# Recommendations
echo "💡 Troubleshooting Tips:"
echo "------------------------"
echo "1. Check the 'Previous Container Logs' above for crash reasons"
echo "2. Verify the image was built: make docker-build"
echo "3. Check RBAC permissions are correctly configured"
echo "4. Verify the HTTP server can bind to port 8080"
echo "5. Check if the pod has enough resources (CPU/Memory)"
echo ""
echo "Common fixes:"
echo "  - Rebuild image: make docker-build"
echo "  - Recreate deployment: make delete && make deploy"
echo "  - Check pod logs: kubectl logs -f <pod-name> -n $NAMESPACE"
echo ""


