package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/demo/todo-app/internal/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	db        *sql.DB
	templates *template.Template
}

func New(db *sql.DB) *Handler {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	return &Handler{
		db:        db,
		templates: tmpl,
	}
}

type PageData struct {
	Todos   []models.Todo
	DBError string
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading all todos")
	data := PageData{
		Todos: []models.Todo{},
	}

	if h.db == nil {
		data.DBError = "Database unavailable"
		h.templates.ExecuteTemplate(w, "index.html", data)
		return
	}

	rows, err := h.db.Query("SELECT id, title, completed FROM todos ORDER BY id")
	if err != nil {
		log.Printf("Failed to load todos from DB: %v", err)
		data.DBError = "Database unavailable"
		h.templates.ExecuteTemplate(w, "index.html", data)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			log.Printf("Error scanning todo: %v", err)
			continue
		}
		data.Todos = append(data.Todos, todo)
	}

	h.templates.ExecuteTemplate(w, "index.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form: %v", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	log.Printf("Adding new todo: %s", title)

	if h.db == nil {
		log.Println("Cannot add todo - database unavailable")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_, err := h.db.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", title, false)
	if err != nil {
		log.Printf("Failed to add todo: %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Toggle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("Invalid todo ID: %s", idStr)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if h.db == nil {
		log.Println("Cannot toggle todo - database unavailable")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_, err = h.db.Exec("UPDATE todos SET completed = NOT completed WHERE id = $1", id)
	if err != nil {
		log.Printf("Failed to toggle todo id=%d: %v", id, err)
	} else {
		log.Printf("Toggled todo id=%d", id)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("Invalid todo ID: %s", idStr)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	log.Printf("Deleting todo id=%d", id)

	if h.db == nil {
		log.Println("Cannot delete todo - database unavailable")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_, err = h.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		log.Printf("Failed to delete todo id=%d: %v", id, err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
