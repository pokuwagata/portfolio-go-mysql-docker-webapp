package main

import (
	"api/external/validater"
	"api/external/database"
	"api/repository"
	"api/usecase"
	"api/controller"
	"api/external/logger"
	"api/external/router"
	"os"
	"strconv"
	"fmt"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	logger.Init(e)
	db := database.NewDB(e)
	defer func() {
		e.Logger.Info("Close MySQL")
		db.Close()
	}()
	validater.Init(e)

	ur := repository.NewUserRepository(db, e)
	ar := repository.NewArticleRepository(db, e)
	uu := usecase.NewUserUsecase(ur, e)
	su := usecase.NewSessionUsecase(ur, e)
	au := usecase.NewArticleUsecase(ar, ur, su)
	sc := controller.NewSessionController(su)
	uc := controller.NewUserController(uu)
	ac := controller.NewArticleController(au)
	router.Init(e, uc, sc, ac)
	port, _ := strconv.Atoi(os.Args[1])
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
