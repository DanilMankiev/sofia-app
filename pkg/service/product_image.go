package service

import (
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ImageService struct {
	repo repository.ProductImage
}

func newProductImageService(repo repository.ProductImage) *ImageService {
	return &ImageService{repo: repo}
}

func (im *ImageService) CreateImage(input entity.ImageInput) error {
	return im.repo.CreateImage(input)
}

func (im *ImageService) DeleteImage(image_id int, prouct_id int) error {
	return im.repo.DeleteImage(image_id, prouct_id)
}
