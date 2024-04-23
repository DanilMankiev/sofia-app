package service

import (
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ProductService struct {
	repo         repository.Product
	categoryrepo repository.Category
}

func newProductService(repo repository.Product, categoryrepo repository.Category) *ProductService {
	return &ProductService{repo: repo, categoryrepo: categoryrepo}
}

func (it *ProductService) GetAllItems(category_id int) ([]entity.Product, error) {

	return it.repo.GetAllItems(category_id)
}

func (it *ProductService) CreateProduct(category_id int, input entity.CreateProduct) (int, error) {
	_, err := it.categoryrepo.GetCategoryById(category_id)
	if err != nil {
		// error
		return 0, err
	}
	return it.repo.CreateProduct(category_id, input)
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
