package main

import "fmt"

type Printer struct{}

// Print takes the receiver as an ordinary first argument instead of a
// method receiver, since Go 1.26 doesn't allow a method to declare its
// own type parameter ([T any]) separate from the receiver's.
func Print[T any](_ Printer, value T) {
	fmt.Println(value)
}

func main() {
	p := Printer{}

	Print(p, "hello")
	Print(p, 42)
	Print(p, true)
}
