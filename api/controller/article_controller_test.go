package controller

import (
	usecaseMocks "api/external/mocks/usecase"
	"api/external/validater"
	"api/model"
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"net/url"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"errors"
	"api/constant"
)

func TestCreate(t *testing.T) {
	a := &model.Article{
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	j, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("an error '%s' was not expected:", err)
	}

	token := "token"

	e := echo.New()
	validater.Init(e)

	req, _ := http.NewRequest(echo.POST, "/article", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, strings.Join([]string{"Bearer", token}, " "))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	au := new(usecaseMocks.ArticleUsecaseMock)
	au.On("Create", context.TODO(), a, token).Return(nil)
	ac := NewArticleController(au)

	if err := ac.Create(c); err != nil {
		t.Fatalf("an error '%s' was not expected:", err)
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("status code '%d' was not expected:", rec.Code)
	}
}

func TestUpdate(t *testing.T) {
	a := &model.Article{
		ID:      int64(1),
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	j, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("an error '%s' was not expected:", err)
	}

	token := "token"

	e := echo.New()
	validater.Init(e)

	req, _ := http.NewRequest(echo.POST, "/article/1", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, strings.Join([]string{"Bearer", token}, " "))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	au := new(usecaseMocks.ArticleUsecaseMock)
	au.On("Update", context.TODO(), a, token).Return(nil)
	ac := NewArticleController(au)

	if err := ac.Update(c); err != nil {
		t.Fatalf("an error '%s' was not expected:", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status code '%d' was not expected:", rec.Code)
	}
}

func TestGetById(t *testing.T) {
	a := &model.Article{
		ID:      int64(1),
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	j, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("an error '%s' was not expected:", err)
	}
	e := echo.New()
	validater.Init(e)

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest(echo.GET, "/article/1", strings.NewReader(string(j)))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		au := new(usecaseMocks.ArticleUsecaseMock)
		au.On("GetById", context.TODO(), a.ID).Return(a, nil)
		ac := NewArticleController(au)

		if err := ac.GetById(c); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("status code '%d' was not expected:", rec.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		req, _ := http.NewRequest(echo.GET, "/article/0", strings.NewReader(string(j)))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:id")
		c.SetParamNames("id")
		c.SetParamValues("0")

		au := new(usecaseMocks.ArticleUsecaseMock)
		mockErr := errors.New("mock error")
		au.On("GetById", context.TODO(), int64(0)).Return(nil, mockErr)
		ac := NewArticleController(au)

		if err := ac.GetById(c); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status code '%d' was not expected:", rec.Code)
		}

		var res model.ErrorResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			if res.Message != mockErr.Error() {
				t.Fatalf("error message %s' was not expected:", res.Message)
			}
		}
	})
}

func TestGetList(t *testing.T) {
	a1 := model.ViewArticle{
		ID:      int64(1),
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	a2 := model.ViewArticle{
		ID:      int64(2),
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	articles := []model.ViewArticle{a1, a2}

	e := echo.New()
	validater.Init(e)

	t.Run("number=1", func(t *testing.T) {
		q := make(url.Values)
		q.Set("number", "1")

		req, _ := http.NewRequest(echo.GET, "/articles?" + q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		au := new(usecaseMocks.ArticleUsecaseMock)
		au.On("GetMaxPageNumber", context.TODO()).Return(2, nil)
		au.On("GetByPageNumber", context.TODO()).Return(articles, nil)
		ac := NewArticleController(au)

		if err := ac.GetList(c); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("status code '%d' was not expected:", rec.Code)
		}
	})

	t.Run("number>1", func(t *testing.T) {
		q := make(url.Values)
		q.Set("number", "2")

		req, _ := http.NewRequest(echo.GET, "/articles?" + q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		au := new(usecaseMocks.ArticleUsecaseMock)
		au.On("GetByPageNumber", context.TODO()).Return(articles, nil)
		ac := NewArticleController(au)

		if err := ac.GetList(c); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("status code '%d' was not expected:", rec.Code)
		}
	})

	t.Run("error", func(t *testing.T) {
		q := make(url.Values)
		q.Set("number", "0")

		req, _ := http.NewRequest(echo.GET, "/articles?" + q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		au := new(usecaseMocks.ArticleUsecaseMock)
		ac := NewArticleController(au)

		if err := ac.GetList(c); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status code '%d' was not expected:", rec.Code)
		}

		var res model.ErrorResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			if res.Message != constant.ERR_INVALID_REQUEST_PARAM {
				t.Fatalf("error message %s' was not expected:", res.Message)
			}
		}
	})
}
