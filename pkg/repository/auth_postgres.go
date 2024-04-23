package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"database/sql"
	
	"firebase.google.com/go/v4/auth"
	
	"github.com/DanilMankiev/sofia-app/entities"
	
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db       *sqlx.DB
	FireAuth *auth.Client
}

func NewAuthPostgres(db *sqlx.DB, FireAuth *auth.Client) *AuthPostgres {
	return &AuthPostgres{db: db,FireAuth:FireAuth}
}

func (r *AuthPostgres) SignUp(user entity.SignUpInput) (string, error) {
	var client entity.User
	if err := r.db.QueryRow("SELECT * FROM %s WHERE email=$1 OR phone=$2", usersTable, user.Email, user.Phone).Scan(&client); err != nil {
		log.Printf("failed to get user by email or phone:%v", err.Error())
	}

	if client.Id != "" {
		return "", errors.New("user with email already exists")
	}
	// Generate uid for user
	uid := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return "", errors.New("internal server error")
	}

	query := fmt.Sprintf("INSERT INTO %s (id,name, surname,phone, email, password_hash) values ($1 ,$2, $3,$4,$5,$6)", usersTable)

	_, err = r.db.Exec(query, uid, user.Name, user.Surname, user.Phone, user.Email, hashedPassword)
	if err != nil {
		log.Printf("failed to insert user into database: %v", err)
		return "", errors.New("internal server error")
	}

	// Create custom token
	customToken, err := r.FireAuth.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("failed to create custom token: %v", err.Error())
		return "", errors.New("internal server error")
	}
	return customToken, nil
}

func (r *AuthPostgres) SignIn(input entity.SignInInput) (string, error) {
	var client entity.User
	if err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", input.Email).Scan(&client.Name,&client.Surname,&client.Password_hash,&client.Phone,&client.Email,&client.Id); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user with email does not exist")
		}
		log.Printf("failed to get user by email from database: %v", err)
		return "", errors.New("internal server error")
	}

	//Compare input password and hash password 
	if err:= bcrypt.CompareHashAndPassword([]byte(client.Password_hash),[]byte(input.Password));err!=nil{
		return "", errors.New("incorrect password")
	}

	//Generate token for the user
	token,err:= r.FireAuth.CustomToken(context.Background(),client.Id)
	if err!=nil{
		log.Printf("failed to generate custom token:%v",err.Error())
		return "", errors.New("internal server error")
	}

	return token,nil

}
