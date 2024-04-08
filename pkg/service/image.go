package service

import (
	"github.com/DanilMankiev/sofia-app"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ImageService struct {
	repo repository.ProductImage
}

func newProductImageService(repo repository.ProductImage) *ImageService {
	return &ImageService{repo: repo}
}

func (im *ImageService) CreateImage(input sofia.ImageInput) error {
	return im.repo.CreateImage(input)
}

func (im *ImageService) GetAllImages(product_id int) ([]string, error) {
	return im.repo.GetAllImages(product_id)
}

func (im *ImageService) GetImageById(product_id int, image_id int) (string, error) {
	return im.repo.GetImageById(product_id, image_id)
}

func (im *ImageService) DeleteImage(image_id int) error {
	return im.repo.DeleteImage(image_id)
}
