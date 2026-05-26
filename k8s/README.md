# Kubernetes Demo in Go

A comprehensive Go application that demonstrates various Kubernetes capabilities using the official `client-go` library. This project showcases how to interact with Kubernetes API programmatically and runs locally using Minikube or any other local Kubernetes cluster.

## 🚀 Features

This demo application showcases:

- **Kubernetes Client Operations**:
  - Listing namespaces, pods, nodes, services, and deployments
  - Creating and updating ConfigMaps
  - Watching pod events in real-time
  - Using both in-cluster and out-of-cluster configuration

- **Production-Ready Patterns**:
  - Health and readiness probes
  - Service Account with RBAC
  - Resource limits and requests
  - Proper logging
  - Multi-stage Docker builds

- **Local Development**:
  - Easy setup with Minikube
  - Makefile for common operations
  - Port forwarding for local access

## 📋 Prerequisites

Before running this project, ensure you have the following installed:

- **Go** 1.21 or higher ([download](https://golang.org/dl/))
- **Docker** ([download](https://www.docker.com/products/docker-desktop))
- **kubectl** ([installation guide](https://kubernetes.io/docs/tasks/tools/))
- **Minikube** ([installation guide](https://minikube.sigs.k8s.io/docs/start/))

## 🛠️ Quick Start

### 1. Setup Minikube

Start a local Kubernetes cluster:

```bash
make setup-minikube
```

This will:
- Start Minikube if it's not already running
- Configure Docker to use Minikube's Docker daemon

### 2. Build and Deploy

Deploy the application to your local Kubernetes cluster:

```bash
make deploy
```

This command will:
- Build the Docker image
- Create the namespace and RBAC resources
- Deploy the application
- Create a NodePort service

### 3. Access the Application

After deployment, access the application using one of these methods:

**Option A: Using Minikube service (recommended)**
```bash
minikube service k8s-demo -n k8s-demo
```

**Option B: Using NodePort**
```bash
# Get the Minikube IP and access on port 30080
curl http://$(minikube ip):30080
```

**Option C: Using port forwarding**
```bash
make port-forward
# Then access at http://localhost:8080
```

### 4. View Application Logs

Watch the application logs in real-time:

```bash
make logs
```

You'll see the application performing various Kubernetes operations:
- Listing all namespaces in the cluster
- Listing pods in the k8s-demo namespace
- Listing cluster nodes
- Creating/updating a ConfigMap
- Listing services
- Getting deployment information
- Watching for pod events

## 📁 Project Structure

```
k8s/
├── main.go                 # Main application code
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies
├── Dockerfile              # Multi-stage Docker build
├── Makefile                # Build and deployment commands
├── README.md               # This file
└── k8s/                    # Kubernetes manifests
    ├── namespace.yaml      # k8s-demo namespace
    ├── serviceaccount.yaml # Service account for the app
    ├── clusterrole.yaml    # RBAC role definition
    ├── clusterrolebinding.yaml # RBAC role binding
    ├── configmap.yaml      # Application configuration
    ├── deployment.yaml     # Application deployment
    └── service.yaml        # NodePort service
```

## 🎯 What the Application Does

The Go application demonstrates various Kubernetes operations:

### 1. **Namespace Operations**
Lists all namespaces in the cluster with their status.

### 2. **Pod Management**
Lists all pods in the current namespace, showing:
- Pod name and status
- Pod IP address
- Container details (name and image)

### 3. **Node Information**
Lists cluster nodes with:
- Node names
- OS information
- Kubernetes version
- Internal IP addresses

### 4. **ConfigMap Operations**
Creates or updates a ConfigMap with:
- Application metadata
- Timestamp of creation/update
- Custom configuration data

### 5. **Service Discovery**
Lists services with:
- Service names and types
- ClusterIP addresses
- Port configurations

### 6. **Deployment Information**
Shows deployment details:
- Replica counts (desired, ready, available)
- Container images

### 7. **Event Watching**
Continuously watches for pod events and logs them in real-time.

## 🎮 Available Make Commands

| Command | Description |
|---------|-------------|
| `make help` | Display all available commands |
| `make setup-minikube` | Setup and start Minikube |
| `make build` | Build the Go application locally |
| `make docker-build` | Build Docker image |
| `make deploy` | Deploy to Kubernetes |
| `make status` | Show deployment status |
| `make logs` | View application logs |
| `make delete` | Delete all Kubernetes resources |
| `make clean` | Clean up everything |
| `make run-local` | Run locally (requires kubeconfig) |
| `make port-forward` | Forward port 8080 to localhost |
| `make shell` | Open shell in a pod |
| `make fmt` | Format Go code |
| `make vet` | Run go vet |

## 🔍 Exploring the Application

### Check Deployment Status

```bash
make status
```

### View Real-time Logs

```bash
make logs
```

### Access the HTTP Endpoints

The application exposes three HTTP endpoints:

```bash
# Health check
curl http://localhost:8080/health

# Readiness check
curl http://localhost:8080/ready

# Main endpoint
curl http://localhost:8080/
```

### Execute Commands in a Pod

```bash
make shell
```

## 🔧 Running Locally (Development)

You can also run the application locally (outside Kubernetes):

```bash
make run-local
```

This requires:
- A valid kubeconfig file at `~/.kube/config`
- Access to a Kubernetes cluster (Minikube, kind, or remote cluster)

The application will automatically use your kubeconfig to connect to the cluster.

## 🔑 RBAC Permissions

The application uses a ServiceAccount with a ClusterRole that grants:

- **Read permissions** on:
  - Namespaces
  - Pods
  - Services
  - Nodes
  - Deployments
  
- **Write permissions** on:
  - ConfigMaps (create, update, patch)

This demonstrates proper security practices by granting only the minimum required permissions.

## 📚 Learning Resources

This project demonstrates concepts from:

- [Kubernetes Official Documentation](https://kubernetes.io/docs/)
- [client-go Library](https://github.com/kubernetes/client-go)
- [Kubernetes API Concepts](https://kubernetes.io/docs/reference/using-api/api-concepts/)

## 🧹 Cleanup

To remove all resources:

```bash
make clean
```

To stop Minikube:

```bash
minikube stop
```

To completely delete the Minikube cluster:

```bash
minikube delete
```

## 🐛 Troubleshooting

### Pods in CrashLoopBackOff or Restart Loop?

If your pods are restarting continuously, use the debug script to diagnose the issue:

```bash
./debug-pod.sh
```

This script will show:
- Pod status and events
- Current and previous container logs
- RBAC configuration
- Docker image status
- Deployment details

**Common causes of restart loops:**
1. **HTTP Server Binding Issues**: The app has been updated to ensure the HTTP server binds successfully before health checks start
2. **RBAC Permission Errors**: Verify the ServiceAccount and ClusterRole are correctly configured
3. **Image Not Found**: Rebuild the image with `make docker-build`
4. **Resource Limits**: Check if pods have enough CPU/Memory

**Manual debugging steps:**
```bash
# Check pod logs (current)
kubectl logs -n k8s-demo -l app=k8s-demo

# Check pod logs (previous crash)
kubectl logs -n k8s-demo -l app=k8s-demo --previous

# Check pod events
kubectl get events -n k8s-demo --sort-by='.lastTimestamp'

# Describe the pod for detailed status
kubectl describe pod -n k8s-demo -l app=k8s-demo
```

**Quick fixes:**
```bash
# Rebuild and redeploy
make delete
make deploy

# Just rebuild the image
make docker-build
kubectl rollout restart deployment/k8s-demo -n k8s-demo
```

### Pods not starting?

Check pod status and events:
```bash
kubectl describe pod -n k8s-demo -l app=k8s-demo
```

### Can't access the service?

Ensure Minikube is running and check service status:
```bash
minikube status
make status
```

### Permission errors?

Verify RBAC resources are created:
```bash
kubectl get clusterrole k8s-demo-reader
kubectl get clusterrolebinding k8s-demo-reader-binding
```

### Go module checksum errors?

If you encounter checksum mismatch errors during build:
```bash
go clean -modcache
go mod download
go mod verify
```

## 🚀 Next Steps

To extend this project, consider:

1. **Add More Operations**: Implement create/update/delete for other resources
2. **Add Metrics**: Integrate Prometheus metrics
3. **Add Tests**: Write unit and integration tests
4. **Add CI/CD**: Create GitHub Actions workflows
5. **Add Helm Charts**: Package the application with Helm
6. **Add Custom Resources**: Work with CRDs (Custom Resource Definitions)
7. **Add Operators**: Build a Kubernetes operator

## 📄 License

This is a demo project for educational purposes.

## 🤝 Contributing

Feel free to fork, modify, and use this project for learning Kubernetes and Go!

