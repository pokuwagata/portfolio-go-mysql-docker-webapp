package repository

import (
	"api/model"
	"context"
	"database/sql"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func(ar *ArticleRepository) Store(ctx context.Context, a *model.Article) error {
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