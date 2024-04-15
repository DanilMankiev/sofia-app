package service

import (
	"github.com/DanilMankiev/sofia-app"

	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user sofia.User) (int, error)
	GenegateToken(username, password string) (string, error)
}

type ListOfproducts interface {
	CreateList(list entity.List) (int, error)
	GetAllLists() ([]entity.List, error)
	GetListById(id int) (entity.List, error)
	UpdateList(id int, input entity.UpdateListInput) error
	DeleteList(id int) error
}

type Product interface {
	GetAllItems(list_id int) ([]entity.Product, error)
	CreateProduct(list_id int, input entity.CreateProduct) (int, error)
	GetItemByid(product_id int) (entity.Product, error)
	DeleteItem(product_id int) error
	UpdateItem(product_id int, input entity.UpdateProductInput) error
}

type ProductImage interface {
	CreateImage(input entity.ImageInput) error
	GetAllImages(product_id int) ([]string, error)
	GetImageById(product_id int, image_id int) (string, error)
	DeleteImage(image_id int) error
}

type Review interface{
	CreateReview(input entity.CreateReview) (int,error)
	GetAllReview() ([]entity.Review,error)
	DeleteReview(id int) error
	UpdateReview(id int, input entity.UpdateReview) error
}

type Blog interface{
	CreateBlog(input entity.CreateBlog) (int,error)
	GetAllBlog() ([]entity.Blog,error)
	DeleteBlog(id int) error
	UpdateBlog(id int, input entity.UpdateBlog) error
	GetBlogById(id int) (entity.Blog,error)
}

type BlogImage interface{
	CreateImage(input entity.ImageInputBlog) error
	GetAllImages(id int) ([]string, error)
	GetImageById(id int, image_id int) (string, error)
	DeleteImage(image_id int) error
}

type Service struct {
	Authorization
	ListOfproducts
	Product
	ProductImage
	BlogImage
	Review
	Blog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  newAuthService(repos.Authorization),
		ListOfproducts: newListService(repos.ListOfproducts),
		Product:        newProductService(repos.Product, repos.ListOfproducts),
		ProductImage:   newProductImageService(repos.ProductImage),
		Review: newReviewService(repos.Review),
		Blog:  newBlogService(repos.Blog),
		BlogImage: newBlogImageService(repos.BlogImage),
	}
}
