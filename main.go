package main

import (
	"log"
	"net/http"
)

func main() {
	storage := NewTodoStoragePostgres()
	service := NewTodoService(storage)
	controller := NewTodoController(service)

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		switch r.Method {
		case http.MethodPost:
			controller.Create(w, r)
		case http.MethodPut:
			controller.Update(w, r)
		case http.MethodDelete:
			controller.Delete(w, r)
		case http.MethodGet:
			controller.List(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
