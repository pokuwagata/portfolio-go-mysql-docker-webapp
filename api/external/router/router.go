package router

import (
	"api/controller"
	"api/external/jwtauth"
	"net/http"
	"github.com/labstack/echo"
)

func Init(
	e *echo.Echo,
	uc *controller.UserController,
	sc *controller.SessionController,
	ac *controller.ArticleController) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/user", uc.Create)
	e.POST("/session", sc.Create)
	// 認証・認可が必要なリクエスト
	a := e.Group("/admin")
	jwtauth.Init(a)
	a.DELETE("/user", uc.Delete)
	a.GET("/session", sc.GetUsernameFromToken)
	a.POST("/article", ac.Create)
	a.DELETE("/article", ac.Delete)
	a.GET("/articles", ac.GetList)
}