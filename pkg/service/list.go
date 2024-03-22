package service

import (
	"github.com/DanilMankiev/sofia-app"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ListService struct {
	repo repository.ListOfproducts
}

func newListService(repo repository.ListOfproducts) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) CreateList(list sofia.List) (int, error) {
	return s.repo.CreateList(list)
}

func (s *ListService) GetAllLists() ([]sofia.List, error) {
	return s.repo.GetAllLists()
}

func (s *ListService) GetListById(id int) (sofia.List, error) {
	return s.repo.GetListById(id)
}

func (s *ListService) UpdateList(id int, input sofia.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(id, input)
}

func (s *ListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}
