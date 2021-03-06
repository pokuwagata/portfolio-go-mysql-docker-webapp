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

func TestUpdate(t *testing.T) {
	a := &model.Article{
		ID: int64(1),
		Title:   "タイトル",
		Content: "コンテンツ",
	}
	token := "token"

	urm := new(repositoryMocks.UserRepositoryMock)

	t.Run("success", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		arm.On("Update", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil)
		sum.On("GetUsernameFromToken", token).Return("username", nil)

		au := NewArticleUsecase(arm, urm, sum)
		if err := au.Update(context.TODO(), a, token); err != nil {
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
		if err := au.Update(context.TODO(), a, token); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was expected:", mockErr)
		}

		sum.AssertExpectations(t)
	})

	t.Run("errorUpdate", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		mockErr := errors.New("mock error")
		arm.On("Update", mock.Anything, mock.AnythingOfType("*model.Article")).Return(mockErr)
		sum.On("GetUsernameFromToken", token).Return("username", nil)

		au := NewArticleUsecase(arm, urm, sum)
		if err := au.Update(context.TODO(), a, token); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was expected:", mockErr)
		}

		arm.AssertExpectations(t)
		sum.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	a := &model.Article{
		ID: int64(1),
		Title:   "タイトル",
		Content: "コンテンツ",
	}

	urm := new(repositoryMocks.UserRepositoryMock)

	t.Run("success", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		arm.On("GetById", mock.Anything, a.ID).Return(nil, nil)

		au := NewArticleUsecase(arm, urm, sum)
		if _, err := au.GetById(context.TODO(), a.ID); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		arm.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)

		mockErr := errors.New("mock error")
		arm.On("GetById", mock.Anything, a.ID).Return(nil, mockErr)

		au := NewArticleUsecase(arm, urm, sum)
		if _, err := au.GetById(context.TODO(), a.ID); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was expected:", mockErr)
		}

		arm.AssertExpectations(t)
	})
}

func TestGetMaxPageNumber(t *testing.T) {
	urm := new(repositoryMocks.UserRepositoryMock)

	t.Run("success", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		arm.On("GetArticleCount", mock.Anything).Return(1, nil)

		au := NewArticleUsecase(arm, urm, sum)
		if _, err := au.GetMaxPageNumber(context.TODO()); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		arm.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)

		mockErr := errors.New("mock error")
		arm.On("GetArticleCount", mock.Anything).Return(0, mockErr)

		au := NewArticleUsecase(arm, urm, sum)
		if _, err := au.GetMaxPageNumber(context.TODO()); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was expected:", mockErr)
		}

		arm.AssertExpectations(t)
	})
}

func TestGetMaxPageNumberByUser(t *testing.T) {
	urm := new(repositoryMocks.UserRepositoryMock)
	token := "token"
	name := "username"

	t.Run("success", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		sum.On("GetUsernameFromToken", token).Return(name, nil)
		arm.On("GetArticleCountByUser", mock.Anything, name).Return(1, nil)

		au := NewArticleUsecase(arm, urm, sum)
		if _, err := au.GetMaxPageNumberByUser(context.TODO(), token); err != nil {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		arm.AssertExpectations(t)
		sum.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		arm := new(repositoryMocks.ArticleRepositoryMock)
		sum := new(usecaseMocks.SessionUsecaseMock)
		mockErr := errors.New("mock error")

		sum.On("GetUsernameFromToken", token).Return(name, nil)
		arm.On("GetArticleCountByUser", mock.Anything, name).Return(0, mockErr)

		au := NewArticleUsecase(arm, urm, sum)
		if _, err := au.GetMaxPageNumberByUser(context.TODO(), token); err.Error() != mockErr.Error() {
			t.Fatalf("an error '%s' was not expected:", err)
		}

		arm.AssertExpectations(t)
		sum.AssertExpectations(t)
	})
}