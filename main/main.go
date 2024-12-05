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
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	})
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/todos/")
		w.Header().Set("Content-Type", "application/json")
		idStr, _ := strconv.Atoi(id)
		json.NewEncoder(w).Encode(todos[idStr])
	})
	http.ListenAndServe(":8080", nil)
}
