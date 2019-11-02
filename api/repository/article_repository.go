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
	query :=
		`INSERT
			articles
		SET
			title = ?,
			content = ?,
			user_id =(
				SELECT
					id
				FROM
					users
				WHERE
					username = ?
			),
			article_status_id =(
				SELECT
					id
				FROM
					article_statuses
				WHERE
					status = ?
			)`

	if _, err := ar.db.ExecContext(
		ctx, query, a.Title, a.Content, a.Username, a.ArticleStatus); err != nil {
		return err
	}

	return nil
}

func (ar *ArticleRepository) Update(ctx context.Context, a *model.Article) error {
	query :=
		`
		UPDATE
			articles
		SET
			title = ?,
			content = ?,
			user_id =(
				SELECT
					id
				FROM
					users
				WHERE
					username = ?
			),
			article_status_id =(
				SELECT
					id
				FROM
					article_statuses
				WHERE
					status = ?
			)
		WHERE
			id = ?
	`
	if _, err := ar.db.ExecContext(ctx, query, a.Title, a.Content, a.Username, constant.PUBLISHED, a.ID); err != nil {
		return err
	}

	return nil
}

func (ar *ArticleRepository) GetById(ctx context.Context, id int64) (*model.ViewArticle, error) {
	query := 
		`
		SELECT
			id,
			title,
			content,
			updated_at,
			(
				SELECT
					username
				FROM
					users
				WHERE
					id = articles.user_id
			) 
		FROM
			articles
		WHERE
			ID = ?
			AND article_status_id =(
				SELECT
					id
				FROM
					article_statuses
				WHERE
					status = ?
			)
		`

	var a model.ViewArticle
	if err := ar.db.QueryRowContext(ctx, query, id, constant.PUBLISHED).
		Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt, &a.Username); err != nil {
			return nil, err
	}

	return &a, nil
}

func (ar *ArticleRepository) GetArticleCount(ctx context.Context) (int, error) {
	query := 
		`
		SELECT
			count(id)
		FROM
			articles
		WHERE
			article_status_id = (
				SELECT
					id
				FROM
					article_statuses
				WHERE
					status = ?
			)
		`

	var count int
	if err := ar.db.QueryRowContext(ctx, query, constant.PUBLISHED).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (ar *ArticleRepository) GetArticleCountByUser(ctx context.Context, name string) (int, error) {
	// TODO: GetArticleCountとの共通化
	query := 
		`
		SELECT
			count(id)
		FROM
			articles
		WHERE
			user_id = (
				select
					id
				from
					users
				where
					username = ?
			)
			AND article_status_id = (
				SELECT
					id
				FROM
					article_statuses
				WHERE
					status = ?
			)
		`

	var count int
	if err := ar.db.QueryRowContext(ctx, query, name, constant.PUBLISHED).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (ar *ArticleRepository) GetByPageNumber(ctx context.Context, n int) ([]model.ViewArticle, error) {
	var rows *sql.Rows
	// NOTE: index利用のため（サブクエリを使用しないため）に直接指定
	// idが初期データの投入順に依存するため変更時は修正が必要
	artStaId := 1

	if n == 1 {
		atcQuery := `SELECT id, title, content, updated_at, (SELECT username FROM users WHERE id = articles.user_id) ` +
			`FROM articles ` +
			`WHERE article_status_id = ? ` +
			`ORDER BY updated_at DESC LIMIT ?`

		stmt, err := ar.db.PrepareContext(ctx, atcQuery)
		if err != nil {
			return nil, err
		}

		rows, err = stmt.QueryContext(ctx, artStaId, constant.ARTICLES_PER_PAGE)
		if err != nil {
			return nil, err
		}

	} else {
		// updated_atが同一のレコードは存在しない想定
		// paging indexを利用(user_id, article_status_id, updated_at DESC)
		atcQuery :=
			`SELECT id, title, content, updated_at, (SELECT username FROM users WHERE id = articles.user_id) ` +
				`FROM articles ` +
				`WHERE article_status_id = ? ` +
				`AND updated_at < ` +
				`(SELECT updated_at FROM articles WHERE 
						article_status_id = ? ` +
				`ORDER BY updated_at DESC LIMIT 1 OFFSET ?) ` +
				`ORDER BY updated_at DESC LIMIT ?`

		stmt, err := ar.db.PrepareContext(ctx, atcQuery)
		if err != nil {
			return nil, err
		}

		offset := constant.ARTICLES_PER_PAGE*(n-1) - 1

		rows, err = stmt.QueryContext(ctx, artStaId, artStaId, offset, constant.ARTICLES_PER_PAGE)
		if err != nil {
			return nil, err
		}
	}

	var articles []model.ViewArticle

	if rows.Next() {
		var a model.ViewArticle
		_ = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt, &a.Username)
		articles = append(articles, a)
	} else {
		return nil, errors.New("article not found")
	}

	for rows.Next() {
		var a model.ViewArticle
		_ = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt, &a.Username)
		articles = append(articles, a)
	}

	return articles, nil
}

