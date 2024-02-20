package main

import (
	_ "database/sql"
	"os"
	"path/filepath"

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
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs:%s", err.Error())
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	err = godotenv.Load(filepath.Join(dir, ".env"))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "root",
	})
	if err != nil {
		logrus.Fatalf("Failed to init DB:%s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandeler(services)

	srv := new(sofia.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error ocured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
