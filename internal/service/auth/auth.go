package auth

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

const (
	errorHashingPassword = "error hashing the password: %w"
	errorCreatingUser    = "error creating user: %w"
)

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (r *AuthService) Create(user model.User) (int, error) {
	user.Password, _ = generatePasswordHash(user.Password)

	id, err := r.repo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf(errorCreatingUser, err)
	}
	return id, nil
}

func generatePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf(errorHashingPassword, err)
	}
	return string(passwordHash), nil
}
