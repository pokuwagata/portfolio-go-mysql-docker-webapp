package router

import (
	"api/controller"
	"net/http"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo, uc *controller.UserController, sc *controller.SessionController) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/user", uc.Create)
	e.POST("/session", sc.Create)
}