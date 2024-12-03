package service

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"TODO_APP/internal/service/auth"
	// list "TODO_APP/internal/service/todo_list"
	// item "TODO_APP/internal/service/todo_item"
)

type Authorization interface {
	Create(user model.User) (int, error)
	GenerateJWTtoken(username, password string) (string, error)
	ParseJWTtoken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userID int, list model.TodoList) (int, error)
	GetAllLists(userID int) ([]model.TodoList, error)
	DeleteList(userID int, listID int) error
	GetList(userID int, listID int) (model.TodoList, error)
	UpdateList(userID int, listID int, input model.UpdateListInput) error
}

type TodoItem interface {
	// ...
}

type Service struct {
	Authorization
	TodoList
	// TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthService(repo.Authorization),
		//TodoList:      list.NewTodoListService(repo.TodoList),
		//	TodoItem:      todo_item.NewTodoItemService(),
	}
}
