package router

import (
	"api/controller"
	"api/external/jwtauth"
	"github.com/labstack/echo"
	"net/http"
)

func Init(
	e *echo.Echo,
	uc *controller.UserController,
	sc *controller.SessionController,
	ac *controller.ArticleController) {
	// 認証・認可が不要なルーティング
	// TODO: 削除
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/user", uc.Create)
	e.POST("/session", sc.Create)
	e.GET("/articles/:id", ac.GetById)
	e.GET("/articles", ac.GetList)

	// 認証・認可が必要なルーティング
	a := e.Group("/admin")
	jwtauth.Init(a)
	a.GET("/session", sc.GetUsernameFromToken)
	a.POST("/article/:id", ac.Update)
	a.POST("/article", ac.Create)
	a.DELETE("/article", ac.Delete)
	a.GET("/articles", ac.GetListByUser)
}
