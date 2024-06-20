package main

import (
	market "market"
	"market/pkg/handler"
	"market/pkg/repository"
	"market/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("configerr", err.Error())
	}
	if err := godotenv.Load("/home/r_rmarsu/new-project/.env"); err != nil {
		logrus.Fatal("enverr", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("dberr died: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	hub := handler.NewHub()
	handlers := handler.NewHandler(services, hub)

	srv := new(market.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("runerr", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("/home/r_rmarsu/new-project/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
