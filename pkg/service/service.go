package service

import (
	"github.com/DanilMankiev/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user sofia.User) (int, error)
	GenegateToken(username, password string) (string, error)
}

type ListOfproducts interface {
	CreateList(list sofia.List) (int, error)
}

type Service struct {
	Authorization
	ListOfproducts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
