package service

import (
	"context"
	"errors"
	"log"

	"os"
	"fmt"
	"firebase.google.com/go/v4/auth"
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
  
)

func GetPublicKey() string{
	var certPEM = os.Getenv("PRIVATE_KEY")

	return certPEM
}
func createUID() string {
	return uuid.NewString()
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UID string `json:"uid,omitempty"`
}

type AuthService struct {
	repo     repository.Authorization
	FireAuth *auth.Client
}

func newAuthService(repo repository.Authorization, Fireauth *auth.Client) *AuthService {
	return &AuthService{repo: repo, FireAuth: Fireauth}
}

func (s *AuthService) SignUp(user entity.SignUpInput) error {
	uid := createUID()

	err := s.repo.SignUp(user, uid)
	if err != nil {
		return errors.New("failed to signup")
	}

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password).
		PhoneNumber(user.Phone).
		DisplayName(user.Name)

	_, err = s.FireAuth.CreateUser(context.Background(), params)
	if err != nil {
		return errors.New("error creating user")
	}
	return nil
}
func (s *AuthService) SignIn(user entity.SignInInput) (string, error) {

	uid, err := s.repo.SignIn(user)
	if err != nil {
		return "", errors.New("failed to get user id")
	}
	//Generate token for the user
	token, err := s.FireAuth.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("failed to generate custom token:%v", err.Error())
		return "", errors.New("internal server error")
	}
	return token, nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(GetPublicKey()))
    if err != nil {
        fmt.Println("Error parsing private key:", err)
        return "",errors.New("error parsing private key")
    }

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return &key.PublicKey,nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UID, nil
}
