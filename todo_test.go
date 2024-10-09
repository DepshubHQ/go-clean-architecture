package main

import "testing"

// Test for Create method
func TestTodoService_Create(t *testing.T) {
	mockStorage := &TodoStorageInMemory{}
	service := NewTaskService(mockStorage)

	todo := Todo{ID: 1, Title: "Test Todo", Status: TodoState}

	err := service.Create(todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(mockStorage.todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(mockStorage.todos))
	}

	if mockStorage.todos[0] != todo {
		t.Fatalf("expected todo %v, got %v", todo, mockStorage.todos[0])
	}
}

// Test for Update method
func TestTodoService_Update(t *testing.T) {
	mockStorage := &TodoStorageInMemory{todos: []Todo{{ID: 1, Title: "Old Todo", Status: TodoState}}}
	service := NewTaskService(mockStorage)

	updatedTodo := Todo{ID: 1, Title: "Updated Todo", Status: DoingState}

	err := service.Update(updatedTodo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mockStorage.todos[0] != updatedTodo {
		t.Fatalf("expected todo %v, got %v", updatedTodo, mockStorage.todos[0])
	}
}

// Test for Delete method
func TestTodoService_Delete(t *testing.T) {
	mockStorage := &TodoStorageInMemory{todos: []Todo{{ID: 1, Title: "Test Todo", Status: TodoState}}}
	service := NewTaskService(mockStorage)

	err := service.Delete(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(mockStorage.todos) != 0 {
		t.Fatalf("expected 0 todos, got %d", len(mockStorage.todos))
	}
}

// Test for List method
func TestTodoService_List(t *testing.T) {
	mockStorage := &TodoStorageInMemory{todos: []Todo{
		{ID: 1, Title: "Todo 1", Status: TodoState},
		{ID: 2, Title: "Todo 2", Status: DoingState},
	}}
	service := NewTaskService(mockStorage)

	todos, err := service.List()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(todos) != 2 {
		t.Fatalf("expected 2 todos, got %d", len(todos))
	}
}
