package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	config := trace.FlightRecorderConfig{}

	// Create and start the flight recorder
	recorder := trace.NewFlightRecorder(config)
	recorder.Start()
	defer recorder.Stop()

	doWork()

	// Simulate a critical event - capture trace snapshot
	fmt.Println("Critical event occurred! Capturing trace...")

	file, err := os.Create("trace_snapshot.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := recorder.WriteTo(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Trace snapshot saved to trace_snapshot.out (%d bytes)\n", n)
}

func doWork() {
	for i := 0; i < 1000; i++ {
		go func(n int) {
			time.Sleep(time.Millisecond)
			_ = n * n
		}(i)
	}
	time.Sleep(2 * time.Second)
}
