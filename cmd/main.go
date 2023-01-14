package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// * NOTE : PLEASE RUNNING WITH : nodemon --exec go run main.go

const (
	dsn = "root:@tcp(127.0.0.1:3306)/go-restaurant?charset=utf8mb4&parseTime=True&loc=Local"
)

func getMenu(c echo.Context) error {
	menuType := c.FormValue("type")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var foodItem []MenuItem

	db.Where(MenuItem{Type: TypeMenu(menuType)}).Find(&foodItem)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodItem,
	})
}

func main() {
	seedDB()
	e := echo.New()
	e.GET("/menu/food", getMenu)

	e.Logger.Fatal(e.Start(":4000"))
}
