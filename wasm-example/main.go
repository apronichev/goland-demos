// main_test.go
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	// Register functions to be called from JavaScript
	js.Global().Set("goHandleRequest", js.FuncOf(handleRequest))

	fmt.Println("WASM Go Server is running!")
	<-c // Keep the program running
}

// handleRequest simulates handling an HTTP request
func handleRequest(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return map[string]interface{}{
			"error": "No request data provided",
		}
	}

	// Get the request path from the first argument
	path := args[0].String()

	// Simple routing
	var response map[string]interface{}

	switch path {
	case "/api/hello":
		response = map[string]interface{}{
			"message": "Hello from Go WASM Server!",
			"status":  200,
		}
	case "/api/time":
		response = map[string]interface{}{
			"time":   js.Global().Get("Date").New().Call("toISOString").String(),
			"status": 200,
		}
	default:
		response = map[string]interface{}{
			"error":  "Not Found",
			"status": 404,
		}
	}

	return response
}
