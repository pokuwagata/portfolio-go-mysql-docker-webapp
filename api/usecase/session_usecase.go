package usecase

import (
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

// 64 random hexadecimal characters (0-9 and A-F):
// https://www.grc.com/passwords.htm
const SECRET_KEY = "D01287268E2B030F0AE20D718F1EE60CA88FCDEC3D7F9E8368DF2209E90DC4E0"

func NewSessionUsecase(ur *repository.UserRepository) *SessionUsecase {
	return &SessionUsecase{ur}
}

func (su *SessionUsecase) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	// Check Username & Password
	hash, err := su.ur.GetPassword(ctx, s)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s.Password)); err != nil {
		return "", errors.New("password is incorecct")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = s.Username
	claims["password"] = hash
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return t, nil
}
