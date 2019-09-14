package usecase

import (
	"api/constant"
	"api/model"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Store(ctx context.Context, u *model.User) error
}

type UserUsecase struct {
	r UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{r}
}

func (uu *UserUsecase) CreateUser(ctx context.Context, u *model.User) error {
	u.Status = constant.VALID

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return uu.r.Store(ctx, u)
}
