package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	{ID: 4, Name: "David", Email: "david@example.com"},
	{ID: 5, Name: "Eve", Email: "eve@example.com"},
	{ID: 6, Name: "Frank", Email: "frank@example.com"},
	{ID: 7, Name: "Grace", Email: "grace@example.com"},
	{ID: 8, Name: "Hannah", Email: "hannah@example.com"},
	{ID: 9, Name: "Ivan", Email: "ivan@example.com"},
	{ID: 10, Name: "Juliet", Email: "juliet@example.com"},
	{ID: 11, Name: "Karl", Email: "karl@example.com"},
	{ID: 12, Name: "Lily", Email: "lily@example.com"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, user := range users {
		if fmt.Sprint(user.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
