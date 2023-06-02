package controllers

import (
	"kpahmadyani/configs"
	"kpahmadyani/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductController(c echo.Context) error {
	var products []models.Product
	result := configs.DB.Preload("Category").Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResponese{
		Message: "succes",
		Data:    products,
	})
}

func GetCategoryWithProductsController(c echo.Context) error {
	categoryName := c.Param("category")
	var category models.Category
	configs.DB.Where("name = ?", categoryName).First(&category)
	if category.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	var products []models.Product
	configs.DB.Preload("Category").Where("category_id = ?", category.ID).Find(&products)
	return c.JSON(http.StatusOK, models.BaseResponese{
		Message: "Success",
		Data:    products,
	})

}

func AddProductController(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	var category models.Category
	configs.DB.Where("name = ?", product.Category.Name).First(&category)
	if category.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}
	product.CategoryID = category.ID
	configs.DB.Create(&product)
	return c.JSON(http.StatusOK, models.BaseResponese{
		Message: "Berhasil Menambahkan Product",
		Data:    product,
	})
}

func DeleteProductController(c echo.Context) error {
	productID := c.Param("id")
	result := configs.DB.Delete(&models.Product{}, productID)
	if result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product deleted",
	})
}
