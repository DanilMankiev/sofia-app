package repository

import (
	"fmt"

	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type BlogImagePostgres struct {
	db *sqlx.DB
}

func NewBlogImagePostgres(db *sqlx.DB) *BlogImagePostgres {
	return &BlogImagePostgres{db: db}
}

func (bp *BlogImagePostgres) CreateImage(input entity.ImageInputBlog) error {
	query := fmt.Sprintf("INSERT INTO %s (blog_id,img) values ($1,$2)", blogImages)
	_, err := bp.db.Exec(query, input.Blog_id, input.Image)
	if err != nil {
		return err
	}
	return nil
}

func (bp *BlogImagePostgres) GetAllImages(blog_id int) ([]string, error) {
	var output []string

	query := fmt.Sprintf("SELECT (img) from %s WHERE blog_id=$1", blogImages)

	if err := bp.db.Select(&output, query, blog_id); err != nil {
		return output, err // do some
	}
	return output, nil
}

func (bp *BlogImagePostgres) GetImageById(blog_id int, image_id int) (string, error) {
	var output string
	query := fmt.Sprintf("SELECT (img) from %s WHERE blog_id=$1 AND id = $2", blogImages)

	if err := bp.db.Get(&output, query, blog_id, image_id); err != nil {
		return output, err // do some
	}
	return output, nil
}

func (bp *BlogImagePostgres) DeleteImage(image_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", blogImages)
	_, err := bp.db.Exec(query, image_id)
	return err
}
