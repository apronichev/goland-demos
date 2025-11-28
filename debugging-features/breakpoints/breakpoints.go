package main

import (
	"fmt"
	"time"
)

// Simple calculator that processes numbers
func main() {
	fmt.Println("=== Calculator Demo - Use Non-Suspending Breakpoints ===\n")

	numbers := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

	total := 0
	for _, num := range numbers {
		result := processNumber(num)
		total += result
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Printf("\nFinal Total: %d\n", total)
}

// processNumber performs calculations
func processNumber(n int) int {
	doubled := n * 2
	if doubled > 50 {
		return doubled + 10
	}

	return doubled
}
