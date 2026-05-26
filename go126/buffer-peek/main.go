package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf1 := bytes.NewBuffer([]byte("Hello, Go 1.26!"))

	first := make([]byte, 5)
	_, err := buf1.Read(first)
	// first, err := buf2.Peek(5)
	if err != nil {
		return
	}
	fmt.Println("Read:", string(first))
	fmt.Println("Remaining:", buf1.String())
}
