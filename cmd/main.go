package main

import (
	"github.com/ArifMun/go-rest-api.git/cmd/internal/database"
	"github.com/ArifMun/go-rest-api.git/cmd/internal/delivery/rest"
	mRepo "github.com/ArifMun/go-rest-api.git/cmd/internal/repository/menu"
	rUsecase "github.com/ArifMun/go-rest-api.git/cmd/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

// connection
const (
	dbAddress = "host=localhost port=5432 user=postgres password=root dbname=go_resto_app "
)

// main function
func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	restoUscase := rUsecase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUscase)
	rest.LoadRoute(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
