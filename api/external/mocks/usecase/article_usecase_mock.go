package usecase

import (
	"api/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type ArticleUsecaseMock struct {
	mock.Mock
}

func (au *ArticleUsecaseMock) Create(ctx context.Context, a *model.Article, t string) error {
	args := au.Called(ctx, a, t)
	if args.Get(0) != nil {
		return args.Error(0)
	} else {
		return nil
	}
}
