# Go 1.25: New `ignore` Directive in go.mod

The `ignore` directive allows you to exclude specific module versions from dependency resolution.

## Basic Usage

```go
module example.com/myapp

go 1.25

require (
    github.com/lib/pq v1.10.9
)

// Ignore a specific version with security issues
ignore github.com/lib/pq v1.10.7

// Ignore a range of versions
ignore github.com/some/module v1.0.0
ignore github.com/some/module v1.1.0