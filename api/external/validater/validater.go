package validater

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Init(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}
