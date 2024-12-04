package todolist

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"fmt"
)

const (
	errorCreatingList   = "error creating user: %w"
	errorGettingAllList = "error getting all list: %w"
	errorDeletingList   = "error deleting the list: %w"
	errorGettingList    = "error getting the list: %w"
	errorUpdatingList   = "error updating list: %w"
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

func (s *TodoListService) GetAllLists(userID int) ([]model.TodoList, error) {
	lists, err := s.repo.GetAll(userID)
	if err != nil {
		return nil, fmt.Errorf(errorGettingAllList, err)
	}

	return lists, nil
}

func (s *TodoListService) DeleteList(userID int, listID int) error {
	err := s.repo.Delete(userID, listID)
	if err != nil {
		return fmt.Errorf(errorDeletingList, err)
	}

	return nil
}

func (s *TodoListService) GetList(userID int, listID int) (model.TodoList, error) {
	list, err := s.repo.GetById(userID, listID)
	if err != nil {
		return list, fmt.Errorf(errorGettingList, err)
	}

	return list, nil
}

func (s *TodoListService) UpdateList(userID int, listID int, input model.UpdateListInput) error {
	err := input.Validate()
	if err != nil {
		return fmt.Errorf("service sloy error updating list")
	}

	err = s.repo.UpdateById(userID, listID, input)
	if err != nil {
		return fmt.Errorf(errorUpdatingList, err)
	}

	return nil
}
