package repository

import (
	"fmt"

	sofia "github.com/DanilMankiev/todo-app"
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

	query := fmt.Sprintf("INSERT INTO %s (listname) values($1) returning list_id", listsTable)

	row := l.db.QueryRow(query, list.Listname)

	if err := row.Scan(&list_id); err != nil {
		return 0, err
	}

	return list_id, nil

}
