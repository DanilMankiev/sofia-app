package service

import (
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ListService struct {
	repo repository.ListOfproducts
}

func newListService(repo repository.ListOfproducts) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) CreateList(list entity.List) (int, error) {
	return s.repo.CreateList(list)
}

func (s *ListService) GetAllLists() ([]entity.List, error) {
	return s.repo.GetAllLists()
}

func (s *ListService) GetListById(id int) (entity.List, error) {
	return s.repo.GetListById(id)
}

func (s *ListService) UpdateList(id int, input entity.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(id, input)
}

func (s *ListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}
