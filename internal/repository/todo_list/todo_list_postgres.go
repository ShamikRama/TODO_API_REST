package todo_list

import (
	"TODO_APP/internal/model"
	"database/sql"
	"fmt"
)

const (
	todoListsTable  = "todo_lists"
	usersListsTable = "users_list"
)

type TodoListPostgres struct {
	db *sql.DB
}

func NewTodoListPostgres(db *sql.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (r *TodoListPostgres) Create(userID int, list model.TodoList) (int, error) {
	const op = "sql.Create.TodoList"

	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}

	// Создание списка
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES($1, $2) RETURNING id", todoListsTable)
	stmt, err := tx.Prepare(createListQuery)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var listId int
	err = stmt.QueryRow(list.Title, list.Description).Scan(&listId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	// Связывание списка с пользователем
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userID, listId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	// Завершение транзакции
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return listId, nil
}
