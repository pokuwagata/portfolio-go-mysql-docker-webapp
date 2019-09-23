package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func NewDB(e *echo.Echo) *sql.DB {
	// TODO: configファイル化
	db, err := sql.Open("mysql",
		"user:password@tcp(db:3306)/sample_db?parseTime=true")
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	for {
		err = db.Ping()
		if err == nil {
			break
		} else {
			e.Logger.Info(err.Error())
		}
	}

	return db
}
