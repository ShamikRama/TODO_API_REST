package repository

import (
	"TODO_APP/internal/model"
	"database/sql"
)

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userID int, list model.TodoList) (int, error)
	GetAll(userID int) ([]model.TodoList, error)
	GetById(userID int) (model.TodoList, error)
	Delete(userID int, itemID int) error
	Update(userID int, itemID int) error
}

type TodoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(userID int, listID int) ([]model.TodoItem, error)
	GetById(listID int, itemId int) (model.TodoItem, error)
	Delete(userID int, itemID int) error
	Update(userID int, itemID int, input model.UpdateItemInput) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		//TodoList:      NewTodoListPostgres(db),
		//TodoItem:      NewTodoItemPostgres(db),
	}
}
