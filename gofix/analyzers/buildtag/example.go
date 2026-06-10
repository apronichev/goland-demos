//go:build linux || darwin

// `buildtag`: a vet-style analyzer (no auto-fix) that diagnoses malformed
// build directives — wrong spelling (`//go:buil`), misplaced lines (after the
// package clause), inconsistent `//go:build` vs `// +build` constraints, etc.
//
// The directive above is well-formed, so this file is the "passing" baseline.
// To see the analyzer fire, introduce a typo like `//go:buld linux` and run:
//
//	go vet -vettool=$(go env GOTOOLDIR)/fix ./analyzers/buildtag/...
package buildtag

const Tag = "ok"
