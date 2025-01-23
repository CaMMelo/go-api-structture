package service

import (
	"go-api-structure/data"
)

type RemoveTodoService struct {
	Repository data.TodoRepository
}

func NewRemoveTodoService(repo data.TodoRepository) *RemoveTodoService {
	return &RemoveTodoService{Repository: repo}
}

func (s *RemoveTodoService) Execute(id int) error {
	return s.Repository.Remove(id)
}
