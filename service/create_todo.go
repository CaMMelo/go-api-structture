package service

import (
	"go-api-structure/data"
	"go-api-structure/inputs"
	"go-api-structure/model"
)

type CreateTodoService struct {
	Repository data.TodoRepository
}

func NewCreateTodoService(repo data.TodoRepository) *CreateTodoService {
	return &CreateTodoService{Repository: repo}
}

func (s *CreateTodoService) Execute(input inputs.CreateTodoInput) (*model.Todo, error) {
	return s.Repository.Create(input)
}
