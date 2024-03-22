package repository

import (
	"fmt"
	"strings"

	"github.com/DanilMankiev/sofia-app"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (it *ProductPostgres) GetAllItems(list_id int) ([]sofia.Product, error) {
	var products []sofia.Product

	query := fmt.Sprintf("SELECT * FROM %s", productTable)
	err := it.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil

}

func (it *ProductPostgres) CreateProduct(list_id int, input sofia.CreateProduct) (int, error) {
	var product_id int

	query := fmt.Sprintf("INSERT INTO %s (product_name, list_id, description, price) values ($1,$2,$3,$4) RETURNING product_id", productTable)

	row := it.db.QueryRow(query, input.Product_name, list_id, input.Description, input.Price)

	if err := row.Scan(&product_id); err != nil {
		return 0, err
	}

	return product_id, nil

}

func (it *ProductPostgres) GetItemByid(product_id int) (sofia.Product, error) {
	var product sofia.Product

	query := fmt.Sprintf("SELECT * FROM %s WHERE product_id=&1", productTable)

	if err := it.db.Get(&product, query, product_id); err != nil {
		return product, err
	}

	return product, nil
}

func (it *ProductPostgres) DeleteItem(product_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE product_id=$1", productTable)
	_, err := it.db.Exec(query, product_id)
	return err
}

func (it *ProductPostgres) UpdateItem(product_id int, input sofia.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.List_id != nil {
		setValues = append(setValues, fmt.Sprintf("list_id=$%d", argId))
		args = append(args, *input.List_id)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}
	if input.Product_name != nil {
		setValues = append(setValues, fmt.Sprintf("product_name=$%d", argId))
		args = append(args, *input.Product_name)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE product_id = $%d", productTable, setQuery, argId)

	args = append(args, product_id)

	_, err := it.db.Exec(query, args...)

	return err
}
