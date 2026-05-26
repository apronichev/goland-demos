# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a collection of 35+ Go demonstration projects showcasing GoLand/JetBrains IDE features, Go language capabilities (versions 1.20-1.26), and integration patterns. Each subdirectory is an independent project with its own `go.mod` file.

## Project Categories

### Go Language Features
- **go125/** - Go 1.25 features: JSON v2, testing improvements, GC enhancements, repanic, vet analyzers
- **go126/** - Go 1.26 features: Buffer peek operations, multi-handler routing, new variable semantics

### Web Applications
- **todo-app/** - REST API with SQLite backend, demonstrates generics pattern
- **web-notebook/** - Note-taking web app with database persistence
- **grpc-demo/** - Dual gRPC+HTTP server with protocol buffers

### Infrastructure & Cloud
- **k8s/** - Kubernetes client operations with minikube setup
- **terraform/** - Terraform SDK integration with state/plan parsing
- **dev-containers/** - Docker Compose development environment

### IDE Feature Demonstrations
- **debugging-features/** - Debugger capabilities (breakpoints, goroutines, interfaces)
- **inspections-lighting-talk/** - Code inspection patterns
- **data-flow-demo/** - Data flow analysis demonstrations
- **refactoring-demo/** - Refactoring capabilities showcase
- **structure-view/** - Structure view and navigation features

### Testing & Quality
- **fuzz-testing/** - Fuzzing demonstrations
- **profiling-applications/** - Performance profiling and benchmarking
- **vulnerability-checker/** - Security and dependency analysis
- **linters/** - Linter demonstrations

## Common Development Commands

### Building and Running Individual Projects

Each project is independent. Navigate to the project directory first:

```bash
cd <project-name>
go mod download
go run main.go
# or for nested server structure:
go run server/main.go
```

### Web Application Projects

**todo-app:**
```bash
cd todo-app
go run main.go
# Access at http://localhost:8080
# Creates todos.db SQLite database
```

**web-notebook:**
```bash
cd web-notebook
go get modernc.org/sqlite
go run server/main.go
# Access at http://127.0.0.1:8080
```

**grpc-demo:**
```bash
cd grpc-demo
go run main.go
# gRPC server: localhost:50051
# HTTP server: localhost:8080
# OpenAPI spec: http://localhost:8080/grpc.yaml
```

### Kubernetes Project

**k8s/** uses a comprehensive Makefile:
```bash
cd k8s
make help                # Show all available commands
make deploy-all          # Start minikube, build, and deploy
make build               # Build Go application locally
make docker-build        # Build Docker image (uses minikube's Docker)
make apply               # Apply Kubernetes manifests
make status              # Show deployment status
make logs                # View application logs
make port-forward        # Forward port 8080 to localhost
make port-forward-postgres  # Forward PostgreSQL port
make port-forward-redis     # Forward Redis port
make delete              # Delete all Kubernetes resources
make clean               # Clean up everything including binaries
```

### Testing

```bash
# Run tests in a specific project
cd <project-name>
go test -v ./...

# Run fuzzing tests
cd fuzz-testing
go test -fuzz=FuzzReverse

# Run benchmarks with allocation tracking
cd profiling-applications
go test -bench=. -benchmem

# Run allocation output tests (Go 1.25+)
cd go125/test-allocs-output
go test -test.allocationsoutput
```

### Code Quality

```bash
# Format code
go fmt ./...

# Run go vet
go vet ./...

# Tidy dependencies
go mod tidy
```

## Architecture Patterns

### Generic Storage Pattern (todo-app)

The todo-app demonstrates a clean architecture with generics:

```
todo-app/
├── main.go              # HTTP handlers and routing
├── model/
│   └── model.go         # Domain types (TodoItem)
└── storage/
    ├── storage.go       # Generic Storage[T Item] interface
    ├── sqlite.go        # SQLite implementation
    └── inMemory.go      # In-memory implementation
```

The storage layer uses Go generics for type-safe operations:
```go
type Storage[T Item] interface {
    Get(id int) T
    Put(item T) T
    GetAll() []T
    Remove(id int)
}
```

### HTTP Handler Pattern

Most web apps use standard library routing with method-based switching:
```go
http.HandleFunc("/todos", handler)
http.HandleFunc("/todos/{id}", handler)  // Path parameters with PathValue()

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // handle GET
    case http.MethodPost:
        // handle POST
    }
}
```

### Database Integration

Projects using SQLite:
- Use `modernc.org/sqlite` (pure Go implementation)
- Database files created in project root (e.g., `todos.db`, `notebook.db`)
- Can be inspected using IDE's Database tool window

## Key Technologies

**Databases:**
- SQLite via `modernc.org/sqlite` v1.30-v1.39
- PostgreSQL (in k8s demo)
- Redis (in k8s demo)

**RPC & APIs:**
- gRPC with `google.golang.org/grpc` v1.70+
- Protocol Buffers v3
- Standard HTTP REST

**Infrastructure:**
- Kubernetes via `k8s.io/client-go`
- Terraform SDK (plugin-sdk v2.30, terraform-exec v0.19)
- Docker and Docker Compose

## Development Notes

### IDE Feature Demonstrations

Many projects contain intentional `//todo` comments that showcase IDE quick-fixes and inspections:
- "invoke 'implement missing methods' quick fix"
- "invoke 'inline variable' refactoring"
- "type case and see FLCC completion" (Full Line Code Completion)
- "invoke 'handle error' quick fix"

These are educational markers, not actual TODOs to complete.

### Project Independence

Each subdirectory is a standalone project with its own:
- `go.mod` file for dependency management
- Independent versioning (Go 1.20-1.26 across projects)
- No shared code between projects

### Go Version Requirements

Projects demonstrate cutting-edge features:
- **Go 1.26**: terraform/ (latest features)
- **Go 1.25**: linters/, profiling-applications/, go125/* (JSON v2, testing enhancements)
- **Go 1.22+**: Most web applications
- Always check the project's `go.mod` for specific version requirements

### Protocol Buffers (grpc-demo)

Generated code is committed to `proto/` directory:
- `user.pb.go` - Protocol buffer types
- `user_grpc.pb.go` - gRPC service definitions

To regenerate:
```bash
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

### Container Development

The **dev-containers/** project includes:
- `.devcontainer/devcontainer.json` for VS Code/JetBrains
- `docker-compose.yml` with database initialization
- Post-create commands for dependency installation

## Common Patterns to Follow

1. **Error Handling**: Always check and handle errors explicitly
2. **JSON Encoding**: Use `json.NewEncoder(w).Encode()` for HTTP responses
3. **HTTP Status Codes**: Use `http.Status*` constants, never magic numbers
4. **Path Parameters**: Use `r.PathValue("id")` for path variables (Go 1.22+)
5. **Generics**: Use type parameters `[T constraint]` for reusable components
6. **Interfaces**: Define interfaces for abstraction (e.g., Storage layer)

## Troubleshooting

### Port Already in Use
Most web apps use port 8080. If unavailable, modify the `ListenAndServe` call in `main.go`.

### SQLite Database Locked
Stop any running instances of the application before restarting.

### Kubernetes Issues
```bash
cd k8s
make status    # Check current state
make delete    # Clean up and retry
make deploy-all
```

### Missing Dependencies
```bash
go mod download
go mod tidy
```
