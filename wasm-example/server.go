// server.go
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Port to serve on")
	flag.Parse()

	// Serve the WebAssembly file with the correct MIME type
	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, "./main.wasm")
	})

	// Serve static files
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Printf("Server started at http://localhost:%s/wasm-example/", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
