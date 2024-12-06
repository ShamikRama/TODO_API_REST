package auth

import (
	"TODO_APP/internal/model"
	"TODO_APP/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	errorHashingPassword = "error hashing the password: %w"
	errorCreatingUser    = "error creating user: %w"
)

// здесь было еще поле salt которе использовалось в generatepassword
const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
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
		return "wtf1", fmt.Errorf(errorHashingPassword, err)
	}
	return string(passwordHash), nil
}

func (r *AuthService) GenerateJWTtoken(username, password string) (string, error) {

	pass, err := generatePasswordHash(password)
	if err != nil {
		fmt.Printf("Error generate token")
		return "wtf", err
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

func (r *AuthService) ParseJWTtoken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return signingKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
