package infra

import (
	"database/sql"
	"errors"
	"go-api-structure/inputs"
	"go-api-structure/model"

	_ "github.com/mattn/go-sqlite3"
)

type SQLTodoRepository struct {
	db *sql.DB
}

func NewSQLTodoRepository(db *sql.DB) *SQLTodoRepository {
	return &SQLTodoRepository{db: db}
}

func (r *SQLTodoRepository) Create(input inputs.CreateTodoInput) (*model.Todo, error) {
	result, err := r.db.Exec(
		"INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)",
		input.Title, input.Description, false,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:          int(id),
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}, nil
}

func (r *SQLTodoRepository) Remove(id int) error {
	result, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (r *SQLTodoRepository) Update(todo model.Todo) error {
	_, err := r.db.Exec(
		"UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?",
		todo.Title, todo.Description, todo.Completed, todo.ID,
	)
	return err
}

func (r *SQLTodoRepository) GetByID(id int) (*model.Todo, error) {
	row := r.db.QueryRow("SELECT id, title, description, completed FROM todos WHERE id = ?", id)
	var todo model.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err == sql.ErrNoRows {
		return nil, errors.New("todo not found")
	}
	return &todo, err
}
