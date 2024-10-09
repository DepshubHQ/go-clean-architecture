package main

import (
	"encoding/json"
	"net/http"
)

type TodoController struct {
	service *TodoService
}

func NewTodoController(service *TodoService) *TodoController {
	return &TodoController{service}
}

func (c *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	todo, err := c.parseTodoRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := c.service.Create(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	todo, err := c.parseTodoRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := c.service.Update(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	todo, err := c.parseTodoRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := c.service.Delete(todo.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *TodoController) List(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (c *TodoController) parseTodoRequest(r *http.Request) (Todo, error) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return Todo{}, err
	}
	return todo, nil
}
