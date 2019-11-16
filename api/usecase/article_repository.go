package usecase

import (
	"api/model"
	"context"
)

type ArticleRepository interface {
	Store(ctx context.Context, a *model.Article) error
	Update(ctx context.Context, a *model.Article) error
	GetById(ctx context.Context, id int64) (*model.ViewArticle, error)
	GetArticleCount(ctx context.Context, searchParams map[string]string) (int, error)
	GetByPageNumber(ctx context.Context, n int, searchParams map[string]string) ([]model.ViewArticle, error)
	Delete(ctx context.Context, aId int64, un string) (int64, error)
}
