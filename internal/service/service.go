package service

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"TODO_APP/internal/service/auth"
	// "TODO_APP/internal/service/todo_item"
	// "TODO_APP/internal/service/todo_list"
)

type Authorization interface {
	Create(user model.User) (int, error)
	GenerateJWTtoken(username, password string) (string, error)
	// ParseJWTtoken()
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
		Authorization: auth.NewAuthService(repo),
		//	TodoItem:      todo_item.NewTodoItemService(),
		//	TodoList:      todo_list.NewTodoListService(),
	}
}
