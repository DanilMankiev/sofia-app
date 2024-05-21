package repository

import (
	"errors"
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

func (pi *ProductImagePostgres) CreateImage(input entity.ImageInputProduct) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = $1 WHERE id=$2", productTable)
	_, err := pi.db.Exec(query,input.Image,input.Product_id)
	if err != nil {
		return err
	}
	return nil
}


func (pi *ProductImagePostgres) DeleteImage(prouct_id int) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = array_remove(image_all, image_all[array_length(image_all, 1)]) WHERE array_length(image_all, 1) > 0 AND id=$1",productTable)
	_, err := pi.db.Exec(query, prouct_id)
	return err
}

func(pi* ProductImagePostgres) CreatePreviewImage(url string, id int) error{
	query:=fmt.Sprintf("UPDATE %s SET image_preview = $1 WHERE id=$2", productTable)
	_,err:= pi.db.Exec(query,url,id)
	if err != nil {
		return errors.New("failed on update image_preview")
	}
	return nil
}
