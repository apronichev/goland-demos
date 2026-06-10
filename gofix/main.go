// Package main is a stub. The interesting code lives under ./analyzers/<name>/.
//
// Each subdirectory contains a tiny package whose source intentionally triggers
// exactly one of the analyzers shipped with `go tool fix`. Run:
//
//	go fix -diff ./...                 # preview every fix without mutating files
//	go fix ./...                       # apply every fix
//	go fix -diff -minmax ./...         # preview a single analyzer
//	go tool fix help                   # list all analyzers
//	go tool fix help minmax            # details for one analyzer
package main

func main() {}
