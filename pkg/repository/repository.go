package repository

import (
	"trenlly"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user trenlly.User) (int, error)
	GetUser(username, password string) (trenlly.User, error)
}

type Trening interface {
	CreateTren(tren trenlly.Trening, userId int) (int, error)
	GetAll(userId int) ([]trenlly.Trening, error)
	GetById(userId, trenId int) (trenlly.Trening, error)
	UpdateTren(userId, trenId int, input trenlly.UpdateTrenItem) error
	Delete(userId, trenId int) error
}

type Repository struct {
	Authorization
	Trening
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Trening:       NewTrenRepos(db),
	}
}
