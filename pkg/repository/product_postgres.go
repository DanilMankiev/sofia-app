package repository

import (
	"errors"
	"fmt"
	"strings"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (it *ProductPostgres) GetAllItems(category_id int) ([]entity.Product, error) {
	var products []entity.Product

	query := fmt.Sprintf("SELECT * FROM %s WHERE category_id=$1", productTable)
	err := it.db.Select(&products, query,category_id)
	if err != nil {
		return nil, err
	}
	return products, nil

}

func (it *ProductPostgres) CreateProduct(category_id int, input entity.CreateProduct) (int, error) {
	var product_id int
	var category_name string

	query := fmt.Sprintf("SELECT (name) FROM %s WHERE id=$1", categoryTable)

	if err := it.db.Get(&category_name, query, category_id); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (name, category_id,category, description_preview,description_full,image_preview,image_all,composition,purchase,delivery, price,furniture) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id", productTable)

	
	row := it.db.QueryRow(query, input.Name, category_id, category_name, input.DescriptionPreview, input.FullDescription, input.ImagePreview, input.AllImages, input.Composition, input.TermsPurchase, input.Delivery, input.Price,input.Furniture) // TODO

	if err := row.Scan(&product_id); err != nil {
		return 0, err
	}

	return product_id,nil

}

func (it *ProductPostgres) GetItemByid(product_id int) (entity.Product, error) {
	var product entity.Product

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productTable)

	if err := it.db.Get(&product, query, product_id); err != nil {
		return product, err
	}

	return product, nil
}

func (it *ProductPostgres) DeleteItem(product_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", productTable)
	_, err := it.db.Exec(query, product_id)
	return err
}

func (it *ProductPostgres) UpdateItem(product_id int, input entity.UpdateProductInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.CategoryId != nil {
		setValues = append(setValues, fmt.Sprintf("category_id=$%d", argId))
		args = append(args, *input.CategoryId)
		argId++
	}
	if input.CategoryName != nil {
		setValues = append(setValues, fmt.Sprintf("category=$%d", argId))
		args = append(args, *input.CategoryName)
		argId++
	}
	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}
	if input.Composition != nil {
		setValues = append(setValues, fmt.Sprintf("composition=$%d", argId))
		args = append(args, *input.Composition)
		argId++
	}
	if input.Delivery != nil {
		setValues = append(setValues, fmt.Sprintf("delivery=$%d", argId))
		args = append(args, *input.Delivery)
		argId++
	}
	if input.DescriptionPreview != nil {
		setValues = append(setValues, fmt.Sprintf("description_preview=$%d", argId))
		args = append(args, *input.DescriptionPreview)
		argId++
	}
	if input.FullDescription != nil {
		setValues = append(setValues, fmt.Sprintf("description_full=$%d", argId))
		args = append(args, *input.FullDescription)
		argId++
	}
	if input.TermsPurchase != nil {
		setValues = append(setValues, fmt.Sprintf("purchase=$%d", argId))
		args = append(args, *input.TermsPurchase)
		argId++
	}
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Furniture != nil {
		setValues = append(setValues, fmt.Sprintf("furniture=$%d", argId))
		args = append(args, *input.Furniture)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", productTable, setQuery, argId)

	args = append(args, product_id)

	_, err := it.db.Exec(query, args...)
	if err!=nil{
		return errors.New(err.Error())
	}
	return err
}
