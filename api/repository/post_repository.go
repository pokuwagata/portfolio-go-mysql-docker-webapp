package repository

import (
	"api/model"
	"context"
	"database/sql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func(ar *PostRepository) Store(ctx context.Context, p *model.Post) (int64, error) {
	query := `INSERT posts SET user_id=(SELECT id FROM users where username= ?)`
	stmt, err := ar.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, p.Username)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}