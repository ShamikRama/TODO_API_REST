package repository

import (
	"TODO_APP/internal/model"
	"database/sql"
	"fmt"
)

const usersTable = "users"

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	const op = "sql.Auth.CreateUser"

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES($1, $2, $3) RETURNING id", usersTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var id int

	err = stmt.QueryRow(user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	const op = "sql.Auth.GetUser"

	var user model.User
	query := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2", usersTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, password).Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return user, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return user, nil
}
