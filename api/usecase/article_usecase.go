package usecase

import (
	"api/constant"
	"api/model"
	"context"
)

type ArticleUsecase struct {
	ar ArticleRepository
	ur UserRepository
	su SessionUsecaseInterface
}

func NewArticleUsecase(
	ar ArticleRepository,
	ur UserRepository,
	su SessionUsecaseInterface) *ArticleUsecase {
	return &ArticleUsecase{ar, ur, su}
}

func (au *ArticleUsecase) Create(ctx context.Context, a *model.Article, t string) error {
	// tokenから投稿者を設定
	u, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return err
	}

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

func (au *ArticleUsecase) Update(ctx context.Context, a *model.Article, t string) error {
	// tokenから投稿者を設定
	u, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return err
	}

	// 投稿者を設定
	a.Username = u

	err = au.ar.Update(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (au *ArticleUsecase) GetById(ctx context.Context, id int64) (*model.ViewArticle, error) {
	article, err := au.ar.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (au *ArticleUsecase) GetMaxPageNumber(ctx context.Context) (int, error) {
	cnt, err := au.ar.GetArticleCount(ctx)
	if err != nil {
		return 0, err
	}

	max := cnt / constant.ARTICLES_PER_PAGE
	if cnt%constant.ARTICLES_PER_PAGE != 0 {
		max += 1
	}

	return max, nil
}

func (au *ArticleUsecase) GetMaxPageNumberByUser(ctx context.Context, t string) (int, error) {
	// tokenから投稿者を設定
	un, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return 0, err
	}

	cnt, err := au.ar.GetArticleCountByUser(ctx, un)
	if err != nil {
		return 0, err
	}

	max := cnt / constant.ARTICLES_PER_PAGE
	if cnt%constant.ARTICLES_PER_PAGE != 0 {
		max += 1
	}

	return max, nil
}

func (au *ArticleUsecase) GetByPageNumber(ctx context.Context, n int) ([]model.ViewArticle, error) {
	articles, err := au.ar.GetByPageNumber(ctx, n)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (au *ArticleUsecase) GetByUserAndPageNumber(ctx context.Context, n int, t string) ([]model.ViewArticle, error) {
	// tokenから投稿者を設定
	un, err := au.su.GetUsernameFromToken(t)
	if err != nil {
		return nil, err
	}

	articles, err := au.ar.GetByUserAndPageNumber(ctx, un, n)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (au *ArticleUsecase) Delete(ctx context.Context, articleId int64, token string) (int64, error) {
	name, err := au.su.GetUsernameFromToken(token)
	if err != nil {
		return 0, err
	}

	id, err := au.ar.Delete(ctx, articleId, name)
	if err != nil {
		return 0, err
	}

	return id, nil
}
