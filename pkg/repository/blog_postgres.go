package repository

import (
	"fmt"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"

	"strings"
)


type BlogPostgres struct{
	db *sqlx.DB
}

func NewBlogPostgres(db *sqlx.DB) *BlogPostgres{
	return &BlogPostgres{db:db}
}

func(bp *BlogPostgres) CreateBlog(input entity.CreateBlog) (int,error){
	var id int 

	query:=fmt.Sprintf("INSERT INTO %s (topic,description) values ($1,$2) RETURNING id", blogTable)

	row:=bp.db.QueryRow(query,input.Topic,input.Description)

	if err:=row.Scan(&id);err!=nil{
		return 0,err
	}
	return id,nil
}

func (bp *BlogPostgres) GetAllBlog() ([]entity.Blog,error){
	var output []entity.Blog

	query:=fmt.Sprintf("SELECT * FROM %s",blogTable)

	if err:=bp.db.Select(&output,query);err!=nil{
		return output,err
	}
	return output,nil
}

func (bp *BlogPostgres) GetBlogById(id int) (entity.Blog,error){
	var blog entity.Blog

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", blogTable)

	if err := bp.db.Get(&blog, query, id); err != nil {
		return blog, err
	}

	return blog, nil
}

func (bp *BlogPostgres) DeleteBlog(id int) error{
	query:=fmt.Sprintf("DELETE FROM %s WHERE id=$1",blogTable)
	_,err:=bp.db.Exec(query,id)
	return err
} 

func (bp *BlogPostgres) UpdateBlog(id int, input entity.UpdateBlog) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Topic != nil {
		setValues = append(setValues, fmt.Sprintf("topic=$%d", argId))
		args = append(args, *input.Topic)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", blogTable, setQuery, argId)

	args = append(args, id)

	_, err := bp.db.Exec(query, args...)

	return err
}

