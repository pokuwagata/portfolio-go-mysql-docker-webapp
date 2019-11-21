package usecase

import (
	"api/constant"
	"api/model"
	"context"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	r UserRepository
	e *echo.Echo
}

func NewUserUsecase(r UserRepository, e *echo.Echo) *UserUsecase {
	return &UserUsecase{r, e}
}

func (uu *UserUsecase) CreateUser(ctx context.Context, u *model.User) error {
	u.Status = constant.VALID

	if id, _ := uu.r.GetIdByUsername(ctx, u.Username); id != 0 {
		// ユーザが既に存在する場合
		err := errors.New(constant.ERR_USER_EXISTED)
		uu.e.Logger.Errorf(constant.ERR_APP_ERROR, err)
		uu.e.Logger.Debugf(constant.ERR_APP_ERROR_DEBUG, errors.WithStack(err))
		return err
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return uu.r.Store(ctx, u)
}
