package repository

import (
	market "market"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	Register(user market.User) (int, error)
	GetUser(username, password string) (market.User, error)
}

type Repository struct {
     Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
          Auth: NewAuthPostgres(db),
     }
}