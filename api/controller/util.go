package controller

import (
	"errors"
	"github.com/labstack/echo"
	"strings"
	"api/model"
	"api/constant"
)

func GetTokenFromHeader(c echo.Context) (string, error) {
	// tokenの認証に失敗している場合はこのAPIを呼び出せないためtokenは必ず存在するはず
	value := c.Request().Header.Get(echo.HeaderAuthorization)
	if value == "" {
		err := errors.New(constant.ERR_TOKEN_NOT_FOUND)
		c.Logger().Errorf(constant.ERR_APP_ERROR, err)
		c.Logger().Debugf(constant.ERR_APP_ERROR_DEBUG, err)
		return "", err
	}
	// tokenのみを抽出（e.g. Bearer xxx... → xxx...）
	rawToken := strings.Split(value, "Bearer")
	if len(rawToken) == 1 {
		err := errors.New(constant.ERR_INVALID_TOKEN)
		c.Logger().Errorf(constant.ERR_APP_ERROR, err)
		c.Logger().Debugf(constant.ERR_APP_ERROR_DEBUG, err)
		return "", err
	}
	// 半角空白の除去（rawTokenの先頭に半角空白が入るため）
	token := strings.Join(strings.Fields(rawToken[1]), "")
	return token, nil
}

func GetErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, model.ErrorResponse{Code: status, Message: message})
}
