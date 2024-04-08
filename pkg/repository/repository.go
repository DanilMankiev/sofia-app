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

type ProductImage interface {
	CreateImage(input sofia.ImageInput) error
	GetAllImages(product_id int) ([]string, error)
	GetImageById(product_id int, image_id int) (string, error)
	DeleteImage(image_id int) error
}

type ListImage interface {
}

type Repository struct {
	Authorization
	ListOfproducts
	Product
	ProductImage
	ListImage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		ListOfproducts: NewListPostgres(db),
		Product:        NewProductPostgres(db),
		ProductImage:   NewProductImagePostgres(db),
	}
}
