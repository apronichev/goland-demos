package main

// Package main is the entry point

import "fmt"

// Import fmt to use Println to show output

func FibRecursive(n int) uint64 {
	// Function Creation to accept integer till which the Fibonacci series runs
	if n == 0 {
		return 0
		// Base case for the recursive call
	} else if n == 1 {
		return 1
		// Base case for the recursive call
	} else {
		return FibRecursive(n-1) + FibRecursive(n-2)
		// Recursive call for finding the Fibonacci number
	}
}

func main() {
	// Function call to print out the nth term of the Fibonacci series
	fmt.Println(FibRecursive(3))
}
