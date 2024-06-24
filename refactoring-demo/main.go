package main

import (
	"fmt"
)

func main() {
	//todo Step 2. pass for the 'sum' the following value:'[]int{3, 4, 5}'
	sum := add(3, 4)
	fmt.Println("Sum:", sum)
	//todo Step 6. Rename 'greet' to 'sayHi'. Navigate to it in 'greet.go'.
	greeting := greet("World")
	fmt.Println(greeting)
}

// todo Step 1. Change the signature of the add function to accept 'numbers []int' and return their sum.
// todo iterate through the slice, perform 'sum += num' and 'return sum'.
func add(a int, b int) int {
	return a + b
	//todo Step 4. Extract/inline the loop inside the 'add' function to a new method called 'sumNumbers'.
}

// todo Step 5. Move the 'greet' function to 'greet.go'
func greet(name string) string {
	//todo Step 3. Extract/inline 'message := "Hello, " + name + "!"' to a variable
	return "Hello, " + name + "!"
}
