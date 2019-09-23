package controller

import (
	"strconv"
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

func (ac *ArticleController) GetList(c echo.Context) error {
	n, _ := strconv.Atoi(c.QueryParam("number"))
	c.Logger().Debug(n)
	if n > 0 {
		return ac.getByPageNumber(c, n)
	} else if n == 0 {
		return ac.getAll(c)
	} else {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "リクエストパラメータが不正です"})
	}
}

func (ac *ArticleController) getAll(c echo.Context) error {
	return nil
}

func (ac *ArticleController) getByPageNumber(c echo.Context, n int) error {
	t := GetTokenFromHeader(c)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	articles, err := ac.au.GetByPageNumber(ctx, n, t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, articles)
}