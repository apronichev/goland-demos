package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// BAD: WaitGroup.Add inside goroutine
	for i := 0; i < 3; i++ {
		go func(n int) {
			wg.Add(1) // ❌ go vet: call to (*sync.WaitGroup).Add within goroutine
			defer wg.Done()
			fmt.Printf("Task %d\n", n)
		}(i)
	}
	wg.Wait()

	// BAD: Using fmt.Sprintf for host:port
	host := "example.com"
	port := 8080
	addr := fmt.Sprintf("%s:%d", host, port) // ❌ go vet: use net.JoinHostPort instead

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connection failed:", err)
	} else {
		conn.Close()
	}
}
