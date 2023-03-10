package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	seedDB(db)
	return db
}
