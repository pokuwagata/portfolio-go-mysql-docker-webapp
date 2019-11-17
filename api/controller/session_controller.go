package controller

import (
	"api/model"
	"api/usecase"
	"context"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"api/constant"
)

type SessionController struct {
	su *usecase.SessionUsecase
}

func NewSessionController(su *usecase.SessionUsecase) *SessionController {
	return &SessionController{su}
}

func (sc *SessionController) Create(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "SessionController.Create"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "SessionController.Create"}, constant.LOG_SEPARATOR))

	s := &model.Session{}

	if err := c.Bind(s); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(s); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	t, err := sc.su.CreateSession(ctx, s)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"token": t, "username": s.Username})
}

func (sc *SessionController) GetUsernameFromToken(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "SessionController.GetUsernameFromToken"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "SessionController.GetUsernameFromToken"}, constant.LOG_SEPARATOR))

	token, err := GetTokenFromHeader(c)
	if err != nil {
		return err
	}

	name, err := sc.su.GetUsernameFromToken(token)

	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"username": name})
}