func (ar *ArticleRepository) GetByUserAndPageNumber(ctx context.Context, un string, n int) ([]model.ViewArticle, error) {
	uidQuery := `SELECT id FROM users WHERE username = ?`
	stmt, err := ar.db.PrepareContext(ctx, uidQuery)
	if err != nil {
		return nil, err
	}
	var uid int
	if err := stmt.QueryRowContext(ctx, un).Scan(&uid); err != nil {
		return nil, errors.New("user not found")
	}

	var rows *sql.Rows
	// NOTE: index利用のため（サブクエリを使用しないため）に直接指定
	// idが初期データの投入順に依存するため変更時は修正が必要
	artStaId := 1

	if n == 1 {
		atcQuery := `SELECT id, title, content, updated_at FROM articles ` +
			`WHERE user_id = ? AND article_status_id = ? ` +
			`ORDER BY updated_at DESC LIMIT ?`

		stmt, err = ar.db.PrepareContext(ctx, atcQuery)
		if err != nil {
			return nil, err
		}

		rows, err = stmt.QueryContext(ctx, uid, artStaId, constant.ARTICLES_PER_PAGE)
		if err != nil {
			return nil, err
		}

	} else {
		// updated_atが同一のレコードは存在しない想定
		// paging indexを利用(user_id, article_status_id, updated_at DESC)
		atcQuery :=
			`SELECT id, title, content, updated_at FROM articles WHERE ` +
				`user_id = ? ` +
				`AND article_status_id = ? ` +
				`AND updated_at < ` +
				`(SELECT updated_at FROM articles WHERE 
						user_id = ? AND article_status_id = ? ` +
				`ORDER BY updated_at DESC LIMIT 1 OFFSET ?) ` +
				`ORDER BY updated_at DESC LIMIT ?`

		stmt, err = ar.db.PrepareContext(ctx, atcQuery)
		if err != nil {
			return nil, err
		}

		offset := constant.ARTICLES_PER_PAGE*(n-1) - 1

		rows, err = stmt.QueryContext(ctx, uid, artStaId, uid, artStaId, offset, constant.ARTICLES_PER_PAGE)
		if err != nil {
			return nil, err
		}
	}

	var articles []model.ViewArticle

	if rows.Next() {
		var a model.ViewArticle
		// FIXME: エラーハンドリング
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt)
		articles = append(articles, a)
	} else {
		return nil, errors.New("article not found")
	}

	for rows.Next() {
		var a model.ViewArticle
		// FIXME: エラーハンドリング
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.UpdatedAt)
		articles = append(articles, a)
	}

	return articles, nil
}

func (ar *ArticleRepository) Delete(ctx context.Context, aId int64, un string) (int64, error) {
	query := `UPDATE articles SET article_status_id = (SELECT id FROM article_statuses WHERE status = ?) ` +
		`WHERE id = ? AND user_id = (SELECT id FROM users WHERE username = ?)`

	stmt, err := ar.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, constant.REMOVED, aId, un)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
