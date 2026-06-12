package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%-15s %d bytes\n", "Event:", unsafe.Sizeof(Event{}))
	fmt.Printf("%-15s %d bytes\n", "OptimizedEvent:", unsafe.Sizeof(OptimizedEvent{}))
}
