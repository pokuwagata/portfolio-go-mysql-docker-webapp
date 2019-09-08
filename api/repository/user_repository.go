package repository

import (
	"api/model"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Store(ctx context.Context, u *model.User) error {
	query := `INSERT users SET username=?, password=?, status_id=(SELECT id FROM user_statuses where status= ?)`
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, u.Username, u.Password, u.Status)
	if err != nil {
		return err
	}

	return nil
}