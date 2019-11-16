package repository

import (
	"api/constant"
	"api/model"
	"context"
	"database/sql"
	"strings"
	"github.com/pkg/errors"
	"github.com/labstack/echo"
)

type UserRepository struct {
	db *sql.DB
	e *echo.Echo
}

func NewUserRepository(db *sql.DB, e *echo.Echo) *UserRepository {
	return &UserRepository{db, e}
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

func (ur *UserRepository) GetIdByUsername(ctx context.Context, name string) (int, error) {
	var id int
	query := []string{
		"SELECT",
			"id",
		"FROM users",
		"WHERE",
			"username = ?",
			"AND status_id =(",
				"SELECT",
					"id",
				"FROM user_statuses",
				"where",
					"status = ?",
			")",
	}

	rawQuery := strings.Join(query, constant.HALF_SPACE);

	if err := ur.db.QueryRowContext(ctx, rawQuery, name, constant.VALID).Scan(&id); err != nil {
		err := errors.New(constant.ERR_USER_NOT_FOUND)
		ur.e.Logger.Errorf(constant.ERR_APP_ERROR, err)
		ur.e.Logger.Debugf(constant.ERR_APP_ERROR_DEBUG, errors.WithStack(err))
		return 0, err 
	}
	return id, nil
}