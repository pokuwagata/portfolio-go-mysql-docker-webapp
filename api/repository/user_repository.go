package repository

import (
	"errors"
	"api/constant"
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

func (ur *UserRepository) GetPassword(ctx context.Context, s *model.Session) (string, error) {
	var p string
	query := `SELECT password FROM users WHERE username=? AND status_id=(SELECT id FROM user_statuses where status= ?)`
	// usernameはユニークキー
	if err := ur.db.QueryRowContext(ctx, query, s.Username, constant.VALID).Scan(&p); err != nil {
		return "", errors.New("user not found")
	}
	return p, nil
}