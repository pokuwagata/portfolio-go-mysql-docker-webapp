package usecase

import (
	"api/constant"
	"api/external/jwtauth"
	"api/model"
	"api/repository"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type SessionUsecaseInterface interface {
	CreateSession(ctx context.Context, s *model.Session) (string, error)
	GetUsernameFromToken(t string) (string, error)
}

type SessionUsecase struct {
	ur *repository.UserRepository
	e  *echo.Echo
}

func NewSessionUsecase(ur *repository.UserRepository, e *echo.Echo) *SessionUsecase {
	return &SessionUsecase{ur, e}
}

func (su *SessionUsecase) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	hash, err := su.ur.GetPassword(ctx, s)
	if err != nil {
		err = errors.New(constant.ERR_SIGNUP_FAILED)
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s.Password)); err != nil {
		su.e.Logger.Errorf(constant.ERR_APP_ERROR, err)
		su.e.Logger.Debugf(constant.ERR_APP_ERROR_DEBUG, errors.WithStack(err))
		return "", errors.New(constant.ERR_SIGNUP_FAILED)
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

func (su *SessionUsecase) GetUsernameFromToken(clientToken string) (string, error) {
	token, err := jwt.ParseWithClaims(clientToken, &jwtauth.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtauth.SECRET_KEY), nil
	})
	if err != nil {
		su.e.Logger.Errorf(constant.ERR_APP_ERROR, err)
		su.e.Logger.Debugf(constant.ERR_APP_ERROR_DEBUG, errors.WithStack(err))
		return "", err
	}

	if claims, ok := token.Claims.(*jwtauth.JwtCustomClaims); ok && token.Valid {
		return claims.Name, nil
	} else {
		su.e.Logger.Errorf(constant.ERR_APP_ERROR, err)
		su.e.Logger.Debugf(constant.ERR_APP_ERROR_DEBUG, errors.WithStack(err))
		return "", errors.New(constant.ERR_INVALID_TOKEN)
	}
}
