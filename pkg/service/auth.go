package service

import (

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

// const (
// 	salt       = "12e1dwedmcd23e1rfefwf"
// 	signingKey = "qwecmdwn732rme23f0w"
// 	tokenTTL   = 12 * time.Hour
// )

// type tokenClaims struct {
// 	jwt.StandardClaims
// 	UserId int `json:"user_id"`
// }

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(user entity.SignUpInput) (string, error) {
	return s.repo.SignUp(user)
}

func (s *AuthService) SignIn(user entity.SignInInput) (string, error) {
	return s.repo.SignIn(user)
}

// func (s *AuthService) GenegateToken(username, password string) (string, error) {
// 	user, err := s.repo.GetUser(username, generatePasswordHash(password))
// 	if err != nil {
// 		return "", err
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 		user.Id,
// 	})

// 	return token.SignedString([]byte(signingKey))

// }

// func generatePasswordHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }
