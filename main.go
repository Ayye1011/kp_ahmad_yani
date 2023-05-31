package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id       uint8  `gorm:primaryKey AutoIncrement json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	connectDatabase()
	e := echo.New()
	e.GET("/username", GetUsernameController)
	e.Start(":8000")
}

func GetUsernameController(c echo.Context) error {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(500, nil)
	}
	return c.JSON(200, users)
}

func connectDatabase() {
	dsn := "root:Sitialbir@tcp(127.0.0.1:1011)/unjuk_keterampilan?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	fmt.Println("Success")
}
