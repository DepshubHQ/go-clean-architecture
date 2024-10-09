package main

import "fmt"

type TodoStorageInMemory struct {
	todos []Todo
}

func NewTodoStorageInMemory() *TodoStorageInMemory {
	return &TodoStorageInMemory{
		todos: []Todo{},
	}
}

func (s *TodoStorageInMemory) Create(todo Todo) error {
	s.todos = append(s.todos, todo)
	return nil
}

func (s *TodoStorageInMemory) Update(todo Todo) error {
	for i, t := range s.todos {
		if t.ID == todo.ID {
			s.todos[i] = todo
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (s *TodoStorageInMemory) Delete(id uint) error {
	for i, t := range s.todos {
		if t.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (s *TodoStorageInMemory) List() ([]Todo, error) {
	return s.todos, nil
}
