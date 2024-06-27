package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo-app/model"
	"todo-app/storage"
)

// todo invoke "implement missing methods" quick fix on model.TodoItem
var todoStorage storage.Storage[*model.TodoItem] = storage.NewInMemoryStorage[*model.TodoItem]()

func main() {
	todoStorage.Put(&model.TodoItem{
		Title: "Todo 1",
		Done:  false,
		Due:   "23.07.2024",
	})
	todoStorage.Put(&model.TodoItem{
		Title: "Todo 2",
		Done:  true,
		Due:   "01.01.2023",
	})
	// todo invoke "inline variable" on fs
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	// todo invoke "Generate request in HTTP Client"
	http.HandleFunc("/todos", todoHandler)
	http.HandleFunc("/todos/{id}", todoHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.PathValue("id") != "" {
			loadTodo(w, r)
			return
		}
		listTodos(w, r)
	case http.MethodPost:
		createOrUpdateTodo(w, r)
	case http.MethodDelete:
		deleteTodo(w, r)
		// todo type case and see FLCC completion for "case http.MethodPut:"
	}
}

func loadTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var todo = todoStorage.Get(id)
	if todo == nil {
		// todo FLCC works as a charm here!
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(todo)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// todo FLCC suggests here complete line!
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todoStorage.Remove(id)
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	todos := todoStorage.GetAll()
	w.Header().Set("Content-Type", "application/json")
	// todo invoke "handle error" quick fix on Encode()
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createOrUpdateTodo(w http.ResponseWriter, r *http.Request) {
	item := model.TodoItem{}
	// todo invoke "prepend with &" quick fix on item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		// todo invoke extract variable for "something went wrong"
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	// todo invoke "create method `Validate`" quick fix on Validate()
	err = item.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todoStorage.Put(&item)
}
