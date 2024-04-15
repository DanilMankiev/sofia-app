package main

import (
	"context"
	_ "database/sql"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/sirupsen/logrus"

	_ "github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/DanilMankiev/sofia-app"
	"github.com/DanilMankiev/sofia-app/pkg/handler"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
	"github.com/DanilMankiev/sofia-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	// if err := initConfig(); err != nil {
	// 	logrus.Fatalf("Error initializing configs:%s", err.Error())
	// }

	_, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	if err = godotenv.Load("../.env"); err != nil {
		logrus.Fatalf("error loading .env")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		Username: os.Getenv("username"),
		DBname:   os.Getenv("dbname"),
		SSLMode:  os.Getenv("sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to init DB:%s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandeler(services)

	srv := new(sofia.Server)

	go func() {
		if err := srv.Run(os.Getenv("ports"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
