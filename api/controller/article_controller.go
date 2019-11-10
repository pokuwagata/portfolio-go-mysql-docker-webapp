package controller

import (
	"strings"
	"api/model"
	"api/usecase"
	"context"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"api/constant"
)

type ArticleController struct {
	au usecase.ArticleUsecaseInterface
}

func NewArticleController(au usecase.ArticleUsecaseInterface) *ArticleController {
	return &ArticleController{au}
}

func (ac *ArticleController) Create(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "ArticleController.Create"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "ArticleController.Create"}, constant.LOG_SEPARATOR))

	a := &model.Article{}

	if err := c.Bind(a); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(a); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	token, err := GetTokenFromHeader(c)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := ac.au.Create(ctx, a, token); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, constant.SUCCESS_MESSAGE)
}

func (ac *ArticleController) Update(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "ArticleController.Update"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "ArticleController.Update"}, constant.LOG_SEPARATOR))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	a := &model.Article{}
	a.ID = id

	if err := c.Bind(a); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(a); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	token, err := GetTokenFromHeader(c)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := ac.au.Update(ctx, a, token); err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, constant.SUCCESS_MESSAGE)
}

func (ac *ArticleController) GetById(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "ArticleController.GetById"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "ArticleController.GetById"}, constant.LOG_SEPARATOR))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	article, err := ac.au.GetById(ctx, id)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, article)
}

func (ac *ArticleController) GetList(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "ArticleController.GetList"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "ArticleController.GetList"}, constant.LOG_SEPARATOR))

	n, _ := strconv.Atoi(c.QueryParam("number"))
	if n > 1 {
		articles, err := ac.getByPageNumber(c, n)
		if err != nil {
			return GetErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, model.GetListResponse{Articles: articles})
	} else if n == 1 {
		max, err := ac.getMaxPageNumber(c)
		if err != nil {
			return GetErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		articles, err := ac.getByPageNumber(c, n)
		if err != nil {
			return GetErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, model.FirstGetListResponse{Articles: articles, Max: max})
	} else {
		return GetErrorResponse(c, http.StatusBadRequest, constant.ERR_INVALID_REQUEST_PARAM)
	}
}

func (ac *ArticleController) GetListByUser(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "ArticleController.GetListByUser"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "ArticleController.GetListByUser"}, constant.LOG_SEPARATOR))

	n, _ := strconv.Atoi(c.QueryParam("number"))
	if n > 1 {
		articles, err := ac.getByUserAndPageNumber(c, n)
		if err != nil {
			return GetErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, model.GetListResponse{Articles: articles})
	} else if n == 1 {
		max, err := ac.getMaxPageNumberByUser(c)
		if err != nil {
			return GetErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		articles, err := ac.getByUserAndPageNumber(c, n)
		if err != nil {
			return GetErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, model.FirstGetListResponse{Articles: articles, Max: max})
	} else {
		// TODO: エラーログ
		return GetErrorResponse(c, http.StatusBadRequest, constant.ERR_INVALID_REQUEST_PARAM)
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
	token, err := GetTokenFromHeader(c)
	if err != nil {
		return nil, err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	articles, err := ac.au.GetByUserAndPageNumber(ctx, n, token)
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
	token, err := GetTokenFromHeader(c)
	if err != nil {
		return 0, err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	max, err := ac.au.GetMaxPageNumberByUser(ctx, token)
	if err != nil {
		return 0, err
	}

	return max, nil
}

func (ac *ArticleController) Delete(c echo.Context) error {
	c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_BEGIN, "ArticleController.Delete"}, constant.LOG_SEPARATOR))
	defer c.Logger().Info(strings.Join([]string{constant.LOG_METHOD_END, "ArticleController.Delete"}, constant.LOG_SEPARATOR))

	// TODO: pathParameterに変更
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 64)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	token, err := GetTokenFromHeader(c)

	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// TODO: idを返す必要ある？
	_, err = ac.au.Delete(ctx, id, token)
	if err != nil {
		return GetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, constant.SUCCESS_MESSAGE)
}
