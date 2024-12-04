package service

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"TODO_APP/internal/service/auth"
	item "TODO_APP/internal/service/todo_item"
	list "TODO_APP/internal/service/todo_list"
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
	Create(userID int, listID int, item model.TodoItem) (int, error)
	GetAllItems(userID int, listID int) ([]model.TodoItem, error)
	GetItem(userID int, listID int) (model.TodoItem, error)
	Update(userID int, itemID int, input model.UpdateItemInput) error
	DeleteItem(userID int, itemID int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthService(repo.Authorization),
		TodoList:      list.NewTodoListService(repo.TodoList),
		TodoItem:      item.NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
