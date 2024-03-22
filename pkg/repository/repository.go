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
	GetAllLists() ([]sofia.List, error)
	GetListById(id int) (sofia.List, error)
	UpdateList(id int, input sofia.UpdateListInput) error
	DeleteList(id int) error
}

type Product interface {
	GetAllItems(list_id int) ([]sofia.Product, error)
	CreateProduct(list_id int, input sofia.CreateProduct) (int, error)
	GetItemByid(poduct_id int) (sofia.Product, error)
	DeleteItem(product_id int) error
	UpdateItem(product_id int, input sofia.UpdateItemInput) error
}

type Repository struct {
	Authorization
	ListOfproducts
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		ListOfproducts: NewListPostgres(db),
		Product:        NewProductPostgres(db),
	}
}
