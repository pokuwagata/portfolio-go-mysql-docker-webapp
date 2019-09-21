package usecase

import (
	"api/external/jwtauth"
	"api/model"
	"api/repository"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type SessionUsecase struct {
	ur *repository.UserRepository
}

func NewSessionUsecase(ur *repository.UserRepository) *SessionUsecase {
	return &SessionUsecase{ur}
}

func (su *SessionUsecase) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	hash, err := su.ur.GetPassword(ctx, s)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s.Password)); err != nil {
		return "", errors.New("password is incorecct")
	}

	claims := &jwtauth.JwtCustomClaims{
		Name:     s.Username,
		Password: hash,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtauth.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (su *SessionUsecase) GetUsernameFromToken(t string) (string, error) {
	token, err := jwt.ParseWithClaims(t, &jwtauth.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtauth.SECRET_KEY), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwtauth.JwtCustomClaims); ok && token.Valid {
		return claims.Name, nil
	} else {
		return "", errors.New("token is invalid")
	}
}
