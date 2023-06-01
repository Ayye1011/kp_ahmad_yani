package main

import (
	"fmt"
	"net/http"

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

type BaseResponese struct {
	Message string
	Data    interface{}
}

func main() {
	connectDatabase()
	e := echo.New()
	e.GET("/username", GetUsernameController)
	e.GET("/username/:id", GetDetailUsernameController)
	e.POST("/username", LoginRequest)
	e.DELETE("/username/:id", DeleteUser)
	e.Start(":8000")
}

func GetUsernameController(c echo.Context) error {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(500, nil)
	}
	return c.JSON(200, BaseResponese{
		Message: "Success",
		Data:    users,
	})
}

func GetDetailUsernameController(c echo.Context) error {
	id := c.Param("id")
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(200, BaseResponese{
		Message: "Success",
		Data:    user,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	result := DB.Delete(&User{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}

func LoginRequest(c echo.Context) error {
	var userInput User
	c.Bind(&userInput)
	result := DB.Create(&userInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, BaseResponese{
		Message: "Success",
		Data:    userInput,
	})
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

func migration() {
	DB.AutoMigrate(&User{})
}
