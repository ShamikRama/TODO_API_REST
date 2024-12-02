package todolist

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"fmt"
)

const (
	errorCreatingList = "error creating user: %w"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userID int, list model.TodoList) (int, error) {
	id, err := s.repo.Create(userID, list)
	if err != nil {
		return 0, fmt.Errorf(errorCreatingList, err)
	}
	return id, nil
}
