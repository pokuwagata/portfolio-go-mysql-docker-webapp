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
