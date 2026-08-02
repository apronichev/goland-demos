package main

import (
	"fmt"
	"net/http"
	"time"
)

// doWork simulates a slow operation (a DB query, an RPC, ...).
func doWork() string {
	time.Sleep(2 * time.Second)
	return "work finished"
}

// processHandler kicks off some background processing when the endpoint is
// hit. It looks reasonable, but it leaks a goroutine on every request.
//
// The worker sends its result on an unbuffered channel. The handler only
// waits 500ms before giving up and returning. Once it returns, `result`
// becomes unreachable, so the worker is stuck forever on `result <- ...`
// with nobody left to receive: an orphaned goroutine leak.
func processHandler(w http.ResponseWriter, r *http.Request) {
	result := make(chan string) // unbuffered

	go func() {
		result <- doWork() // blocks forever once the handler has returned
	}()

	select {
	case res := <-result:
		fmt.Fprintln(w, res)
	case <-time.After(500 * time.Millisecond):
		// Bug: we bail out on timeout but never drain `result`.
		http.Error(w, "timed out", http.StatusGatewayTimeout)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /process", processHandler)

	fmt.Println("Listening on http://localhost:8080...\n" +
		"Hit http://localhost:8080/process a few times,\n" +
		"then open the goroutine leak profile in the GoLand profiler.")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("server error:", err)
	}
}
