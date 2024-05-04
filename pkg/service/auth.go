package service

import (
	"context"
	"errors"
	"time"

	"fmt"

	"firebase.google.com/go/v4/auth"
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"crypto/sha1"
)


func createUID() string {
	return uuid.NewString()
}

const (
	signingKey="129rh3r1m9jfiunedhwbivjef09je2ewdew31d39j"
	salt = "newfu9uefp9832hf3r9f2pqfic29qf38nc"
	tokenTTL = time.Hour * 24 *30
)

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

	//Generate token for the user
	uid, err := s.repo.SignIn(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		uid,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}