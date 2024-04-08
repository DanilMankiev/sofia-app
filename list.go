package sofia

import (
	"errors"
)

type List struct {
	List_id     int    `json:"list_id" db:"list_id"`
	Listname    string `json:"listname" binding:"required"`
	Description string `json:"description" binding:"required"`
}

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

type UpdateListInput struct {
	Listname    *string `json:"listname"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Product_name *string `json:"product_name"`
	List_id      *int    `json:"list_id"`
	Description  *string `json:"description"`
	Price        *int    `json:"price"`
}

type ImageInput struct {
	Product_id int    `json:"product_id"`
	Image      string `json:"img"`
}

type ImageOutput struct {
	Url string `json:"url"`
}

func (up UpdateListInput) Validate() error {
	if up.Listname == nil && up.Description == nil {
		return errors.New("update table no validate")
	}
	return nil
}

func (up UpdateItemInput) Validate() error {
	if up.Product_name == nil && up.List_id == nil && up.Description == nil && up.Price == nil {
		return errors.New("update item no valiable")
	}
	return nil
}
