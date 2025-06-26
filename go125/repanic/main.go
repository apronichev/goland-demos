package main

import (
	"fmt"
	"log"
)

// Switch between Go 1.24 and Go 1.25 to see the difference in panicking behavior
func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from: %v", r)
			// Repanic with the same value
			panic(r)
		}
	}()

	// Original panic
	panic("SOMETHING WENT WRONG")
}

func middlewareLayer() {
	defer func() {
		if r := recover(); r != nil {
			// Log and repanic (common in middleware/frameworks)
			fmt.Printf("Middleware caught: %v\n", r)
			panic(r) // Repanic to let caller handle
		}
	}()

	riskyOperation()
}

func main() {
	fmt.Println("Starting program...")
	middlewareLayer()
}
