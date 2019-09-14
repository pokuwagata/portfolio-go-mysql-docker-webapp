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
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	su := usecase.NewSessionUsecase(ur)
	sc := controller.NewSessionController(su)
	router.Init(e, uc, sc)
	port, _ := strconv.Atoi(os.Args[1])
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
