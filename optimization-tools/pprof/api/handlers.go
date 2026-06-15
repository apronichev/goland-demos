package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/goland-demos/optimization-tools/pprof/generator"
	"github.com/goland-demos/optimization-tools/pprof/index"
)

const (
	userCount = 5_000_000
	seed      = 42
)

type Server struct {
	idx *index.Index
}

func NewServer(idx *index.Index) *Server {
	return &Server{idx: idx}
}

func (s *Server) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /stats", s.stats)
	mux.HandleFunc("GET /search", s.search)
	mux.HandleFunc("GET /users/{id}", s.getUser)
	mux.HandleFunc("POST /reindex", s.reindex)
}

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) stats(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, s.idx.Stats())
}

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "missing q", http.StatusBadRequest)
		return
	}
	results := s.idx.Search(q)
	writeJSON(w, http.StatusOK, map[string]any{
		"query": q,
		"count": len(results),
		"users": results,
	})
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	u, ok := s.idx.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, u)
}

func (s *Server) reindex(w http.ResponseWriter, _ *http.Request) {
	log.Printf("reindexing users...")
	users := generator.Generate(userCount, seed)
	s.idx.Build(users)
	writeJSON(w, http.StatusOK, s.idx.Stats())
	log.Printf("reindexing complete")
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
