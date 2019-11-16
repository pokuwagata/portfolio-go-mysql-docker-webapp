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
	args := au.Called(ctx, a)
	if args.Get(0) != nil {
		return args.Error(0)
	} else {
		return nil
	}
}

func (au *ArticleRepositoryMock) GetById(ctx context.Context, id int64) (*model.ViewArticle, error) {
	args := au.Called(ctx, id)
	if args.Get(0) != nil {
		if a, ok := args.Get(0).(*model.ViewArticle); ok {
			return a, nil
		} else {
			return nil, nil
		}
	} else {
		return nil, args.Error(1)
	}
}

func (au *ArticleRepositoryMock) GetArticleCount(ctx context.Context, searchParams map[string]string) (int, error) {
	args := au.Called(ctx)
	if args.Get(0) != 0 {
		return args.Int(0), nil
	} else {
		return 0, args.Error(1)
	}
}

func (au *ArticleRepositoryMock) GetByPageNumber(ctx context.Context, n int, searchParams map[string]string) ([]model.ViewArticle, error) {
	return nil, nil
}

func (au *ArticleRepositoryMock) Delete(ctx context.Context, articleId int64, name string) (int64, error) {
	return 0, nil
}
