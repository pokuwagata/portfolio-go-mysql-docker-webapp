package repository

import (
	"api/constant"
	"api/model"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query :=
		`INSERT
			articles
		SET
			title = ?,
			content = ?,
			user_id =(
				SELECT
					id
				FROM
					users
				WHERE
					username = ?
			),
			article_status_id =(
				SELECT
					id
				FROM
					article_statuses
				WHERE
					status = ?
			)`

	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs("タイトル", "コンテンツ", "テストユーザ", constant.PUBLISHED).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ar := NewArticleRepository(db)
	article := model.Article{Title: "タイトル", Content: "コンテンツ", Username: "テストユーザ", ArticleStatus: constant.PUBLISHED}
	if err = ar.Store(context.TODO(), &article); err != nil {
		t.Errorf("error was not expected : %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
