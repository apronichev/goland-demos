package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://go.dev/blog")
	defer resp.Body.Close()
	if err != nil {
		return
	}
	fmt.Printf("Status code: %d", resp.StatusCode)
}
