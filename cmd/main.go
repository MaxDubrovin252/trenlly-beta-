package main

import (
	"log/slog"
	"os"
	"trenlly/config"
	"trenlly/internal/handler"
	"trenlly/pkg/repository"
	"trenlly/pkg/service"

	"trenlly/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := config.InitConfig(); err != nil {
		logrus.Fatal(err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Errorf("enviromental values not set:%s", err.Error())
		os.Exit(1)
	}

	db, err := repository.NewDB(repository.Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		UserName: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Errorf("cannot connect to db :%s", err.Error())
	}

	logrus.Info("successfuly connect to db ", slog.String("PORT", viper.GetString("db.port")))

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	srv := new(server.SRV)

	if err := srv.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Errorf("cannot run server:%s", err.Error())
	}

}
