package service

import (
     market "market"
     "market/pkg/repository"
)

type Auth interface {
	Register(user market.User) (int, error)
	GenerateToken(username ,password string) (string, error)
	ParseToken(token string) (int, error)
}


type Service struct {
     Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
          Auth: NewAuthService(repos.Auth),
     }
}

