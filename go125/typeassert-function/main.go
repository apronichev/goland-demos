package main

import (
	"fmt"
	"reflect"
	"time"
)

type MysteryBox struct {
	content interface{}
}

type Person struct {
	Name string
	Age  int
}

func main() {
	boxes := []MysteryBox{
		{content: "Hello, World!"},
		{content: 42},
		{content: Person{Name: "Alice", Age: 30}},
		{content: 3.14159},
	}

	fmt.Println("Opening mystery boxes the OLD way (before Go 1.25):")
	openBoxesOldWay(boxes)

	fmt.Println("\nOpening mystery boxes the NEW way (Go 1.25+):")
	openBoxesNewWay(boxes)

	fmt.Println("\nPerformance Comparison:")
	demonstratePerformance()
}

// The OLD way - using Interface() then type assertion
func openBoxesOldWay(boxes []MysteryBox) {
	for i, box := range boxes {
		v := reflect.ValueOf(box.content)

		// Old way: First convert to interface{}, then do type assertion
		// This creates an unnecessary copy in memory!
		switch content := v.Interface().(type) {
		case string:
			fmt.Printf("  Box %d: Found a message: %q\n", i+1, content)
		case int:
			fmt.Printf("  Box %d: Found a number: %d\n", i+1, content)
		case Person:
			fmt.Printf("  Box %d: Found a person: %s (age %d)\n", i+1, content.Name, content.Age)
		case float64:
			fmt.Printf("  Box %d: Found a decimal: %.2f\n", i+1, content)
		}
	}
}

// The NEW way - using TypeAssert directly
func openBoxesNewWay(boxes []MysteryBox) {
	for i, box := range boxes {
		v := reflect.ValueOf(box.content)

		if str, ok := reflect.TypeAssert[string](v); ok {
			fmt.Printf("  Box %d: Found a message: %q\n", i+1, str)
		} else if num, ok := reflect.TypeAssert[int](v); ok {
			fmt.Printf("  Box %d: Found a number: %d\n", i+1, num)
		} else if person, ok := reflect.TypeAssert[Person](v); ok {
			fmt.Printf("  Box %d: Found a person: %s (age %d)\n", i+1, person.Name, person.Age)
		} else if decimal, ok := reflect.TypeAssert[float64](v); ok {
			fmt.Printf("  Box %d: Found a decimal: %.2f\n", i+1, decimal)
		}
	}
}

// Demonstrate the performance benefit
func demonstratePerformance() {
	// Create a large slice of values to process
	const iterations = 1000000
	testValue := Person{Name: "Test", Age: 25}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		v := reflect.ValueOf(testValue)
		_ = v.Interface().(Person)
	}
	oldDuration := time.Since(start)

	start = time.Now()
	for i := 0; i < iterations; i++ {
		v := reflect.ValueOf(testValue)
		_, _ = reflect.TypeAssert[Person](v)
	}
	newDuration := time.Since(start)

	fmt.Printf("  🐌 Old way took: %v\n", oldDuration)
	fmt.Printf("  🚀 New way took: %v\n", newDuration)
	fmt.Printf("  📊 Processing %d items\n", iterations)
}
