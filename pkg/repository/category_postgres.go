package repository

import (
	"fmt"
	"strings"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (l *CategoryPostgres) CreateCategory(Category entity.Category) (int, error) {
	var Category_id int

	query := fmt.Sprintf("INSERT INTO %s (name,description) values ($1,$2) RETURNING id", categoryTable)

	row := l.db.QueryRow(query, Category.Name, Category.Description)

	if err := row.Scan(&Category_id); err != nil {
		return 0, err
	}

	return Category_id, nil

}

func (l *CategoryPostgres) GetAllCategorys() ([]entity.Category, error) {
	var Categorys []entity.Category
	query := fmt.Sprintf("SELECT * FROM %s", categoryTable)
	err := l.db.Select(&Categorys, query)
	if err != nil {
		return nil, err
	}

	return Categorys, nil
}

func (l *CategoryPostgres) GetCategoryById(id int) (entity.Category, error) {
	var Category entity.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", categoryTable)

	err := l.db.Get(&Category, query, id)

	return Category, err
}
func (l *CategoryPostgres) DeleteCategory(id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id=$1", categoryTable)

	_, err := l.db.Exec(query, id)

	return err
}

func (l *CategoryPostgres) UpdateCategory(id int, input entity.UpdateCategoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name!= nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", categoryTable, setQuery, argId)

	args = append(args, id)

	_, err := l.db.Exec(query, args...)

	return err
}
