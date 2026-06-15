package main

import (
	"log"
	"net/http"

	"github.com/goland-demos/optimization-tools/pprof/api"
	"github.com/goland-demos/optimization-tools/pprof/index"
)

const addr = ":8080"

func main() {
	idx := index.New()

	mux := http.NewServeMux()
	api.NewServer(idx).Routes(mux)

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
