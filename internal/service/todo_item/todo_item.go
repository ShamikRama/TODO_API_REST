package todo_item

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"fmt"
)

type TodoItemService struct {
	repo     repository.TodoItem
	repolist repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, repolist repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		repolist: repolist,
	}
}

func (s *TodoItemService) Create(userID int, listID int, item model.TodoItem) (int, error) {
	_, err := s.repolist.GetById(userID, listID)
	if err != nil {
		return 0, fmt.Errorf("list does not exists or does not belongs to user")
	}

	return s.repo.Create(listID, item)
}

func (s *TodoItemService) GetAllItems(userID int, listID int) ([]model.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}
