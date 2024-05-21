package repository

import (
	"errors"
	"fmt"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) GetUser(uid string) (entity.UserDisplay,error){
	var user entity.UserDisplay
	query:=fmt.Sprintf("SELECT id,name,surname,email,phone FROM %s WHERE id = $1", usersTable)

	if err := up.db.Get(&user, query, uid); err != nil {
		return user, err
	}
	return user,nil
}

func (up *UserPostgres) CreateFavorites(uid string, id int) error {
	query:= fmt.Sprintf("INSERT INTO %s (uid, id) values ($1,$2)", favoritesTable)
	_, err:=up.db.Exec(query,uid,id)
	if err!=nil{
		return errors.New("failed add to favorites")
	}
	return nil
}

func (up *UserPostgres) GetAllFavorites(uid string) ([]entity.Product,error){
	var products []entity.Product
	query:= fmt.Sprintf("SELECT p.id,p.name,p.category_id,p.description_preview,p.price,p.category,p.delivery,p.purchase,p.composition,p.image_preview,p.description_full,p.image_all,p.furniture FROM %s p JOIN %s u ON p.id=u.id WHERE u.uid=$1;",productTable,favoritesTable)
	err:=up.db.Select(&products,query,uid)
	if err!=nil{
		return products,errors.New("failed to get all favorites")
	}
	return products,nil
}

func (up *UserPostgres) DeleteFavorites(uid string,id int) error{
	query:=fmt.Sprintf("DELETE FROM %s WHERE uid=$1 and id=$2",favoritesTable)
	_,err:=up.db.Exec(query,uid,id)
	if err!=nil{
		return err
	}
	return nil
}