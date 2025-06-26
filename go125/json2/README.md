# Go 1.25: Experimental JSON v2 Package

A faster, more consistent JSON implementation available via build experiment.

The new implementation performs substantially better than the existing one under many scenarios. In general, encoding performance is at parity between the implementations and decoding is substantially faster in the new one. See the [repository](github.com/go-json-experiment/jsonbench) for more detailed analysis.

## Key Benefits
* Faster Decoding: often much faster than v1
* Better Error Messages: more descriptive and helpful
* Consistent Behavior: fewer edge cases and surprises
* Future-Proof: foundation for upcoming JSON features

## Enabling JSON v2

```bash
# Build with experiment
GOEXPERIMENT=jsonv2 go build

# Run with experiment
GOEXPERIMENT=jsonv2 go run main.go

# Test with experiment
GOEXPERIMENT=jsonv2 go test ./...

