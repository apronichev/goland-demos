// Package greetings provides simple functions for generating greeting
// and farewell messages for a given name.
//
// This package can be useful for demonstrating how to document Go packages
// and how to expose functions with proper comments.
//
// Example usage:
//
//	msg := greetings.Hello("Alice")
//	fmt.Println(msg)
//
// Supported functions:
//   - Hello: Returns a friendly greeting.
//   - Goodbye: Returns a polite farewell.
package greetings

import "fmt"

// Hello returns a personalized greeting message for the given name.
//
// Notes:
//  1. The name is inserted directly into the message.
//  2. If an empty string is passed, the result will include just the punctuation.
//
// Example:
//
//	Hello("Alice") returns "Hello, Alice!"
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Goodbye returns a personalized farewell message for the given name.
//
// Notes:
//  1. This function works similarly to Hello but with a goodbye format.
//  2. Suitable for logging, CLI output, or friendly terminal responses.
//
// Example:
//
//	Goodbye("Bob") returns "Goodbye, Bob!"
func Goodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}
