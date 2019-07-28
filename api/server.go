package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	log.Print(writer);

	// var datasource string
	// env := os.Getenv("env")
	// switch env {
	// case "PROD":
	// 	datasource = os.Getenv("CLEARDB_USER") +
	// 		":" +
	// 		os.Getenv("CLEARDB_PASSWORD") +
	// 		"@tcp(" +
	// 		os.Getenv("CLEARDB_HOST") +
	// 		":3306)/" +
	// 		os.Getenv("CLEARDB_DBNAME")
	// case "DEV":
	// 	datasource = "user:password@(db:3306)/sample_db"
	// default:
	// 	return
	// }

	// db, err := sql.Open("mysql", datasource)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Fprintf(writer, "%v", db)

	// err = db.Ping()
	// if err != nil {
	// 	// do something here
	// 	log.Fatal(err)
	// }
	// defer db.Close()
}

type loggingResponseWriter struct {
	status int
	body   string
	http.ResponseWriter
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *loggingResponseWriter) Write(body []byte) (int, error) {
	w.body = string(body)
	return w.ResponseWriter.Write(body)
}

func responseLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loggingRW := &loggingResponseWriter{
					ResponseWriter: w,
			}
			h.ServeHTTP(loggingRW, r)
			log.Println("Status : ", loggingRW.status, "Response : ", loggingRW.body)
	})
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	openLogFile("logs/system.log");
	fmt.Printf("Starting server at Port %d", port)
	// http.HandleFunc("/", handler)
	http.Handle("/", responseLogger(http.HandlerFunc(handler)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil));
}
