package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var db *sql.DB

func initDB() {
	var err error
	dbPath := filepath.Join("./data", "notes.db") // Store in a 'data' subdirectory
	os.MkdirAll("./data", 0755)                   // Create the directory if it doesn't exist

	log.Printf("Database path: %s", dbPath)
	db, err = sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS notes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL
    );`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Database initialized successfully")
}

func addNote(content string) error {
	result, err := db.Exec("INSERT INTO notes (content) VALUES (?)", content)
	if err != nil {
		log.Printf("Error adding note: %v", err)
		return err
	}

	id, _ := result.LastInsertId()
	log.Printf("Added note with ID: %d", id)
	return nil
}

func getNotes() ([]Note, error) {
	rows, err := db.Query("SELECT id, content FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
