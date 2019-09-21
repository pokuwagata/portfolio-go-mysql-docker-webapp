package controller

import (
	"strings"
	"github.com/labstack/echo"
)

func GetTokenFromHeader(c echo.Context) string {
	// tokenの認証に失敗している場合はこのAPIを呼び出せないためtokenは必ず存在するはず
	value := c.Request().Header.Get("Authorization")
	// tokenのみを抽出（e.g. Bearer xxx... → xxx...）
	rawToken := strings.Split(value, "Bearer")[1]
	// 半角空白の除去（rawTokenの先頭に半角空白が入るため）
	token := strings.Join(strings.Fields(rawToken), "")
	return token
}