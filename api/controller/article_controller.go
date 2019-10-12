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

func (ac *ArticleController) Update(c echo.Context) error {
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
	err := ac.au.Update(ctx, a, t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "success")
}

func (ac *ArticleController) Get(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	article, err := ac.au.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
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
	} else {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "リクエストパラメータが不正です"})
	}
}

func (ac *ArticleController) GetListByUser(c echo.Context) error {
	n, _ := strconv.Atoi(c.QueryParam("number"))
	if n > 1 {
		articles, err := ac.getByUserAndPageNumber(c, n)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, model.GetListResponse{Articles: articles})
	} else if n == 1 {
		max, err := ac.getMaxPageNumberByUser(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		}

		articles, err := ac.getByUserAndPageNumber(c, n)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, model.FirstGetListResponse{Articles: articles, Max: max})
	} else {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "リクエストパラメータが不正です"})
	}
}

func (ac *ArticleController) getByPageNumber(c echo.Context, n int) ([]model.ViewArticle, error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	articles, err := ac.au.GetByPageNumber(ctx, n)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ac *ArticleController) getByUserAndPageNumber(c echo.Context, n int) ([]model.ViewArticle, error) {
	t := GetTokenFromHeader(c)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	articles, err := ac.au.GetByUserAndPageNumber(ctx, n, t)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ac *ArticleController) getMaxPageNumber(c echo.Context) (int, error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	max, err := ac.au.GetMaxPageNumber(ctx)
	if err != nil {
		return 0, err
	}

	return max, nil
}

func (ac *ArticleController) getMaxPageNumberByUser(c echo.Context) (int, error) {
	t := GetTokenFromHeader(c)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	max, err := ac.au.GetMaxPageNumberByUser(ctx, t)
	if err != nil {
		return 0, err
	}

	return max, nil
}

func (ac *ArticleController) Delete(c echo.Context) error {
	// TODO: pathParameterに変更
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	token := GetTokenFromHeader(c)
	_, err = ac.au.Delete(ctx, id, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
