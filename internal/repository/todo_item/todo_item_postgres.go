package todo_item

import (
	"TODO_APP/internal/model"
	"database/sql"
	"fmt"
	"strings"
)

const (
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
	usersListsTable = "users_lists"
)

type TodoItemPostgres struct {
	db *sql.DB
}

func NewTodoItemPostgres(db *sql.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(listID int, item model.TodoItem) (int, error) {
	const op = "sql.TodoItem.Create"

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	createItemQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	createListsItemQuery := fmt.Sprintf(
		"INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)

	_, err = tx.Exec(createListsItemQuery, listID, itemId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return itemId, tx.Commit()

}

func (r *TodoItemPostgres) GetAll(userID int, listID int) ([]model.TodoItem, error) {
	const op = "sql.TodoItem.GetAll"

	var items []model.TodoItem

	getAllItemQuery := fmt.Sprintf(
		"SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND li.list_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)

	rows, err := r.db.Query(getAllItemQuery, userID, listID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.TodoItem
		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userID int, itemID int) (model.TodoItem, error) {
	const op = "sql.TodoItem.GetById"

	var item model.TodoItem

	getItemQuery := fmt.Sprintf(
		"SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND ti.id= $2",
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.QueryRow(getItemQuery, userID, itemID).Scan(&item)
	if err != nil {
		return item, fmt.Errorf("%s : %w", op, err)
	}

	return item, nil
}

func (r *TodoItemPostgres) DeleteById(userID int, itemID int) error {
	const op = "sql.TodoItem.DeleteById"

	getItemQuery := fmt.Sprintf(
		"DELETE FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND ti.id= $2",
		todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(getItemQuery, userID, itemID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (r *TodoItemPostgres) UpdateById(userID int, itemID int, input model.UpdateItemInput) error {
	const op = "sql.TodoItem.DeleteById"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++

	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Description)
		argId++

	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Done)
		argId++

	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)

	args = append(args, userID, itemID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
