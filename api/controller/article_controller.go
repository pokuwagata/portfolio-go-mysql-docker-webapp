package controller

import (
	"context"
	"api/model"
	"net/http"
	"github.com/labstack/echo"
	"api/usecase"
)

type ArticleController struct {
	au *usecase.ArticleUsecase
}

func NewArticleController (au *usecase.ArticleUsecase) *ArticleController {
	return &ArticleController{au}
}

func (ac *ArticleController) Create(c echo.Context) error {
	a := &model.Article{}

	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err := c.Validate(a); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	t := GetTokenFromHeader(c)
	err := ac.au.CreateArticle(ctx, a, t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}