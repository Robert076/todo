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

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	})
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		id := strings.TrimPrefix(r.URL.Path, "/todos/")
		w.Header().Set("Content-Type", "application/json")
		idInt, err := strconv.Atoi(id)

		if len(todos) <= idInt || idInt < 0 || err != nil {
			json.NewEncoder(w).Encode(nil)
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
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)

		if len(todos) <= idInt || idInt < 0 || err != nil {
			json.NewEncoder(w).Encode(nil)
		} else {
			json.NewEncoder(w).Encode("Deleting id " + id)
			todos = append(todos[:idInt], todos[idInt+1:]...)
		}
	})
	http.HandleFunc("/todos/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		title := r.URL.Query().Get("title")
		if title == "" {
			json.NewEncoder(w).Encode("Please provide a title")
			return
		}
		description := r.URL.Query().Get("description")
		if description == "" {
			json.NewEncoder(w).Encode("Please provide a description")
		}
		var newTodo todo
		lastId, _ := strconv.Atoi(todos[len(todos)-1].ID)
		newTodo.ID = strconv.Itoa(lastId + 1)
		newTodo.Title = title
		newTodo.Description = description
		todos = append(todos, newTodo)
	})
	http.ListenAndServe(":8080", nil)
}
