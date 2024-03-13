package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
	"github.com/Futturi/vktest/internal/utils"
	"github.com/golang-jwt/jwt"
)

const (
	jwtkey = "erwiughoweijfrpqiwjbnqhioqeuhghowqop[ew[slcnbtisdkcngueklalx]]"
)

type Auth_Service struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *Auth_Service {
	return &Auth_Service{repo: repo}
}

func (a *Auth_Service) SignUp(User models.User) (int, error) {
	user := models.User{Id: User.Id, Username: User.Username, Password: utils.HashPass(User.Password)}
	return a.repo.SignUp(user)
}

type Claims struct {
	Id      int `json:"id"`
	IsAdmin bool
	jwt.StandardClaims
}

func (a *Auth_Service) SignIn(User models.User) (string, error) {
	user := models.User{Username: User.Username, Password: utils.HashPass(User.Password)}
	id, err := a.repo.SignIn(user)
	if err != nil {
		return "", err
	}
	claims := Claims{id, false, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtkey))
}

func (a *Auth_Service) ParseToken(header string) (string, string, error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(jwtkey), nil
	})
	if err != nil {
		return "", "", err
	}
	Claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", "", errors.New("token claims are not of type *tokenClaims")
	}
	return fmt.Sprintf("%d", Claims.Id), fmt.Sprintf("%v", Claims.IsAdmin), nil
}

func (a *Auth_Service) SignUpAdmin(Admin models.Admin) (int, error) {
	admin := models.Admin{Username: Admin.Username, Password: utils.HashPass(Admin.Password)}
	return a.repo.SignUpAdmin(admin)
}

func (a *Auth_Service) SignInAdmin(Admin models.Admin) (string, error) {
	admin := models.Admin{Username: Admin.Username, Password: utils.HashPass(Admin.Password)}
	id, err := a.repo.SignInAdmin(admin)
	if err != nil {
		return "", err
	}
	claims := Claims{id, true, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtkey))
}
