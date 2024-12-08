package todo_list

import (
	"TODO_APP/internal/model"
	"database/sql"
	"fmt"
	"strings"
)

const (
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
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
	const op = "sql.TodoList.Create"

	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}

	createListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES($1, $2) RETURNING id", todoListsTable)
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

	createUsersListQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, list_id) VALUES($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userID, listId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return listId, nil
}

func (r *TodoListPostgres) GetAll(userID int) ([]model.TodoList, error) {
	const op = "sql.TodoList.GetAllList"

	var lists []model.TodoList

	query := fmt.Sprintf(
		"SELECT todo_lists.id, todo_lists.title, todo_lists.description FROM %s INNER JOIN %s users_lists ON todo_lists.id = users_lists.list_id WHERE users_lists.user_id = $1",
		todoListsTable, usersListsTable)

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var list model.TodoList
		err := rows.Scan(&list.Id, &list.Title, &list.Description)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		lists = append(lists, list)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return lists, nil
}

func (r *TodoListPostgres) Delete(userID int, listID int) error {
	const op = "sql.TodoList.DeleteList"

	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}

	// Удаляем запись из таблицы users_lists
	deleteUsersListQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE user_id = $1 AND list_id = $2", usersListsTable)
	_, err = tx.Exec(deleteUsersListQuery, userID, listID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	// Удаляем запись из таблицы todo_lists
	deleteListQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1", todoListsTable)
	_, err = tx.Exec(deleteListQuery, listID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return nil
}

func (r *TodoListPostgres) GetById(userID int, listID int) (model.TodoList, error) {
	const op = "sql.TodoList.GettingList"

	var list model.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2", todoListsTable, usersListsTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return list, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID, listID).Scan(&list.Id, &list.Title, &list.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return list, fmt.Errorf("%s: list not found: %w", op, err)
		}
		return list, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return list, nil
}

func (r *TodoListPostgres) UpdateById(userID int, listID int, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listID, userID)

	_, err := r.db.Exec(query, args...)
	return err
}

// Предположим, input.Title = "New Title" и input.Description = "New Description", listID = 101, userID = 202.

// Тогда:

// setValues = []string{"title=$1", "description=$2"}

// args = []interface{}{"New Title", "New Description", 101, 202}

// Запрос: UPDATE todo_lists tl SET title=$1, description=$2 FROM users_lists ul WHERE tl.id = ul.list_id AND ul.list_id=$3 AND ul.user_id=$4

// Выполнение с аргументами: "New Title", "New Description", 101, 202

// База данных заменяет $1 на "New Title", $2 на "New Description" и т.д.
