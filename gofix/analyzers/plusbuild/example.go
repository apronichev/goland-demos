//go:build linux || darwin
// +build linux darwin

// `plusbuild`: the `// +build` line is the pre-Go-1.18 syntax and is now
// redundant whenever a matching `//go:build` line is present.
package plusbuild

const Platform = "unix-ish"
