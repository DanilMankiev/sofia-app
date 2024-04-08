package service

import (
	"github.com/DanilMankiev/sofia-app"

	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user sofia.User) (int, error)
	GenegateToken(username, password string) (string, error)
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
	GetItemByid(product_id int) (sofia.Product, error)
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

type Service struct {
	Authorization
	ListOfproducts
	Product
	ProductImage
	ListImage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  newAuthService(repos.Authorization),
		ListOfproducts: newListService(repos.ListOfproducts),
		Product:        newProductService(repos.Product, repos.ListOfproducts),
		ProductImage:   newProductImageService(repos.ProductImage),
	}
}
