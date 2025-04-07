package service

import (
	"trenlly"
	"trenlly/pkg/repository"
)

type Authorization interface {
	CreateUser(user trenlly.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Trening interface {
	CreateTren(tren trenlly.Trening, userId int) (int, *trenlly.Result, error)
	GetAll(userId int) ([]trenlly.Trening, error)
	GetById(userId, trenId int) (trenlly.Trening, error)
	UpdateTren(userId, trenId int, input trenlly.UpdateTrenItem) error
	Delete(userId, trenId int) error
}

type Service struct {
	Authorization
	Trening
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Trening:       NewTreningService(repos.Trening),
	}
}
