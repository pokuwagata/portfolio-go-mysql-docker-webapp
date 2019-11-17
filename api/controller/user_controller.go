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
	u *usecase.UserUsecase
}

func NewUserController(u *usecase.UserUsecase) *UserController {
	return &UserController{u}
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

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uc.u.CreateUser(ctx, u)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, constant.SUCCESS_MESSAGE)
}
