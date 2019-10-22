package usecase

import (
	repositoryMocks "api/external/mocks/repository"
	usecaseMocks "api/external/mocks/usecase"
	"api/model"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreate(t *testing.T) {
	a := &model.Article{
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	token := "token"

	urm := new(repositoryMocks.UserRepositoryMock)

	t.Run("success", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		arm.On("Store", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil)
		sum.On("GetUsernameFromToken", token).Return("username", nil)

		au := NewArticleUsecase(arm, urm, sum)
		if err := au.Create(context.TODO(), a, token); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		arm.AssertExpectations(t)
		sum.AssertExpectations(t)
	})

	t.Run("errorGetUsernameFromToken", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		mockErr := errors.New("mock error")
		sum.On("GetUsernameFromToken", token).Return("", mockErr)

		au := NewArticleUsecase(arm, urm, sum)
		if err := au.Create(context.TODO(), a, token); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was expected:", mockErr)
		}

		sum.AssertExpectations(t)
	})

	t.Run("errorStore", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		mockErr := errors.New("mock error")
		arm.On("Store", mock.Anything, mock.AnythingOfType("*model.Article")).Return(mockErr)
		sum.On("GetUsernameFromToken", token).Return("username", nil)

		au := NewArticleUsecase(arm, urm, sum)
		if err := au.Create(context.TODO(), a, token); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was expected:", mockErr)
		}

		arm.AssertExpectations(t)
		sum.AssertExpectations(t)
	})
}
