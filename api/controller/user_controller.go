package controller

import (
	"api/usecase"
	"api/model"
	"context"
	apiErr "api/error"
	"net/http"
	"github.com/labstack/echo"
)

type UserController struct {
	u *usecase.UserUsecase
}

func NewUserController(u *usecase.UserUsecase) *UserController {
	return &UserController{u}
}

func (uc *UserController) Create(c echo.Context) error {
	u := &model.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, apiErr.Response{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, apiErr.Response{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uc.u.CreateUser(ctx, u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apiErr.Response{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}