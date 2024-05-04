package service

import (
	"context"
	"errors"
	"log"

	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"firebase.google.com/go/v4/auth"
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetPublicKey() *rsa.PublicKey{
	var certPEM = "-----BEGIN CERTIFICATE-----\nMIIDHTCCAgWgAwIBAgIJALeM3DBeoI3wMA0GCSqGSIb3DQEBBQUAMDExLzAtBgNV\nBAMMJnNlY3VyZXRva2VuLnN5c3RlbS5nc2VydmljZWFjY291bnQuY29tMB4XDTI0\nMDQzMDA3MzIyMVoXDTI0MDUxNjE5NDcyMVowMTEvMC0GA1UEAwwmc2VjdXJldG9r\nZW4uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wggEiMA0GCSqGSIb3DQEBAQUA\nA4IBDwAwggEKAoIBAQCZ5H6KYWuP1+SwCsN9tQXsD4JXii7FEJr/NGoQnHePBobr\nOzHaSdyxov7o3XqtouXmRDetVpANdph2r5+rTFY1231KehQF9HosYDA4zT/Ph32Z\n+kpS9Xlkg+515lowQtYGnwlAmnHirivTgUmrHR7GaVVOo7K5erD1tbFiIjTgtHNR\njlxVS786WEdvOkVodJQcKX5/5FDlI01AAbnbLf+iKpCq/bXNGFQI/6r47TTo9qEm\nUAHoPaqW2LceGpH1qqoBoBRfsS4qaWbAxsHjs0cur4x4Ai1c+iJbFNHQfbis/BzM\n5d5DLDFV4n3ZPie03aJoonWbiJX1eTW1No4XyozfAgMBAAGjODA2MAwGA1UdEwEB\n/wQCMAAwDgYDVR0PAQH/BAQDAgeAMBYGA1UdJQEB/wQMMAoGCCsGAQUFBwMCMA0G\nCSqGSIb3DQEBBQUAA4IBAQBGERUt+83Ar/OjpwpG9n1hsgM5X5TBrZXMPLpzlr0Y\nDOSB3svrvwBOcJftddUIStJKaEaFwuK+N6TuxtYbcE8tBF7QG1H1M7OdIb8j1o4j\naGggP9ziXiFgRHBADd8o4gHgeBygfZQUU73XHDu1jSzNsUELF0mUt5ffKxSoRtq2\ne1ng74n9sBmExN7HNW8DnyXyF21AnFeCqY3ttTY4KttsGKIXJB1PKXZ31wbTTeVH\njmn+QRC6co2ENNCgCtWr1GiBrgkve8HbtR1qbSDnpBiGAdH+yxBWCRNTEEPW4E7b\nZTGhgbFh1YNFf/+ihvomrfCeCdfwbQEkvs6hhQAI4nTC\n-----END CERTIFICATE-----\n"

    // Парсинг PEM-данных сертификата
    block, _ := pem.Decode([]byte(certPEM))
    if block == nil {
        log.Printf("Ошибка парсинга PEM блока")
    }

    // Извлечение открытого ключа из PEM-блока
    cert, err := x509.ParseCertificate(block.Bytes)
    if err != nil {
        log.Printf("Ошибка извлечения открытого ключа: ", err)
    }
	
	return cert.PublicKey.(*rsa.PublicKey)
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
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok:= token.Method.(*jwt.SigningMethodRSA);!ok{
			return nil, fmt.Errorf("неверный алгоритм подписи: %v", token.Header["alg"])
		}
		return GetPublicKey(),nil
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
