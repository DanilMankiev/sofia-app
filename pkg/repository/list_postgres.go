package repository

import (
	"fmt"
	"strings"

	sofia "github.com/DanilMankiev/sofia-app"
	"github.com/jmoiron/sqlx"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (l *ListPostgres) CreateList(list sofia.List) (int, error) {
	var list_id int

	query := fmt.Sprintf("INSERT INTO %s (listname,description) values ($1,$2) RETURNING list_id", listsTable)

	row := l.db.QueryRow(query, list.Listname, list.Description)

	if err := row.Scan(&list_id); err != nil {
		return 0, err
	}

	return list_id, nil

}

func (l *ListPostgres) GetAllLists() ([]sofia.List, error) {
	var lists []sofia.List
	query := fmt.Sprintf("SELECT * FROM %s", listsTable)
	err := l.db.Select(&lists, query)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (l *ListPostgres) GetListById(id int) (sofia.List, error) {
	var list sofia.List

	query := fmt.Sprintf("SELECT * FROM %s WHERE list_id= $1", listsTable)

	err := l.db.Get(&list, query, id)

	return list, err
}
func (l *ListPostgres) DeleteList(id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.list_id=$1", listsTable)

	_, err := l.db.Exec(query, id)

	return err
}

func (l *ListPostgres) UpdateList(id int, input sofia.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Listname != nil {
		setValues = append(setValues, fmt.Sprintf("listname=$%d", argId))
		args = append(args, *input.Listname)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE list_id = $%d", listsTable, setQuery, argId)

	args = append(args, id)

	_, err := l.db.Exec(query, args...)

	return err
}
