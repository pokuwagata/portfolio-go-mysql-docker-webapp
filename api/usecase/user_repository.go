package usecase

import (
	"api/model"
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, u *model.User) error
	GetPassword(ctx context.Context, s *model.Session) (string, error)
	GetIdByUsername(ctx context.Context, name string) (int, error)
	Exists(ctx context.Context, name string) (bool, error)
}
