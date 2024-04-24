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
	query = fmt.Sprintf("UPDATE %s SET image_preview = image_all[1] WHERE id=$1", productTable)
	_,err= pi.db.Exec(query,input.Product_id)
	if err != nil {
		return err
	}
	return nil
}

// func (pi *ProductImagePostgres) GetAllImages(product_id int) ([]string, error) {
// 	var output []string

// 	query := fmt.Sprintf("SELECT (img) from %s WHERE product_id=$1", productImagesTable)

// 	if err := pi.db.Select(&output, query, product_id); err != nil {
// 		return output, err // TODO
// 	}
// 	return output, nil
// }


func (pi *ProductImagePostgres) DeleteImage(image_id int, prouct_id int) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = array_remove(image_all, image_all[array_length(image_all, 1)]) WHERE array_length(image_all, 1) > 0;",productTable)
	_, err := pi.db.Exec(query)
	return err
}
