package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuType string

const (
	MenuTypeFood   = "food"
	MenuTypeDrinks = "drinks"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

// connection
const (
	dbAddress = "host=localhost port=5432 user=postgres password=root dbname=go_resto_app "
)

// DB seed
func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmi",
			OrderCode: "bakmie",
			Price:     40000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Mi Ayam",
			OrderCode: "MA",
			Price:     20000,
			Type:      MenuTypeFood,
		},
	}
	drinksMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "ET",
			Price:     2000,
			Type:      MenuTypeDrinks,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "EJ",
			Price:     2000,
			Type:      MenuTypeDrinks,
		},
	}

	fmt.Println(foodMenu, drinksMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&MenuItem{})
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

// main function
func main() {
	seedDB()
	e := echo.New()

	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
