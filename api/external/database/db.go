package database

import (
	"database/sql"
	"context"
	"database/sql/driver"
	// "api/constant"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/shogo82148/go-sql-proxy"
)

func NewDB(e *echo.Echo) *sql.DB {
	sql.Register("mysql-proxy", proxy.NewProxyContext(&mysql.MySQLDriver{}, &proxy.HooksContext{
		Open: func(_ context.Context, _ interface{}, conn *proxy.Conn) error {
			e.Logger.Info("Open MySQL")
			return nil
		},
		Exec: func(_ context.Context, _ interface{}, stmt *proxy.Stmt, args []driver.NamedValue, result driver.Result) error {
			e.Logger.Infof("Exec: %s; args = %+v", stmt.QueryString, args)
			return nil
		},
		Query: func(_ context.Context, _ interface{}, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows) error {
			e.Logger.Infof("Query: %s; args = %+v", stmt.QueryString, args)
			return nil
		},
	}))
	// TODO: configファイル化
	db, err := sql.Open("mysql-proxy",
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
