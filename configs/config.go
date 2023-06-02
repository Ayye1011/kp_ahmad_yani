package configs

import (
	"fmt"
	"kpahmadyani/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Sitialbir@tcp(127.0.0.1:1011)/unjuk_keterampilan_ok?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	fmt.Println("Success")
}

func migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Product{})
}
