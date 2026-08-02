// Package greeting builds localized greetings.
//
// This package exists to showcase GoLand MCP's semantic, type-aware
// refactoring. Try `rename_refactoring` on Greeter, Greet, or Default:
// every call site below AND in cmd/greet updates safely, while the
// unrelated Greet function in the root main.go is left untouched —
// something plain text search (grep/sed) cannot get right.
package greeting

import "strings"

// Greeter renders greetings using a fixed prefix.
type Greeter struct {
	Prefix string
}

// Greet returns the greeting for name (e.g. "Hello, Gopher!").
func (g Greeter) Greet(name string) string {
	return g.Prefix + ", " + name + "!"
}

// Default returns the default English greeter.
func Default() Greeter {
	return Greeter{Prefix: "Hello"}
}

// Shout upper-cases a greeting for emphasis.
func Shout(s string) string {
	return strings.ToUpper(s)
}
