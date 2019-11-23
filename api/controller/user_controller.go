package controller

import (
	"api/constant"
	"api/model"
	"api/usecase"
	"context"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type UserController struct {
	uu *usecase.UserUsecase
	su *usecase.SessionUsecase
}

func NewUserController(uu *usecase.UserUsecase, su *usecase.SessionUsecase) *UserController {
	return &UserController{uu, su}
}

func (uc *UserController) Create(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "UserController.Create"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "UserController.Create"}, constant.LOG_SEPARATOR))
	u := &model.User{}

	if err := c.Bind(u); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	s := &model.Session{Username: u.Username, Password: u.Password}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uc.uu.CreateUser(ctx, u)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := uc.su.CreateSession(ctx, s)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"token": token, "username": s.Username})
}
