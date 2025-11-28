package main

import (
	"fmt"
)

// Persona represents a person with first and last names
type Persona struct {
	FirstName string
	LastName  string
}

// NewPersona creates a new Persona instance
func NewPersona(firstName, lastName string) *Persona {
	return &Persona{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// Greet returns a greeting message from this persona
func (p *Persona) Greet() string {
	return fmt.Sprintf("Hello, I am Mr. %s. Give me a list of numbers and I will summarize them.", p.LastName)
}

func main() {
	speaker := NewPersona("John", "Doe")

	fmt.Println(speaker.Greet())

	sum := add(3, 4)
	fmt.Printf("Answer is %v.\n", sum)
}

func add(a, b int) int {
	return a + b
}
