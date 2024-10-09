package main

import (
	"fmt"
)

type TodoStorage interface {
	Create(todo Todo) error
	Update(todo Todo) error
	Delete(id uint) error
	List() ([]Todo, error)
}

type TodoService struct {
	storage TodoStorage
}

func NewTodoService(storage TodoStorage) *TodoService {
	return &TodoService{storage: storage}
}

const MaxTitleLength = 5

func (s *TodoService) Create(todo Todo) error {
	if len(todo.Title) < MaxTitleLength {
		return fmt.Errorf("title should be at least %d characters", MaxTitleLength)
	}
	return s.storage.Create(todo)
}

func (s *TodoService) Update(todo Todo) error {
	return s.storage.Update(todo)
}

func (s *TodoService) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *TodoService) List() ([]Todo, error) {
	return s.storage.List()
}
