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

func (bp *BlogImagePostgres) DeleteImage(id int) error {
	query := fmt.Sprintf("UPDATE %s SET image_all = array_remove(image_all, image_all[array_length(image_all, 1)]) WHERE array_length(image_all, 1) > 0 AND id=$1;",blogTable)
	_, err := bp.db.Exec(query,id)
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