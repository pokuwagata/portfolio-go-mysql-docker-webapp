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

func (au *ArticleUsecaseMock) Update(ctx context.Context, a *model.Article, t string) error {
	return nil
}

func (au *ArticleUsecaseMock) GetById(ctx context.Context, id int64) (*model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleUsecaseMock) GetMaxPageNumber(ctx context.Context) (int, error) {
	return 0, nil
}

func (au *ArticleUsecaseMock) GetMaxPageNumberByUser(ctx context.Context, t string) (int, error) {
	return 0, nil
}

func (au *ArticleUsecaseMock) GetByPageNumber(ctx context.Context, n int) ([]model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleUsecaseMock) GetByUserAndPageNumber(ctx context.Context, n int, t string) ([]model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleUsecaseMock) Delete(ctx context.Context, articleId int64, token string) (int64, error) {
	return 0, nil
}
