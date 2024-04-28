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

func (bs *BlogImageService) DeleteImage(image_id int) error {
	return bs.repo.DeleteImage(image_id)
}

func(bs * BlogImageService) CreatePreviewImage(url string,id int) error{
	return bs.repo.CreatePreviewImage(url,id)
}