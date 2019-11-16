package mocks

import (
	"api/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (ur *UserRepositoryMock) Store(ctx context.Context, u *model.User) error {
	return nil
}

func (ur *UserRepositoryMock) GetPassword(ctx context.Context, s *model.Session) (string, error) {
	return "", nil
}

func (ur *UserRepositoryMock) GetIdByUsername(ctx context.Context, name string) (int, error) {
	return 0, nil
}
