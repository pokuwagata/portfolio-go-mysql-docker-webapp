package controller

import (
	"strings"
	"context"
	"net/http"
	"api/usecase"
	"api/model"
	"github.com/labstack/echo"
)

type SessionController struct {
	su *usecase.SessionUsecase
}

func NewSessionController (su *usecase.SessionUsecase) *SessionController {
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
	// tokenの認証に失敗している場合はこのAPIを呼び出せないためtokenは必ず存在するはず
	value := c.Request().Header.Get("Authorization")
	// tokenのみを抽出（e.g. Bearer xxx... → xxx...）
	rawToken := strings.Split(value, "Bearer")[1]
	// 半角空白の除去（rawTokenの先頭に半角空白が入るため）
	token := strings.Join(strings.Fields(rawToken), "")
	name, err := sc.su.GetUsernameFromToken(token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"username": name})
}