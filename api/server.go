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

// hogehoge
func main() {
	e := echo.New()
	logger.Init(e)
	db := database.NewDB(e)
	defer db.Close()
	validater.Init(e)

	ur := repository.NewUserRepository(db)
	ar := repository.NewArticleRepository(db)
	pr := repository.NewPostRepository(db)
	uu := usecase.NewUserUsecase(ur)
	su := usecase.NewSessionUsecase(ur)
	au := usecase.NewArticleUsecase(ar, ur, pr, su)
	sc := controller.NewSessionController(su)
	uc := controller.NewUserController(uu)
	ac := controller.NewArticleController(au)
	router.Init(e, uc, sc, ac)
	port, _ := strconv.Atoi(os.Args[1])
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
