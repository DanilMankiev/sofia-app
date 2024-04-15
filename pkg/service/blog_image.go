package service

import (
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type BlogImageService struct {
	repo repository.BlogImage
}

func newBlogImageService(repo repository.BlogImage) *BlogImageService {
	return &BlogImageService{repo: repo}
}

func (bs *BlogImageService) CreateImage(input entity.ImageInputBlog) error {
	return bs.repo.CreateImage(input)
}

func (bs *BlogImageService) GetAllImages(product_id int) ([]string, error) {
	return bs.repo.GetAllImages(product_id)
}

func (bs *BlogImageService) GetImageById(product_id int, image_id int) (string, error) {
	return bs.repo.GetImageById(product_id, image_id)
}

func (bs *BlogImageService) DeleteImage(image_id int) error {
	return bs.repo.DeleteImage(image_id)
}
