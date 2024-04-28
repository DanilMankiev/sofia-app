package main

import (
	"context"
	_ "database/sql"
	"os"
	"os/signal"
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
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs:%s", err.Error())
	}


	if err := godotenv.Load("../.env"); err != nil {
		logrus.Fatalf("error loading .env")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:    viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to init DB:%s", err.Error())
	}
	opt := option.WithCredentialsFile(viper.GetString("service-account"))
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
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
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
