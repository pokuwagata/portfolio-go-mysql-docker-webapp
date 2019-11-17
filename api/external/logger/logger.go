package logger

import (
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"api/constant"
)

func Init(e *echo.Echo) {
	file, err := os.OpenFile(constant.LOG_FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: file,
	}))

	e.Logger.SetHeader(`{"time":"${time_rfc3339_nano}","level":"${level}"}`)
	e.Logger.SetOutput(file)

	mode := os.Getenv("ENV")
	switch(mode) {
		case "PRD":
			e.Logger.SetLevel(2)
			e.Logger.Info("log level = INFO")
		case "DEV":
			e.Logger.SetLevel(1)
			e.Logger.Info("log level = DEBUG")
		default:
			e.Logger.SetLevel(1)
			e.Logger.Info("log level = DEBUG")
	}

}