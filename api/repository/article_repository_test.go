package repository

import (
	"api/constant"
	"api/model"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"
	"github.com/labstack/echo"
	"strings"
)

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := []string {
		"INSERT",
			"articles",
		"SET",
			"title = ?,",
			"content = ?,",
			"user_id =(",
				"SELECT",
					"id",
				"FROM",
					"users",
				"WHERE",
					"username = ?",
			"),",
			"article_status_id =(",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
			")",
		}

	mock.ExpectExec(regexp.QuoteMeta(strings.Join(query, constant.HALF_SPACE))).
		WithArgs("タイトル", "コンテンツ", "テストユーザ", constant.PUBLISHED).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ar := NewArticleRepository(db, echo.New())
	article := model.Article{Title: "タイトル", Content: "コンテンツ", Username: "テストユーザ", ArticleStatus: constant.PUBLISHED}
	if err := ar.Store(context.TODO(), &article); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := []string {
		"UPDATE",
			"articles",
		"SET",
			"title = ?,",
			"content = ?,",
			"user_id =(",
				"SELECT",
					"id",
				"FROM",
					"users",
				"WHERE",
					"username = ?",
			"),",
			"article_status_id =(",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
			")",
		"WHERE",
			"id = ?",
	}

	mock.ExpectExec(regexp.QuoteMeta(strings.Join(query, constant.HALF_SPACE))).
		WithArgs("タイトル", "コンテンツ", "テストユーザ", constant.PUBLISHED, int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ar := NewArticleRepository(db, echo.New())
	article := model.Article{ID: int64(1), Title: "タイトル", Content: "コンテンツ", Username: "テストユーザ", ArticleStatus: constant.PUBLISHED}
	if err := ar.Update(context.TODO(), &article); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := []string{
		"SELECT",
			"id,",
			"title,",
			"content,",
			"updated_at,",
			"(",
				"SELECT",
					"username",
				"FROM",
					"users",
				"WHERE",
					"id = articles.user_id",
			") ",
		"FROM",
			"articles",
		"WHERE",
			"ID = ?",
			"AND article_status_id =(",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
			")",
	}

	rawQuery := strings.Join(query, constant.HALF_SPACE)

	row := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "username"}).
		AddRow(int64(1), "タイトル", "コンテンツ", time.Now(), "テストユーザ")

	mock.ExpectQuery(regexp.QuoteMeta(rawQuery)).
		WithArgs(int64(1), constant.PUBLISHED).
		WillReturnRows(row)

	ar := NewArticleRepository(db, echo.New())
	if _, err := ar.GetById(context.TODO(), int64(1)); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestArticleCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := []string{
		"SELECT",
			"count(id)",
		"FROM",
			"articles",
		"WHERE",
			"article_status_id = (",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
				")",
		}

	row := sqlmock.NewRows([]string{"count"}).AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(strings.Join(query, constant.HALF_SPACE))).
		WithArgs(constant.PUBLISHED).
		WillReturnRows(row)

	ar := NewArticleRepository(db, echo.New())
	if _, err := ar.GetArticleCount(context.TODO(), nil); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestArticleCountByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := []string{
		"SELECT",
			"count(id)",
		"FROM",
			"articles",
		"WHERE",
			"article_status_id = (",
				"SELECT",
					"id",
				"FROM",
					"article_statuses",
				"WHERE",
					"status = ?",
				")",
			"AND",
				"user_id = ?",
		}

	row := sqlmock.NewRows([]string{"count"}).AddRow(1)
	userId := "1"

	mock.ExpectQuery(regexp.QuoteMeta(strings.Join(query, constant.HALF_SPACE))).
		WithArgs(constant.PUBLISHED, userId).
		WillReturnRows(row)

	ar := NewArticleRepository(db, echo.New())
	if _, err := ar.GetArticleCount(context.TODO(), map[string]string{constant.USER_ID_COLUMN : userId}); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetByPageNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := ``

	row := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "username"}).
		AddRow(int64(1), "タイトル", "コンテンツ", time.Now(), "テストユーザ")

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1), constant.PUBLISHED).
		WillReturnRows(row)

	ar := NewArticleRepository(db, echo.New())
	if _, err := ar.GetById(context.TODO(), int64(1)); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}