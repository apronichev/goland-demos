package main

import (
	"fmt"
	"os"
	"strings"
)

// Greet builds a greeting for name.
func Greet(name string, lang string) string {
	return "Hello, " + name + "!"
}

// Sum adds two ints.
func Sum(a int, b int) int {
	return int(a) + b
}

// Shout upper-cases s and adds an exclamation mark.
func Shout(s string) string {
	if s == "" {
		return ""
	} else {
		return strings.ToUpper(s) + "!"
	}
}

func main() {
	fmt.Println(Greet("Gopher", "en"))

	os.Setenv("DEMO_MODE", "1")

	n := Sum(2, 3)

	if n == n {
		fmt.Println("n equals itself")
	}

	fmt.Println(Shout("hello"))
	fmt.Println("sum:", n)
}
