package repository

import (
	"firebase.google.com/go/v4/auth"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	SignUp(user entity.SignUpInput) (string, error)
	SignIn(input entity.SignInInput) (string, error)
}

type Category interface {
	CreateCategory(category entity.Category) (int, error)
	GetAllCategorys() ([]entity.Category, error)
	GetCategoryById(id int) (entity.Category, error)
	UpdateCategory(id int, input entity.UpdateCategoryInput) error
	DeleteCategory(id int) error
}

type Product interface {
	GetAllItems(category_id int) ([]entity.Product, error)
	CreateProduct(category_id int, input entity.CreateProduct) (int, error)
	GetItemByid(poduct_id int) (entity.Product, error)
	DeleteItem(product_id int) error
	UpdateItem(product_id int, input entity.UpdateProductInput) error
}

type ProductImage interface {
	CreateImage(input entity.ImageInput) error
	DeleteImage(image_id int, prouct_id int) error
}

type Review interface {
	CreateReview(input entity.CreateReview) (int, error)
	GetAllReview() ([]entity.Review, error)
	DeleteReview(id int) error
	UpdateReview(id int, input entity.UpdateReview) error
}

type Blog interface {
	CreateBlog(input entity.CreateBlog) (int, error)
	GetAllBlog() ([]entity.Blog, error)
	DeleteBlog(id int) error
	UpdateBlog(id int, input entity.UpdateBlog) error
	GetBlogById(id int) (entity.Blog, error)
}

type BlogImage interface {
	CreateImage(input entity.ImageInputBlog) error
	GetAllImages(blog_id int) ([]string, error)
	GetImageById(blog_id int, image_id int) (string, error)
	DeleteImage(image_id int) error
}

type Repository struct {
	Authorization
	Category
	Product
	ProductImage
	Review
	Blog
	BlogImage
}

func NewRepository(db *sqlx.DB, FireAuth *auth.Client) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db, FireAuth),
		Category:      NewCategoryPostgres(db),
		Product:       NewProductPostgres(db),
		ProductImage:  NewProductImagePostgres(db),
		Review:        NewReviewPostgres(db),
		Blog:          NewBlogPostgres(db),
		BlogImage:     NewBlogImagePostgres(db),
	}
}
