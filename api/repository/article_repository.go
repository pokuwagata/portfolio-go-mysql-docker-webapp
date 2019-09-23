package repository

import (
	"api/constant"
	"api/model"
	"context"
	"database/sql"
	"errors"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func (ar *ArticleRepository) Store(ctx context.Context, a *model.Article) error {
	query := `INSERT articles SET title=?, content=?, post_id=?, ` +
		`user_id=(SELECT id FROM users where username=?), ` +
		`article_status_id=(SELECT id FROM article_statuses where status= ?)`
	stmt, err := ar.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.Title, a.Content, a.PostId, a.Username, a.ArticleStatus)
	if err != nil {
		return err
	}

	return nil
}

func (ar *ArticleRepository) GetByPageNumber(ctx context.Context, un string, n int) ([]model.ViewArticle, error) {
	uidQuery := `select id from users where username = ?`
	stmt, err := ar.db.PrepareContext(ctx, uidQuery)
	if err != nil {
		return nil, err
	}
	var uid int
	if err := stmt.QueryRowContext(ctx, un).Scan(&uid); err != nil {
		return nil, errors.New("user not found")
	}

	var rows *sql.Rows
	if n == 1 {
		atcQuery := `SELECT id, title, content, updated_at FROM articles ` +
			`WHERE user_id = ? AND article_status_id = (SELECT id FROM article_statuses WHERE status = ? ) ` +
			`ORDER BY updated_at DESC LIMIT ?`

		stmt, err = ar.db.PrepareContext(ctx, atcQuery)
		if err != nil {
			return nil, err
		}

		rows, err = stmt.QueryContext(ctx, uid, constant.PUBLISHED, constant.ARTICLES_PER_PAGE)
		if err != nil {
			return nil, err
		}

	} else {
		// updated_atが同一のレコードが存在しない想定
		atcQuery := `SELECT id, title, content, updated_at FROM articles WHERE updated_at < ` +
			`(SELECT updated_at FROM articles WHERE user_id = ? AND article_status_id = (SELECT id FROM article_statuses WHERE status = ? ) ` +
			`ORDER BY updated_at DESC LIMIT 1 OFFSET ?) ` +
			`AND user_id = ? AND article_status_id = (SELECT id FROM article_statuses WHERE status = ? ) ` + 
			`ORDER BY updated_at DESC LIMIT ?`

		stmt, err = ar.db.PrepareContext(ctx, atcQuery)
		if err != nil {
			return nil, err
		}

		offset := constant.ARTICLES_PER_PAGE*(n-1) - 1

		rows, err = stmt.QueryContext(ctx, uid, constant.PUBLISHED, offset, uid, constant.PUBLISHED, constant.ARTICLES_PER_PAGE)
		if err != nil {
			return nil, err
		}
	}

	var articles []model.ViewArticle

	if rows.Next() {
		var a model.ViewArticle
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt)
		articles = append(articles, a)
	} else {
		return nil, errors.New("article not found")
	}

	for rows.Next() {
		var a model.ViewArticle
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt)
		articles = append(articles, a)
	}

	return articles, nil
}
