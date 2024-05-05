package service

import (
	

	"firebase.google.com/go/v4/auth"
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type Authorization interface {
	SignUp(user entity.SignUpInput) error
	SignIn(entity.SignInInput) (string, error)
	ParseToken(token string) (string,error)

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
	CreateImage(input entity.ImageInputProduct) error
	DeleteImage(image_id int, prouct_id int) error
	CreatePreviewImage(url string,id int) error
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
	DeleteImage(image_id int) error
	CreatePreviewImage(url string,id int) error
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

func NewService(repos *repository.Repository,FireAuth *auth.Client) *Service {
	return &Service{
		Authorization:      newAuthService(repos.Authorization, FireAuth),
		Category: newCategoryService(repos.Category),
		Product:            newProductService(repos.Product, repos.Category),
		ProductImage:       newProductImageService(repos.ProductImage),
		Review:             newReviewService(repos.Review),
		Blog:               newBlogService(repos.Blog),
		BlogImage:          newBlogImageService(repos.BlogImage),
	}
}
