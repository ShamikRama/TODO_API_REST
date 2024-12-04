package repository

import (
	"TODO_APP/internal/model"
	auth "TODO_APP/internal/repository/auth"
	item "TODO_APP/internal/repository/todo_item"
	list "TODO_APP/internal/repository/todo_list"
	"database/sql"
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userID int, list model.TodoList) (int, error)
	GetAll(userID int) ([]model.TodoList, error)
	GetById(userID int, listID int) (model.TodoList, error)
	Delete(userID int, listID int) error
	UpdateById(userID int, listID int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(userID int, listID int) ([]model.TodoItem, error)
	GetById(listID int, itemId int) (model.TodoItem, error)
	DeleteById(userID int, itemID int) error
	UpdateById(userID int, itemID int, input model.UpdateItemInput) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(db),
		TodoList:      list.NewTodoListPostgres(db),
		TodoItem:      item.NewTodoItemPostgres(db),
	}
}
