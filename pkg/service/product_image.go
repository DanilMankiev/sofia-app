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

func (im *ImageService) CreateImage(input entity.ImageInputProduct) error {
	return im.repo.CreateImage(input)
}

func (im *ImageService) DeleteImage(product_id int) error {
	return im.repo.DeleteImage(product_id)
}

func(im *ImageService) CreatePreviewImage(url string, id int) error{
	return im.repo.CreatePreviewImage(url, id)
}