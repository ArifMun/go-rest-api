package database

import (
	"github.com/ArifMun/go-rest-api.git/cmd/internal/model"
	"github.com/ArifMun/go-rest-api.git/cmd/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmi",
			OrderCode: "bakmie",
			Price:     40000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Mi Ayam",
			OrderCode: "MA",
			Price:     20000,
			Type:      constant.MenuTypeFood,
		},
	}
	drinksMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "ET",
			Price:     2000,
			Type:      constant.MenuTypeDrinks,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "EJ",
			Price:     2000,
			Type:      constant.MenuTypeDrinks,
		},
	}

	// db.AutoMigrate(&model.MenuItem{})
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
