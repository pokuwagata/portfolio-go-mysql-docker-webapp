package controller

import (
	"api/model"
	"api/usecase"
	"context"
	"github.com/labstack/echo"
	"net/http"
)

type SessionController struct {
	su *usecase.SessionUsecase
}

func NewSessionController(su *usecase.SessionUsecase) *SessionController {
	return &SessionController{su}
}

func (sc *SessionController) Create(c echo.Context) error {
	s := &model.Session{}

	if err := c.Bind(s); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err := c.Validate(s); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	t, err := sc.su.CreateSession(ctx, s)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"token": t, "username": s.Username})
}

func (sc *SessionController) GetUsernameFromToken(c echo.Context) error {
	token, err := GetTokenFromHeader(c)
	if err != nil {
		return err
	}

	name, err := sc.su.GetUsernameFromToken(token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"username": name})
}
