package auth

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	errorHashingPassword = "error hashing the password: %w"
	errorCreatingUser    = "error creating user: %w"
)

// здесь было еше поле salt которе использовалось в generatepassword
const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (r *AuthService) GenerateJWTtoken(username, password string) (string, error) {
	pass, err := generatePasswordHash(password)
	if err != nil {
		fmt.Printf("Error generate token")
		return "", err
	}

	user, err := r.repo.GetUser(username, pass)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))

}

func (r *AuthService) ParseJWTtoken(token string) (int, error) {
	return 0, nil
}
