package infra

import (
	"errors"
	"go-api-structure/inputs"
	"go-api-structure/model"
)

type InMemoryTodoRepository struct {
	todos *map[int]model.Todo
}

func NewInMemoryTodoRepository(todos *map[int]model.Todo) *InMemoryTodoRepository {
	return &InMemoryTodoRepository{todos: todos}
}

func (r *InMemoryTodoRepository) Create(input inputs.CreateTodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		ID:          len(*r.todos) + 1,
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}
	(*r.todos)[todo.ID] = *todo
	return todo, nil
}

func (r *InMemoryTodoRepository) Remove(id int) error {
	if _, exists := (*r.todos)[id]; !exists {
		return errors.New("todo not found")
	}
	delete(*r.todos, id)
	return nil
}

func (r *InMemoryTodoRepository) Update(todo model.Todo) error {
	(*r.todos)[todo.ID] = todo
	return nil
}

func (r *InMemoryTodoRepository) GetByID(id int) (*model.Todo, error) {
	todo, exists := (*r.todos)[id]
	if !exists {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}
