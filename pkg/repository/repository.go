package repository

import (
	"github.com/DanilMankiev/sofia-app"
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user sofia.User) (int, error)
	GetUser(username, password string) (sofia.User, error)
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
	GetItemByid(poduct_id int) (entity.Product, error)
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

type BlogImage interface {
	CreateImage(input entity.ImageInputBlog) error
	GetAllImages(blog_id int) ([]string, error)
	GetImageById(blog_id int, image_id int) (string, error)
	DeleteImage(image_id int) error
}

type Repository struct {
	Authorization
	ListOfproducts
	Product
	ProductImage
	Review
	Blog
	BlogImage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		ListOfproducts: NewListPostgres(db),
		Product:        NewProductPostgres(db),
		ProductImage:   NewProductImagePostgres(db),
		Review: NewReviewPostgres(db),
		Blog: NewBlogPostgres(db),
		BlogImage: NewBlogImagePostgres(db),
	}
}
