package infra

import (
	"database/sql"
	"errors"
	"go-api-structure/views"

	_ "github.com/mattn/go-sqlite3"
)

type SQLQueryRepository struct {
	db *sql.DB
}

func NewSQLQueryRepository(db *sql.DB) *SQLQueryRepository {
	return &SQLQueryRepository{db: db}
}

func (r *SQLQueryRepository) GetAll() ([]views.TodoView, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	viewsList := []views.TodoView{}
	for rows.Next() {
		var todo views.TodoView
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		viewsList = append(viewsList, todo)
	}

	return viewsList, nil
}

func (r *SQLQueryRepository) GetByID(id int) (*views.TodoView, error) {
	row := r.db.QueryRow("SELECT id, title, description, completed FROM todos WHERE id = ?", id)
	var todo views.TodoView
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err == sql.ErrNoRows {
		return nil, errors.New("todo not found")
	}
	return &todo, err
}
