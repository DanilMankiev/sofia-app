package entity

import (
	"errors"

	"github.com/lib/pq"
)

type Product struct {
	Id                 int            `json:"id" db:"id"`
	Name               string         `json:"name" binding:"required"`
	CategoryId         string         `json:"category_id" db:"category_id"`
	CategoryName       string         `json:"category_name" db:"category"`
	DescriptionPreview string         `json:"description_preview" binding:"required" db:"description_preview"`
	FullDescription    string         `json:"description_full" binding:"required" db:"description_full"`
	ImgaePreview       string         `json:"image_preview" db:"image_preview"`
	AllImages          pq.StringArray `json:"image_all" db:"image_all"`
	Composition        string         `json:"composition" binding:"required" db:"composition"`
	TermsPurchase      string         `json:"purchase" binding:"required" db:"purchase"`
	Delivery           string         `json:"delivery" binding:"required" db:"delivery"`
	Price              int            `json:"price" binding:"required" db:"price"`
	Furniture          string         `json:"furniture" binding:"required" db:"furniture"`
}

type CreateProduct struct {
	Name               string         `json:"name" binding:"required"`
	DescriptionPreview string         `json:"description_preview" binding:"required"`
	FullDescription    string         `json:"description_full" binding:"required"`
	ImagePreview       string         `json:"image_preview"`
	AllImages          pq.StringArray `json:"image_all"`
	Composition        string         `json:"composition" binding:"required"`
	TermsPurchase      string         `json:"purchase" binding:"required"`
	Delivery           string         `json:"delivery" binding:"required"`
	Price              int            `json:"price" binding:"required"`
	Furniture          string         `json:"furniture" binding:"required" db:"furniture"`
}

type UpdateProductInput struct {
	Name               *string         `json:"name"`
	CategoryId         *string         `json:"category_id"`
	CategoryName       *string         `json:"category_name"`
	DescriptionPreview *string         `json:"description_preview"`
	FullDescription    *string         `json:"description_full"`
	ImgaePreview       *string         `json:"iamge_preview"`
	AllImages          *pq.StringArray `json:"iamge_all"`
	Composition        *string         `json:"composition"`
	TermsPurchase      *string         `json:"purchase"`
	Delivery           *string         `json:"delivery"`
	Price              *int            `json:"price"`
	Furniture          *string          `json:"furniture"`
}

func (up UpdateProductInput) Validate() error {
	if up.Name == nil {
		return errors.New("update item no valiable") // TODO
	}
	return nil
}
