package entity

import "errors"

type Product struct {
	Product_id   int    `json:"product_id" db:"product_id"`
	Product_name string `json:"product_name" binding:"required"`
	List_id      int    `json:"list_id"`
	Description  string `json:"description" binding:"required"`
	Price        int    `json:"price" binding:"required"`
}

type CreateProduct struct {
	Product_name string `json:"product_name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Price        int    `json:"price" binding:"required"`
}

type UpdateProductInput struct {
	Product_name *string `json:"product_name"`
	List_id      *int    `json:"list_id"`
	Description  *string `json:"description"`
	Price        *int    `json:"price"`
}

func (up UpdateProductInput) Validate() error {
	if up.Product_name == nil && up.List_id == nil && up.Description == nil && up.Price == nil {
		return errors.New("update item no valiable")
	}
	return nil
}
