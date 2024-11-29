package service

import "TODO_APP/internal/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		// Authorization: NewAuthService(repo.Authorization),
		//	TodoItem:      NewTodoItemService(),
		//	TodoList:      NewTodoListService(),
	}
}
