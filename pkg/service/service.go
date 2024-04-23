package service

import (
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type Authorization interface {
	SignUp(user entity.SignUpInput) (string, error)
	SignIn(entity.SignInInput) (string, error)
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
	GetAllImages(id int) ([]string, error)
	GetImageById(id int, image_id int) (string, error)
	DeleteImage(image_id int) error
}

type Service struct {
	Authorization
	Category
	Product
	ProductImage
	BlogImage
	Review
	Blog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:      newAuthService(repos.Authorization),
		Category: newCategoryService(repos.Category),
		Product:            newProductService(repos.Product, repos.Category),
		ProductImage:       newProductImageService(repos.ProductImage),
		Review:             newReviewService(repos.Review),
		Blog:               newBlogService(repos.Blog),
		BlogImage:          newBlogImageService(repos.BlogImage),
	}
}
