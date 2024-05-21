package service

import (
	"context"
	"errors"
	"log"
	"time"

	"os"

	"firebase.google.com/go/v4/auth"
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	tokenTTL = time.Hour * 24 * 30
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
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password).
		PhoneNumber(user.Phone).
		DisplayName(user.Name)

	_, err := s.FireAuth.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.repo.SignUp(user, uid)
	if err != nil {
		return err
	}
	return nil
}
func (s *AuthService) SignIn(user entity.SignInInput) (string,string, error) {

	uid, err := s.repo.SignIn(user)
	if err != nil {
		return "","", err
	}
	refreshtoken_jwt:= jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		}, 
		uid,
	})
	
	accessToken, err := s.FireAuth.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("failed to generate custom token:%v", err.Error())
		return "","", errors.New("internal server error")
	}
	refreshToken,err:=refreshtoken_jwt.SignedString([]byte(os.Getenv("REFRESH_KEY")))
	if err!=nil{
		return "","",err
	}	
	err= s.repo.CreateRefreshToken(uid,refreshToken)
	if err!=nil{
		return "","",err
	}
	return accessToken,refreshToken,nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(GetPublicKey()))
    if err != nil {
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
		return "", errors.New("token claims are not valid")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
        return "", errors.New("token is expired")
	}
	return claims.UID, nil
}
func ParseRefreshToken(refreshToken string)(string, error){
	token,err:= jwt.ParseWithClaims(refreshToken,&tokenClaims{},func(token *jwt.Token) (interface{}, error) {
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,errors.New("invalid signing method")
		}
		return []byte(os.Getenv("REFRESH_KEY")), nil
		})
	if err!=nil{
		return "",err
	}
	claims,ok:=token.Claims.(*tokenClaims)
	if !ok{
		return "", errors.New("token claims not valid")
	}

	return claims.UID, nil
}

// func (s *AuthService) GenerateTokens(uid string)(string,string,error){
// 	refreshtoken_jwt:= jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
// 		}, 
// 		uid,
// 	})
	
// 	token, err := s.FireAuth.CustomToken(context.Background(), uid)
// 	if err != nil {
// 		log.Printf("failed to generate custom token:%v", err.Error())
// 		return "","", errors.New("internal server error")
// 	}
// 	refreshToken,err:=refreshtoken_jwt.SignedString([]byte(os.Getenv("REFRESH_KEY")))
// 	if err!=nil{
// 		return "","",err
// 	}
// 	err= s.repo.CreateRefreshToken(uid,refreshToken)
// 	if err!=nil{
// 		return "","",err
// 	}
// 	return  token,refreshToken,nil
// }

func (s * AuthService) RefreshToken(refreshToken string) (string,string,error){
	uid,err:=ParseRefreshToken(refreshToken)
	if err!=nil{
		return "","",err
	}
	ok, err:=s.repo.ValidateToken(refreshToken,uid)
	if err!=nil{
		return "","",err
	}
	if !ok{
		return "","",errors.New("wrong refresh token")
	}
	refreshtoken_jwt:= jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		}, 
		uid,
	})
	err= s.repo.CreateRefreshToken(uid,refreshtoken_jwt.Raw)
	if err!=nil{
		return "","",err
	}
	accessToken, err := s.FireAuth.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("failed to generate custom token:%v", err.Error())
		return "","", errors.New("internal server error")
	}
	refreshTokenNew,err:=refreshtoken_jwt.SignedString([]byte(os.Getenv("REFRESH_KEY")))
	if err!=nil{
		return "","",err
	}	

	return accessToken,refreshTokenNew,nil
}