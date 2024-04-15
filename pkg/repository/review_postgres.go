package repository

import (
	"fmt"
	"strings"

	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres{
	return &ReviewPostgres{db:db}
}

func(rv *ReviewPostgres) CreateReview(input entity.CreateReview) (int,error){
	var review_id int

	query:=fmt.Sprintf("INSERT INTO %s (name,surname,description, topic) VALUES ($1,$2,$3,$4) RETURNING id",reviewTable)

	row:=rv.db.QueryRow(query,input.Name,input.Surname,input.Description,input.Topic)

	if err:=row.Scan(&review_id);err!=nil{
		return 0,err
	}

	return review_id,nil
}

func (rv *ReviewPostgres) GetAllReview() ([]entity.Review,error){
	var output []entity.Review

	query:=fmt.Sprintf("SELECT * FROM %s",reviewTable)

	if err:=rv.db.Select(&output,query);err!=nil{
		return output,err
	}
	return output,nil
}

func (rv *ReviewPostgres) DeleteReview(id int) error{
	query:=fmt.Sprintf("DELETE FROM %s WHERE id=$1",reviewTable)
	_,err:=rv.db.Exec(query,id)
	return err
} 

func (rv *ReviewPostgres) UpdateReview(id int, input entity.UpdateReview) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}
	if input.Topic!=nil{
		setValues = append(setValues, fmt.Sprintf("topic=$%d",argId))
		args = append(args, *input.Topic)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", reviewTable, setQuery, argId)

	args = append(args, id)

	_, err := rv.db.Exec(query, args...)

	return err
}