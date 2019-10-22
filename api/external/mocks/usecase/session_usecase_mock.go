package usecase

import (
	"api/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type SessionUsecaseMock struct {
	mock.Mock
}

func (su *SessionUsecaseMock) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	return "", nil
}

func (su *SessionUsecaseMock) GetUsernameFromToken(clientToken string) (string, error) {
	args := su.Called(clientToken)
	return args.String(0), args.Error(1)
}
