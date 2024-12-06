package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []todo{
	{ID: "1", Title: "Fill up on gas", Description: "We will be going on a roadtrip and there is no gas in the car"},
	{ID: "2", Title: "Do the Java homework", Description: "In a few days we will have the final exam in Java. We must be prepared"},
	{ID: "3", Title: "Cook a healthy meal", Description: "We must get all the protein our body needs, so cook something healthy"},
	{ID: "4", Title: "Write a RESTful API in GO", Description: "Go is a fun language so we must learn it"},
	{ID: "5", Title: "Clean the house", Description: "We will have guests coming over so let's clean the house"},
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return true
	}
	return false
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if methodNotAllowed(w, r, http.MethodGet) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	})
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if methodNotAllowed(w, r, http.MethodGet) {
			return
		}
		id := strings.TrimPrefix(r.URL.Path, "/todos/")
		w.Header().Set("Content-Type", "application/json")
		idInt, err := strconv.Atoi(id)

		if len(todos) <= idInt || idInt < 0 || err != nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
		} else {
			for _, todo := range todos {
				if todo.ID == id {
					json.NewEncoder(w).Encode(todo)
					return
				}
			}
		}
	})
	http.HandleFunc("/todos/delete", func(w http.ResponseWriter, r *http.Request) {
		if methodNotAllowed(w, r, http.MethodDelete) {
			return
		}
		id := r.URL.Query().Get("id")

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				json.NewEncoder(w).Encode("Todo with ID " + id + " deleted")
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)
	})
	http.HandleFunc("/todos/post", func(w http.ResponseWriter, r *http.Request) {
		if methodNotAllowed(w, r, http.MethodPost) {
			return
		}
		title := r.URL.Query().Get("title")
		if title == "" {
			http.Error(w, "Please provide a title", http.StatusBadRequest)
			return
		}
		description := r.URL.Query().Get("description")
		if description == "" {
			http.Error(w, "Please provide a description", http.StatusBadRequest)
			return
		}
		var newTodo todo
		lastId, _ := strconv.Atoi(todos[len(todos)-1].ID)
		newTodo.ID = strconv.Itoa(lastId + 1)
		newTodo.Title = title
		newTodo.Description = description
		todos = append(todos, newTodo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTodo)
	})
	http.HandleFunc("/todos/put", func(w http.ResponseWriter, r *http.Request) {
		if methodNotAllowed(w, r, http.MethodPut) {
			return
		}
		id := r.URL.Query().Get("id")
		updatedTitle := r.URL.Query().Get("title")
		updatedDescription := r.URL.Query().Get("description")

		for i, _ := range todos {
			if todos[i].ID == id {
				todos[i].Title = updatedTitle
				todos[i].Description = updatedDescription
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(todos[i])
				return
			}
		}
		http.Error(w, "Not found", http.StatusNotFound)
	})
	http.ListenAndServe(":8080", nil)
}
