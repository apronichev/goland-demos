package main

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-app/model"
	"todo-app/storage"
)

// todo invoke "implement missing methods" quick fix on model.TodoItem
var todoStorage storage.Storage[*model.TodoItem] = storage.NewInMemoryStorage[*model.TodoItem]()

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTodos(w, r)
	case http.MethodPost:
		createOrUpdateTodo(w, r)
		// todo type case and see FLCC completion for "case http.MethodPut:"
	}
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
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	todoStorage.Put(&item)
}

func main() {
	// todo invoke "Generate request in HTTP Client"
	http.HandleFunc("/", todoHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
