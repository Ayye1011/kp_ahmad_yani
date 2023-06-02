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
	User     string `json:"user"`
	Password string `json:"password"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique" json:"category"`
}

type Product struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}

type BaseResponese struct {
	Message string      `json:"mesage"`
	Data    interface{} `json:"data`
}

func main() {
	connectDatabase()

	e := echo.New()
	// controller user
	e.GET("/user", GetUserController)
	e.GET("/user/:id", GetDetailUserController)
	e.POST("/user", LoginRequest)
	e.DELETE("/user/:id", DeleteUser)
	e.PUT("/user/:id", UpdateUser)

	//controller product
	e.GET("/products", GetProductController)
	e.GET("/products/:category", GetCategoryWithProductsController)
	e.POST("/products", AddProductController)
	e.DELETE("products/:id", DeleteProductController)

	e.Start(":8000")
}

func GetUserController(c echo.Context) error {
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

func GetDetailUserController(c echo.Context) error {
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

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	updatedUser := new(User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	user.User = updatedUser.User
	user.Password = updatedUser.Password

	DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated",
		"data":    user,
	})
}

func GetProductController(c echo.Context) error {
	var products []Product
	result := DB.Preload("Category").Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, BaseResponese{
		Message: "succes",
		Data:    products,
	})
}

func GetCategoryWithProductsController(c echo.Context) error {
	categoryName := c.Param("category")
	var category Category
	DB.Where("name = ?", categoryName).First(&category)
	if category.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	var products []Product
	DB.Preload("Category").Where("category_id = ?", category.ID).Find(&products)
	return c.JSON(http.StatusOK, BaseResponese{
		Message: "Success",
		Data:    products,
	})

}

func AddProductController(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	var category Category
	DB.Where("name = ?", product.Category.Name).First(&category)
	if category.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}
	product.CategoryID = category.ID
	DB.Create(&product)
	return c.JSON(http.StatusOK, BaseResponese{
		Message: "Berhasil Menambahkan Product",
		Data:    product,
	})
}

func DeleteProductController(c echo.Context) error {
	productID := c.Param("id")
	result := DB.Delete(&Product{}, productID)
	if result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product deleted",
	})
}

func connectDatabase() {
	dsn := "root:Sitialbir@tcp(127.0.0.1:1011)/unjuk_keterampilan_ok?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	fmt.Println("Success")
}

func migration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Product{})
}
