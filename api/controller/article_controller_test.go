package controller

import (
	usecaseMocks "api/external/mocks/usecase"
	"api/external/validater"
	"api/model"
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
