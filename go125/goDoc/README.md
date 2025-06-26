# Go 1.25: Documentation Server with `go doc -http`

The new `-http` flag starts a local web server for browsing Go documentation.

## Basic Usage

```bash
# Start doc server on default port (6060)
go doc -http=:6060

# Custom port
go doc -http=:8080

# With specific package
go doc -http=:8080 encoding/json
