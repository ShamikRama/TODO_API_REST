package service

import "TODO_APP/internal/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
