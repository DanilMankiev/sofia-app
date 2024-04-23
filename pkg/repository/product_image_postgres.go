package repository

import (
	"fmt"

	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type ProductImagePostgres struct {
	db *sqlx.DB
}

func NewProductImagePostgres(db *sqlx.DB) *ProductImagePostgres {
	return &ProductImagePostgres{db: db}
}

func (pi *ProductImagePostgres) CreateImage(input entity.ImageInput) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = ARRAY_APPEND(image_all, $1) WHERE id=$2", productTable)
	_, err := pi.db.Exec(query,input.Image,input.Product_id)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("UPDATE %s SET image_preview = (SELECT image_all[1] FROM %s WHERE image_preview IS NULL)", productTable,productTable)
	_,err= pi.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (pi *ProductImagePostgres) GetAllImages(product_id int) ([]string, error) {
	var output []string

	query := fmt.Sprintf("SELECT (img) from %s WHERE product_id=$1", productImagesTable)

	if err := pi.db.Select(&output, query, product_id); err != nil {
		return output, err // TODO
	}
	return output, nil
}

func (pi *ProductImagePostgres) GetImageById(product_id int, image_id int) (string, error) {
	var output string
	query := fmt.Sprintf("SELECT (img) from %s WHERE product_id=$1 AND id = $2", productImagesTable)

	if err := pi.db.Get(&output, query, product_id, image_id); err != nil {
		return output, err // TODO
	}
	return output, nil
}

func (pi *ProductImagePostgres) DeleteImage(image_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", productImagesTable)
	_, err := pi.db.Exec(query, image_id)
	return err
}
