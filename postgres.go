package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type TodoStoragePostgres struct {
	db *sql.DB
}

func NewTodoStoragePostgres() *TodoStoragePostgres {
	db, err := sql.Open("postgres", "postgres://postgres@localhost/todos?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &TodoStoragePostgres{
		db: db,
	}
}

func (s *TodoStoragePostgres) Create(todo Todo) error {
	_, err := s.db.Exec("INSERT INTO todos (id, title, status) VALUES ($1, $2, $3)", todo.ID, todo.Title, todo.Status)
	return err
}

func (s *TodoStoragePostgres) Update(todo Todo) error {
	_, err := s.db.Exec("UPDATE todos SET title = $1, status = $2 WHERE id = $3", todo.Title, todo.Status, todo.ID)
	return err
}

func (s *TodoStoragePostgres) Delete(id uint) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (s *TodoStoragePostgres) List() ([]Todo, error) {
	rows, err := s.db.Query("SELECT id, title, status FROM todos")
	if err != nil {
		return nil, err
	}

	todos := []Todo{}
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status)

		if err != nil {
			log.Println(err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
