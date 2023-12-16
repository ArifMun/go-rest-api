package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmi",
			OrderCode: "bakmie",
			Price:     40000,
		},
		{
			Name:      "Mi Ayam",
			OrderCode: "MA",
			Price:     20000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinksMenu(c echo.Context) error {
	drinksMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "ET",
			Price:     2000,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "EJ",
			Price:     2000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinksMenu,
	})
}

func main() {
	e := echo.New()

	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drinks", getDrinksMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
