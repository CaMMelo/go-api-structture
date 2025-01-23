package service

import (
	"go-api-structure/data"
)

type ToggleTodoService struct {
	Repository data.TodoRepository
}

func NewToggleTodoService(repo data.TodoRepository) *ToggleTodoService {
	return &ToggleTodoService{Repository: repo}
}

func (s *ToggleTodoService) Execute(id int) error {
	todo, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}
	todo.Toggle()
	s.Repository.Update(*todo)
	return nil
}
