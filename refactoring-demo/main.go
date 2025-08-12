package main

import (
	"fmt"
)

type Speaker interface {
	Present(name string) string
	Say(name string) string
	Inspire() string
}

type Persona struct {
	firstName string
	lastName  string
}

func (p Persona) Present(name string) string {
	//TODO implement me
	panic("implement me")
}

func (p Persona) Say(name string) string {
	//TODO implement me
	panic("implement me")
}

func (p Persona) Inspire() string {
	//TODO implement me
	panic("implement me")
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
