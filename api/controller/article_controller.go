package controller

import (
	"api/model"
	"api/usecase"
	"context"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type ArticleController struct {
	au *usecase.ArticleUsecase
}

func NewArticleController(au *usecase.ArticleUsecase) *ArticleController {
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
	if n > 1 {
		articles, err := ac.getByPageNumber(c, n)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, model.GetListResponse{Articles: articles})
	} else if n == 1 {
		max, err := ac.getMaxPageNumber(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		}

		articles, err := ac.getByPageNumber(c, n)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, model.FirstGetListResponse{Articles: articles, Max: max})
	} else if n == 0 {
		return ac.getAll(c)
	} else {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "リクエストパラメータが不正です"})
	}
}

func (ac *ArticleController) getAll(c echo.Context) error {
	return nil
}

func (ac *ArticleController) getByPageNumber(c echo.Context, n int) ([]model.ViewArticle, error) {
	t := GetTokenFromHeader(c)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	articles, err := ac.au.GetByPageNumber(ctx, n, t)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ac *ArticleController) getMaxPageNumber(c echo.Context) (int, error) {
	t := GetTokenFromHeader(c)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	max, err := ac.au.GetMaxPageNumber(ctx, t)
	if err != nil {
		return 0, err
	}

	return max, nil
}
