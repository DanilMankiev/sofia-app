package service

import (
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func newCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category entity.Category) (int, error) {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetAllCategorys() ([]entity.Category, error) {
	return s.repo.GetAllCategorys()
}

func (s *CategoryService) GetCategoryById(id int) (entity.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s *CategoryService) UpdateCategory(id int, input entity.UpdateCategoryInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateCategory(id, input)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}
