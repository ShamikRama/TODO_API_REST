package repository

import (
	"TODO_APP/internal/model"
	auth "TODO_APP/internal/repository/auth"
	todolist "TODO_APP/internal/repository/todo_list"
	"database/sql"
	// todoitem "TODO_APP/internal/repository/todo_item"
)

type Repository struct {
	Authorization
	TodoList
	// TodoItem
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userID int, list model.TodoList) (int, error)
	//GetAll(userID int) ([]model.TodoList, error)
	//GetById(userID int) (model.TodoList, error)
	//Delete(userID int, itemID int) error
	//Update(userID int, itemID int) error
}

// type TodoItem interface {
// 	Create(listId int, item model.TodoItem) (int, error)
// 	GetAll(userID int, listID int) ([]model.TodoItem, error)
// 	GetById(listID int, itemId int) (model.TodoItem, error)
// 	Delete(userID int, itemID int) error
// 	Update(userID int, itemID int, input model.UpdateItemInput) error
// }

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(db),
		TodoList:      todolist.NewTodoListPostgres(db),
		//TodoItem:      todoitem.NewTodoItemPostgres(db),
	}
}
