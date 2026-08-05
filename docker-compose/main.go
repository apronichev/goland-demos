package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type server struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

// env returns the value of the environment variable, or fallback when unset.
func env(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func main() {
	ctx := context.Background()

	dsn := env("DATABASE_URL", "postgres://demo:demo@postgres:5432/demo?sslmode=disable")
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: env("REDIS_ADDR", "redis:6379"),
	})
	defer rdb.Close()

	srv := &server{db: db, rdb: rdb}

	mux := http.NewServeMux()
	mux.HandleFunc("/", srv.handleIndex)
	mux.HandleFunc("/healthz", srv.handleHealth)
	mux.HandleFunc("/users", srv.handleUsers)
	mux.HandleFunc("/visits", srv.handleVisits)

	addr := ":" + env("PORT", "8080")
	log.Printf("listening on %s — try http://localhost%s", addr, addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<!doctype html>
<html>
<head><title>Docker Compose demo</title></head>
<body style="font-family: sans-serif; max-width: 40rem; margin: 3rem auto;">
  <h1>Docker Compose demo</h1>
  <p>A multi-container Go app wired to Postgres and Redis.</p>
  <ul>
    <li><a href="/users">/users</a> — list users from Postgres</li>
    <li><a href="/visits">/visits</a> — increment a Redis visit counter</li>
    <li><a href="/healthz">/healthz</a> — liveness probe</li>
  </ul>
</body>
</html>`)
}

// handleHealth reports whether both backing services are reachable.
func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	status := map[string]string{"postgres": "ok", "redis": "ok"}
	code := http.StatusOK

	if err := s.db.Ping(ctx); err != nil {
		status["postgres"] = err.Error()
		code = http.StatusServiceUnavailable
	}
	if err := s.rdb.Ping(ctx).Err(); err != nil {
		status["redis"] = err.Error()
		code = http.StatusServiceUnavailable
	}

	writeJSON(w, code, status)
}

// handleUsers returns every row from the users table seeded by db/init.sql.
func (s *server) handleUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, "SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	type user struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	users := make([]user, 0)
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, users)
}

// handleVisits increments and returns a counter stored in Redis.
func (s *server) handleVisits(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	count, err := s.rdb.Incr(ctx, "visits").Result()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]int64{"visits": count})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
