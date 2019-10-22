package mocks

import (
	"api/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type ArticleRepositoryMock struct {
	mock.Mock
}

func (au *ArticleRepositoryMock) Store(ctx context.Context, a *model.Article) error {
	args := au.Called(ctx, a)
	if args.Get(0) != nil {
		return args.Error(0)
	} else {
		return nil
	}
}

func (au *ArticleRepositoryMock) Update(ctx context.Context, a *model.Article) error {
	return nil
}

func (au *ArticleRepositoryMock) GetById(ctx context.Context, id int64) (*model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleRepositoryMock) GetArticleCount(ctx context.Context) (int, error) {
	return 0, nil
}

func (au *ArticleRepositoryMock) GetArticleCountByUser(ctx context.Context, un string) (int, error) {
	return 0, nil
}

func (au *ArticleRepositoryMock) GetByPageNumber(ctx context.Context, n int) ([]model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleRepositoryMock) GetByUserAndPageNumber(ctx context.Context, un string, n int) ([]model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleRepositoryMock) Delete(ctx context.Context, aId int64, un string) (int64, error) {
	return 0, nil
}
