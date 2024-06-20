package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	market "market"
	"market/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const salt = "dy234ghuwgfwefwe663fegcssgf6"
const signingKey = "hueiw&fhueg3rw9fh43"

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user market.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.Register(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
			"iat": time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
