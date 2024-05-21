package repository

import (
	"errors"
	"fmt"
	"log"

	"database/sql"

	"github.com/DanilMankiev/sofia-app/entities"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db       *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignUp(user entity.SignUpInput, uid string) error {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return errors.New("internal server error")
	}

	query := fmt.Sprintf("INSERT INTO %s (id,name, surname,phone, email, password_hash) values ($1 ,$2, $3,$4,$5,$6)", usersTable)

	_, err = r.db.Exec(query, uid, user.Name, user.Surname, user.Phone, user.Email, hashedPassword)
	if err != nil {
		log.Printf("failed to insert user into database: %v", err)
		return errors.New("internal server error")
	}

	return  nil
}

func (r *AuthPostgres) SignIn(input entity.SignInInput) (string, error) {
	var client entity.User
	if err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", input.Email).Scan(&client.Name,&client.Surname,&client.Password_hash,&client.Phone,&client.Email,&client.Id); err != nil {
		if err == sql.ErrNoRows {
			return "",errors.New("user with email does not exist")
		}
		log.Printf("failed to get user by email from database: %v", err)
		return "",errors.New("internal server error")
	}

	//Compare input password and hash password 
	if err:= bcrypt.CompareHashAndPassword([]byte(client.Password_hash),[]byte(input.Password));err!=nil{
		return "",errors.New("incorrect password")
	}

	return client.Id,nil
}

func(r *AuthPostgres) CreateRefreshToken(uid string, refreshToken string) error{
	query:=fmt.Sprintf("INSERT INTO %s (id,token) values ($1,$2) ON CONFLICT (id) DO UPDATE SET token=$2", refreshTokenTable)
	if _,err:=r.db.Exec(query,uid,refreshToken);err!=nil{
		return errors.New("internal server error")
	}
	return nil
}

func (r *AuthPostgres) ValidateToken(token string, uid string) (bool,error){
	var dbtoken string
	query:=fmt.Sprintf("SELECT (token) FROM %s WHERE id=$1", refreshTokenTable)
	if err:=r.db.Get(&dbtoken,query,uid);err!=nil{
		return false,err
	}
	return token==dbtoken,nil
}