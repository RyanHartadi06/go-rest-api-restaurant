package database

import (
	"rest-api-restaurant/internal/model"
	"rest-api-restaurant/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodItem := []model.MenuItem{
		{
			Name:      "Burger",
			OrderCode: "BGR",
			Price:     100,
			Type:      constant.FOOD,
		},
		{
			Name:      "Bakmi",
			OrderCode: "BGR",
			Price:     200,

			Type: constant.FOOD,
		},
	}
	drinkItem := []model.MenuItem{
		{
			Name:      "Nutrisari",
			OrderCode: "NTR",
			Price:     100,

			Type: constant.DRINK,
		},
		{
			Name:      "Soda Gembira",
			OrderCode: "SGM",
			Price:     200,

			Type: constant.DRINK,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodItem)
		db.Create(&drinkItem)
	}

}
