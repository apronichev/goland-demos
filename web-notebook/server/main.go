package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	initDB()
	http.HandleFunc("/notes", notesHandler)
	http.Handle("/", http.FileServer(http.Dir("./client")))

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		notes, err := getNotes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(notes)
	case "POST":
		var note struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := addNote(note.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
