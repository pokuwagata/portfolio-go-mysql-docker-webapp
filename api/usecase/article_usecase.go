package usecase

import (
	"api/constant"
	"api/model"
	"api/repository"
	"context"
)

type ArticleUsecase struct {
	ar *repository.ArticleRepository
	ur *repository.UserRepository
	pr *repository.PostRepository
	su *SessionUsecase
}

func NewArticleUsecase(
	ar *repository.ArticleRepository,
	ur *repository.UserRepository,
	pr *repository.PostRepository,
	su *SessionUsecase) *ArticleUsecase {
	return &ArticleUsecase{ar, ur, pr, su}
}

func (au *ArticleUsecase) CreateArticle(ctx context.Context, a *model.Article, t string) error {
	// tokenから投稿者を設定
	u, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return err
	}

	p := &model.Post{Username: u}

	// 投稿イベントを作成
	pid, err := au.pr.Store(ctx, p)
	if err != nil {
		return err
	}
	a.PostId = pid

	// 投稿者を設定
	a.Username = u

	// 作成時の投稿ステータスは有効
	a.ArticleStatus = constant.PUBLISHED

	err = au.ar.Store(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (au *ArticleUsecase) GetMaxPageNumber(ctx context.Context, t string) (int, error) {
	// tokenから投稿者を設定
	un, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return 0, err
	}

	cnt, err := au.ar.GetArticleCount(ctx, un)
	if err != nil {
		return 0, err
	}

	max := cnt / constant.ARTICLES_PER_PAGE
	if cnt % constant.ARTICLES_PER_PAGE != 0 {
		max+=1
	}

	return max, nil;
}

func (au *ArticleUsecase) GetByPageNumber(ctx context.Context, n int, t string) ([]model.ViewArticle, error) {
	// tokenから投稿者を設定
	un, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return nil, err
	}

	articles, err := au.ar.GetByPageNumber(ctx, un, n)
	if err != nil {
		return nil, err
	}

	return articles, nil;
}
