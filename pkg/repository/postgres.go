package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	usersTable = "users"
	UsersList  = "users_lists"
	trenTable  = "trening"
)

type Config struct {
	Port     string
	Host     string
	Password string
	UserName string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sqlx.DB, error) {
	Connsrt := fmt.Sprintf("port=%s host=%s password=%s user=%s dbname=%s sslmode=%s",
		cfg.Port, cfg.Host, cfg.Password, cfg.UserName, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Open("postgres", Connsrt)

	if err != nil {
		logrus.Error("cannot set config for db", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
