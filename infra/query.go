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
		viewsList = append(viewsList, views.TodoView{ID: todo.ID, Title: todo.Title, Description: todo.Description, Completed: todo.Completed})
	}
	return viewsList, nil
}

func (r *QueryRepository) GetByID(id int) (*views.TodoView, error) {
	todo, exists := (*r.todos)[id]
	if !exists {
		return nil, errors.New("todo not found")
	}
	view := views.TodoView{ID: todo.ID, Title: todo.Title, Description: todo.Description, Completed: todo.Completed}
	return &view, nil
}
