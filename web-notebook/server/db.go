package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS notes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL
    );`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}

func addNote(content string) error {
	_, err := db.Exec("INSERT INTO notes (content) VALUES (?)", content)
	return err
}

func getNotes() ([]string, error) {
	rows, err := db.Query("SELECT content FROM notes")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var notes []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}
		notes = append(notes, content)
	}
	return notes, nil
}
