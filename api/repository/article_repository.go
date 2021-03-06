package repository

import (
	"api/constant"
	"api/model"
	"context"
	"database/sql"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"strings"
)

type ArticleRepository struct {
	db *sql.DB
	e *echo.Echo
}

func NewArticleRepository(db *sql.DB, e *echo.Echo) *ArticleRepository {
	return &ArticleRepository{db, e}
}

func (ar *ArticleRepository) Store(ctx context.Context, a *model.Article) error {
	query := []string {
		"INSERT",
			"articles",
		"SET",
			"title = ?,",
			"content = ?,",
			"user_id =(",
				"SELECT",
					"id",
				"FROM",
					"users",
				"WHERE",
					"username = ?",
			"),",
			"article_status_id =(",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
			")",
	}
	rawQuery := strings.Join(query, constant.HALF_SPACE)

	if _, err := ar.db.ExecContext(
		ctx, rawQuery, a.Title, a.Content, a.Username, a.ArticleStatus); err != nil {
			ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
			ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
		return err
	}

	return nil
}

func (ar *ArticleRepository) Update(ctx context.Context, a *model.Article) error {
	query := []string {
		"UPDATE",
			"articles",
		"SET",
			"title = ?,",
			"content = ?,",
			"user_id =(",
				"SELECT",
					"id",
				"FROM",
					"users",
				"WHERE",
					"username = ?",
			"),",
			"article_status_id =(",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
			")",
		"WHERE",
			"id = ?",
	}

	rawQuery := strings.Join(query, constant.HALF_SPACE)

	if _, err := ar.db.ExecContext(ctx, rawQuery, a.Title, a.Content, a.Username, constant.PUBLISHED, a.ID); err != nil {
			ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
			ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
		return err
	}

	return nil
}

func (ar *ArticleRepository) GetById(ctx context.Context, id int64) (*model.ViewArticle, error) {
	query := []string{
		"SELECT",
			"id,",
			"title,",
			"content,",
			"updated_at,",
			"(",
				"SELECT",
					"username",
				"FROM",
					"users",
				"WHERE",
					"id = articles.user_id",
			") ",
		"FROM",
			"articles",
		"WHERE",
			"ID = ?",
			"AND article_status_id =(",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
			")",
	}

	rawQuery := strings.Join(query, constant.HALF_SPACE)

	var a model.ViewArticle
	if err := ar.db.QueryRowContext(ctx, rawQuery, id, constant.PUBLISHED).
		Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt, &a.Username); err != nil {
			ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
			ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
			return nil, err
	}

	return &a, nil
}

func (ar *ArticleRepository) GetArticleCount(ctx context.Context, searchParams map[string]string) (int, error) {
	args := []interface{}{constant.PUBLISHED}

	base := []string{
		"SELECT",
			"count(id)",
		"FROM",
			"articles",
		}

	where := []string{
		"WHERE",
			"article_status_id = (",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
				")",
		}
		
	for key, value := range searchParams {
		where = append(where, []string{
			"AND",
				key + " " + "=" + " " + "?",
		}...)
		args = append(args, value)
	}

	query := append(base, where...)
	rawQuery := strings.Join(query, constant.HALF_SPACE)

	var count int
	if err := ar.db.QueryRowContext(ctx, rawQuery, args...).Scan(&count); err != nil {
		ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
		ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
		return 0, err
	}

	return count, nil
}

func (ar *ArticleRepository) GetByPageNumber(ctx context.Context, n int, searchParams map[string]string) ([]model.ViewArticle, error) {
	var rows *sql.Rows
	// NOTE: index利用のため（サブクエリを使用しないため）にクエリを分割
	artStaId, err := ar.getArticleStatusId(ctx, constant.PUBLISHED)
	if err != nil {
		return nil, err
	}

	args := []interface{}{artStaId}

	base := []string{
		"SELECT",
			"id,",
			"title,",
			"content,",
			"updated_at,",
			"(",
				"SELECT",
					"username",
				"FROM users",
				"WHERE",
					"id = articles.user_id",
			")",
		"FROM articles",
	}

	where := []string{
		"WHERE",
			"article_status_id = ?",
	}

	if n > 1 {
		where = append(where, []string{
			"AND updated_at < (",
			"SELECT",
				"updated_at",
			"FROM articles",
			"WHERE",
				"article_status_id = ?",
			"ORDER BY",
				"updated_at DESC",
			"LIMIT",
				"1 OFFSET ?",
		")",
		}...)

		offset := constant.ARTICLES_PER_PAGE*(n-1) - 1

		args = append(args, []interface{}{artStaId, offset}...)
	}

	for key, value := range searchParams {
		where = append(where, []string{
			"AND",
				key + " " + "=" + " " + "?",
		}...)
		args = append(args, value)
	}

	order := []string{
		"ORDER BY",
			"updated_at DESC",
	}

	limit := []string{
		"LIMIT",
			"?",
	}

	args = append(args, constant.ARTICLES_PER_PAGE)

	query := append(base, where...)
	query = append(query, order...)
	query = append(query, limit...)
	
	rawQuery := strings.Join(query, constant.HALF_SPACE);

	rows, err = ar.db.QueryContext(ctx, rawQuery, args...)
	if err != nil {
		ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
		ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
		return nil, err
	}

	var articles []model.ViewArticle

	if rows.Next() {
		var a model.ViewArticle
		_ = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt, &a.Username)
		articles = append(articles, a)
	} else {
		// 0件の場合は空を渡す
		return []model.ViewArticle{}, nil
	}

	for rows.Next() {
		var a model.ViewArticle
		_ = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt, &a.Username)
		articles = append(articles, a)
	}

	return articles, nil
}

func (ar *ArticleRepository) Delete(ctx context.Context, articleId int64, name string) (int64, error) {
	query := []string{
		"UPDATE articles",
		"SET",
			"article_status_id = (",
				"SELECT",
					"id",
				"FROM article_statuses",
				"WHERE",
					"status = ?",
			")",
		"WHERE",
			"id = ?",
			"AND user_id = (",
				"SELECT",
					"id",
				"FROM users",
				"WHERE",
					"username = ?",
			")",
	}

	rawQuery := strings.Join(query, constant.HALF_SPACE);

	res, err := ar.db.ExecContext(ctx, rawQuery, constant.REMOVED, articleId, name)
	if err != nil {
		ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
		ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (ar *ArticleRepository) getArticleStatusId(ctx context.Context, status string) (int64, error) {
	query := []string{
		"SELECT",
			"id",
		"FROM",
			"article_statuses",
		"WHERE",
			"status = ?",
	}

	rawQuery := strings.Join(query, constant.HALF_SPACE)
	var id int64
	if err:= ar.db.QueryRowContext(ctx, rawQuery, status).Scan(&id); err != nil {
		ar.e.Logger.Errorf(constant.ERR_SQL_MESSAGE, err)
		ar.e.Logger.Debugf(constant.ERR_SQL_MESSAGE_DEBUG, errors.WithStack(err))
		return 0, err
	}

	return id, nil
}
