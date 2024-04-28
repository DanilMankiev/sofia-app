package repository

import (
	"fmt"

	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
	"errors"
)

type BlogImagePostgres struct {
	db *sqlx.DB
}

func NewBlogImagePostgres(db *sqlx.DB) *BlogImagePostgres {
	return &BlogImagePostgres{db: db}
}

func (bp *BlogImagePostgres) CreateImage(input entity.ImageInputBlog) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = $1 WHERE id=$2", blogTable)
	_, err := bp.db.Exec(query,input.Image,input.Blog_id)
	if err != nil {
		return err
	}
	return nil
}

// func (bp *BlogImagePostgres) GetAllImages(blog_id int) ([]string, error) {
// 	var output []string

// 	query := fmt.Sprintf("SELECT (img) from %s WHERE blog_id=$1", blogImagesTable)

// 	if err := bp.db.Select(&output, query, blog_id); err != nil {
// 		return output, err // do some
// 	}
// 	return output, nil
// }

// func (bp *BlogImagePostgres) GetImageById(blog_id int, image_id int) (string, error) {
// 	var output string
// 	query := fmt.Sprintf("SELECT (img) from %s WHERE blog_id=$1 AND id = $2", blogImagesTable)

// 	if err := bp.db.Get(&output, query, blog_id, image_id); err != nil {
// 		return output, err // do some
// 	}
// 	return output, nil
// }

func (bp *BlogImagePostgres) DeleteImage(image_id int) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = array_remove(image_all, image_all[array_length(image_all, 1)]) WHERE array_length(image_all, 1) > 0;",blogTable)
	_, err := bp.db.Exec(query)
	return err
}

func(bp * BlogImagePostgres) CreatePreviewImage(url string,id int) error{
	query:=fmt.Sprintf("UPDATE %s SET image_preview = $1 WHERE id=$2", blogTable)
	_,err:= bp.db.Exec(query,url,id)
	if err != nil {
		return errors.New("failed on update image_preview")
	}
	return nil
}