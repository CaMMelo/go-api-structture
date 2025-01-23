package infra

import (
	"errors"
	"go-api-structure/model"
	"go-api-structure/views"
)

type QueryRepository struct {
	todos *map[int]model.Todo
}

func NewQueryRepository(todos *map[int]model.Todo) *QueryRepository {
	return &QueryRepository{todos: todos}
}

func (r *QueryRepository) GetAll() ([]views.TodoView, error) {
	viewsList := []views.TodoView{}
	for _, todo := range *r.todos {
		viewsList = append(viewsList, views.NewTodoView(todo.ID, todo.Title, todo.Description, todo.Completed))
	}
	return viewsList, nil
}

func (r *QueryRepository) GetByID(id int) (*views.TodoView, error) {
	todo, exists := (*r.todos)[id]
	if !exists {
		return nil, errors.New("todo not found")
	}
	view := views.NewTodoView(todo.ID, todo.Title, todo.Description, todo.Completed)
	return &view, nil
}
