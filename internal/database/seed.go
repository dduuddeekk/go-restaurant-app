package database

import (
	"github.com/dduuddeekk/go-restaurant-app/internal/model"
	"github.com/dduuddeekk/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasigoreng",
			Price:     15000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Babi Guling",
			OrderCode: "babiguling",
			Price:     30000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Arak Bali",
			OrderCode: "arakbali",
			Price:     14000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Teh",
			OrderCode: "esteh",
			Price:     3000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
