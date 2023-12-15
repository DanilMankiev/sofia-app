package main

import (
	"log"

	"github.com/DanilMankiev/todo-app"
	"github.com/DanilMankiev/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while runnt=ing http server: %s", err.Error())
	}
}
