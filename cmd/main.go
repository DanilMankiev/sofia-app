package main

import (
	"context"
	_ "database/sql"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	firebase "firebase.google.com/go/v4"
	"log"
	"github.com/sirupsen/logrus"

	_ "github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/DanilMankiev/sofia-app"
	"github.com/DanilMankiev/sofia-app/pkg/handler"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
	"github.com/DanilMankiev/sofia-app/pkg/service"
	"github.com/spf13/viper"


	"google.golang.org/api/option"
	
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
	opt := option.WithCredentialsFile("../service-account-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Create a Firebase auth client instance
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase auth client: %v", err)
	}

	repos := repository.NewRepository(db, authClient)
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

	_,err=firebase.NewApp(context.Background(),nil)
	if err!=nil{
		log.Fatalf("error initalizing app: %v\n", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
