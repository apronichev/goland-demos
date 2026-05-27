# GoLand Demos

A collection of 35+ Go demonstration projects showcasing [GoLand](https://www.jetbrains.com/go/) / JetBrains IDE features, Go language capabilities across versions 1.20–1.26, and real-world integration patterns. Each subdirectory is an independent, self-contained project.

## Table of Contents

- [Go Language Features](#go-language-features)
- [Web Applications](#web-applications)
- [Infrastructure & Cloud](#infrastructure--cloud)
- [IDE Feature Demonstrations](#ide-feature-demonstrations)
- [Testing & Quality](#testing--quality)
- [Other Demos](#other-demos)
- [Getting Started](#getting-started)
- [Project Index](#project-index)

---

## Go Language Features

### go125/ — Go 1.25 Feature Demos

A suite of focused examples, each in its own subdirectory:

| Subdirectory | What it shows |
|---|---|
| `flight-recorder` | Flight recorder / execution tracing |
| `garbage-collector` | GC enhancements and tuning |
| `go-version-m-json` | `go version -m` JSON output format |
| `goDoc` | GoDoc improvements |
| `ignore` | New `ignore` directive |
| `json2` | JSON v2 package |
| `repanic` | Re-panic behavior |
| `root-type` | Root type changes |
| `test-allocs-output` | `go test -test.allocationsoutput` flag |
| `testing-sync` | Testing synchronization helpers |
| `typeassert-function` | Type assertion function improvements |
| `vet-analyzers` | New `go vet` analyzers |

### go126/ — Go 1.26 Feature Demos

| Subdirectory | What it shows |
|---|---|
| `buffer-peek` | `bytes.Buffer` peek operations |
| `new-multi-handler` | Multi-handler HTTP routing |
| `the-new-variable` | New variable semantics |

### generic-type-aliases/ (Go 1.23)

Demonstrates generic type aliases introduced in Go 1.23.

---

## Web Applications

### todo-app/ — REST API with Vue.js Frontend

A classic to-do list app demonstrating GoLand features with a Go REST backend and Vue.js frontend.

- SQLite storage via `modernc.org/sqlite`
- Generic `Storage[T]` interface pattern
- Standard library HTTP routing with `r.PathValue()` (Go 1.22)

```bash
cd todo-app && go run main.go
# http://localhost:8080
```

### web-notebook/ — Note-taking App

A note-taking web application with SQLite persistence.

```bash
cd web-notebook && go run server/main.go
# http://127.0.0.1:8080
```

### grpc-demo/ — Dual gRPC + HTTP Server

Serves both gRPC (port 50051) and HTTP REST (port 8080) from a single binary using Protocol Buffers v3.

```bash
cd grpc-demo && go run main.go
# gRPC: localhost:50051  |  HTTP: localhost:8080
# OpenAPI spec: http://localhost:8080/grpc.yaml
```

### python-in-go/ — Go + Django Integration

Demonstrates managing a Django server's lifecycle from Go and communicating with its API.

```bash
cd python-in-go && go run main.go
```

### wasm-example/ — WebAssembly

A minimal Go program compiled to WebAssembly with a companion HTML page.

---

## Infrastructure & Cloud

### k8s/ — Kubernetes Client Demo

Comprehensive Kubernetes operations using the official `client-go` library against a local Minikube cluster.

```bash
cd k8s
make help          # list all targets
make deploy-all    # start minikube, build image, deploy
make status        # check deployment health
make logs          # tail application logs
make port-forward  # forward port 8080 to localhost
make delete        # tear down all resources
```

### terraform/ — Terraform SDK Integration (Go 1.26)

Programmatic Terraform state and plan parsing using `hashicorp/terraform-exec` and the Plugin SDK.

```bash
cd terraform && go run main.go
```

### go-terraform-demo/ — Terraform Basics Demo

A companion demo covering Terraform fundamentals (HCL, providers, state) with Go.

### dev-containers/ — Docker Compose Dev Environment

A `.devcontainer` configuration for VS Code and JetBrains, with a `docker-compose.yml` that sets up database services and runs post-create dependency installation.

---

## IDE Feature Demonstrations

| Project | GoLand feature highlighted |
|---|---|
| `debugging-features/` | Breakpoints, goroutine inspection, interface debugging |
| `inspections-lighting-talk/` | Code inspection quick-fixes |
| `data-flow-demo/` | Data flow analysis |
| `refactoring-demo/` | Inline variable, extract method, rename refactoring |
| `structure-view/` | Structure view and code navigation |
| `structural-search/` | Structural search and replace patterns |

> Many files contain intentional `//todo` comments as educational markers for IDE quick-fix demos — these are **not** real tasks to complete.

---

## Testing & Quality

### fuzz-testing/ (Go 1.22)

Fuzzing demonstrations using the built-in `go test -fuzz` engine.

```bash
cd fuzz-testing && go test -fuzz=FuzzReverse
```

### profiling-applications/ (Go 1.25)

Performance profiling and benchmarking examples.

```bash
cd profiling-applications && go test -bench=. -benchmem
```

### linters/

Linter configuration and demonstration patterns for use with GoLand's built-in linter integration.

### vulnerability-checker/ (Go 1.22)

Dependency vulnerability scanning and security analysis using `govulncheck`.

---

## Other Demos

### multidb-support/

Multi-database support demo using MySQL, PostgreSQL, and SQLite via Docker Compose. Includes Sakila sample schema and data dumps.

---

## Getting Started

Every project is independent. The general workflow is:

```bash
cd <project-name>
go mod download
go run main.go          # or: go run server/main.go
```

### Prerequisites

- **Go 1.22+** for most projects; individual `go.mod` files specify exact requirements
- **Go 1.26** for `terraform/`
- **Docker** for `dev-containers/` and `multidb-support/`
- **Minikube** for `k8s/`
- **protoc** to regenerate gRPC stubs in `grpc-demo/`
- **Python / Django** for `python-in-go/`

### Common Commands

```bash
# Format
go fmt ./...

# Vet
go vet ./...

# Test
go test -v ./...

# Tidy dependencies
go mod tidy
```

---

## Project Index

| Project | Go version | Description |
|---|---|---|
| `data-flow-demo` | 1.21 | Data flow analysis IDE demo |
| `debugging-features` | 1.24 | Debugger capabilities |
| `dev-containers` | 1.20 | Docker Compose dev environment |
| `fuzz-testing` | 1.22 | Fuzzing with `go test -fuzz` |
| `generic-type-aliases` | 1.23 | Generic type alias syntax |
| `go-terraform-demo` | 1.22 | Terraform fundamentals with Go |
| `go125/*` | 1.25 | Go 1.25 feature examples (12 demos) |
| `go126/*` | 1.26 | Go 1.26 feature examples (3 demos) |
| `grpc-demo` | 1.23 | gRPC + HTTP dual server |
| `inspections-lighting-talk` | 1.24 | Code inspection quick-fixes |
| `k8s` | 1.21 | Kubernetes client-go operations |
| `linters` | — | Linter configuration demos |
| `multidb-support` | — | MySQL / PostgreSQL / SQLite |
| `profiling-applications` | 1.25 | Benchmarking and profiling |
| `python-in-go` | 1.24 | Go managing a Django server |
| `refactoring-demo` | 1.22 | Refactoring IDE features |
| `structural-search` | — | Structural search patterns |
| `structure-view` | 1.22 | Structure view navigation |
| `terraform` | 1.26 | Terraform SDK integration |
| `todo-app` | 1.22 | REST API + Vue.js + SQLite |
| `vulnerability-checker` | 1.22 | Dependency security scanning |
| `wasm-example` | — | Go compiled to WebAssembly |
| `web-notebook` | 1.24 | Note-taking app with SQLite |
