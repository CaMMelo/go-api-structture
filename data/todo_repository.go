package data

import (
	"go-api-structure/inputs"
	"go-api-structure/model"
	"go-api-structure/views"
)

type TodoRepository interface {
	Create(todo inputs.CreateTodoInput) (*model.Todo, error)
	Remove(id int) error
	Update(todo model.Todo) error
	GetByID(id int) (*model.Todo, error)
}

type TodoQueryRepository interface {
	GetAll() ([]views.TodoView, error)
	GetByID(id int) (*views.TodoView, error)
}
