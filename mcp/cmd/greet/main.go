// Command greet consumes the greeting package across a package boundary,
// so a rename of Greeter/Greet/Default in ../../greeting updates this
// call site too — the cross-package demo for GoLand MCP rename_refactoring.
package main

import (
	"fmt"

	"mcp/greeting"
)

func main() {
	g := greeting.Default()
	fmt.Println(g.Greet("Gopher"))
	fmt.Println(greeting.Shout(g.Greet("world")))
}
