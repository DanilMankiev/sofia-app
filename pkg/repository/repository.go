package repository

import (
	"github.com/DanilMankiev/sofia-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user sofia.User) (int, error)
	GetUser(username, password string) (sofia.User, error)
}

type ListOfproducts interface {
	CreateList(list sofia.List) (int, error)
}

type Item interface {
}

type Repository struct {
	Authorization
	ListOfproducts
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		ListOfproducts: NewListPostgres(db),
	}
}
