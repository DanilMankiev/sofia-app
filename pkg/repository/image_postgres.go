package repository

import (
	"fmt"

	"github.com/DanilMankiev/sofia-app"
	"github.com/jmoiron/sqlx"
)

type ProductImagePostgres struct {
	db *sqlx.DB
}

func NewProductImagePostgres(db *sqlx.DB) *ProductImagePostgres {
	return &ProductImagePostgres{db: db}
}

func (pi *ProductImagePostgres) CreateImage(input sofia.ImageInput) error {
	query := fmt.Sprintf("INSERT INTO %s (product_id,img) values ($1,$2)", imageTable)
	_, err := pi.db.Exec(query, input.Product_id, input.Image)
	if err != nil {
		return err
	}
	return nil
}

func (pi *ProductImagePostgres) GetAllImages(product_id int) ([]string, error) {
	var output []string

	query := fmt.Sprintf("SELECT (img) from %s WHERE product_id=$1", imageTable)

	if err := pi.db.Select(&output, query, product_id); err != nil {
		return output, err // do some
	}
	return output, nil
}

func (pi *ProductImagePostgres) GetImageById(product_id int, image_id int) (string, error) {
	var output string
	query := fmt.Sprintf("SELECT (img) from %s WHERE product_id=$1 AND id = $2", imageTable)

	if err := pi.db.Get(&output, query, product_id, image_id); err != nil {
		return output, err // do some
	}
	return output, nil
}

func (pi *ProductImagePostgres) DeleteImage(image_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", imageTable)
	_, err := pi.db.Exec(query, image_id)
	return err
}
