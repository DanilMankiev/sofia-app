package service

import (
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ProductService struct {
	repo     repository.Product
	listrepo repository.ListOfproducts
}

func newProductService(repo repository.Product, listrepo repository.ListOfproducts) *ProductService {
	return &ProductService{repo: repo, listrepo: listrepo}
}

func (it *ProductService) GetAllItems(list_id int) ([]entity.Product, error) {

	return it.repo.GetAllItems(list_id)
}

func (it *ProductService) CreateProduct(list_id int, input entity.CreateProduct) (int, error) {
	_, err := it.listrepo.GetListById(list_id)
	if err != nil {
		// error
		return 0, err
	}
	return it.repo.CreateProduct(list_id, input)
}

func (it *ProductService) GetItemByid(product_id int) (entity.Product, error) {
	return it.repo.GetItemByid(product_id)
}

func (it *ProductService) DeleteItem(product_id int) error {
	return it.repo.DeleteItem(product_id)
}

func (it *ProductService) UpdateItem(product_id int, input entity.UpdateProductInput) error {
	return it.repo.UpdateItem(product_id, input)
}
