package main

import (
	"fmt"
)

type Persona struct {
	firstName string
	lastName  string
}

func main() {
	sum := add(3, 4)
	speaker1 := Persona{"John", "Doe"}
	greeting := greet(speaker1.lastName)
	fmt.Println(greeting)
	fmt.Printf("Answer is %v.", sum)
}

func add(a int, b int) int {
	return a + b
}

func greet(name string) string {
	return "Hello, I am mr. " + name + ". Give me a list of numbers and I will summarize them."
}
