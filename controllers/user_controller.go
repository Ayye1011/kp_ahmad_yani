package controllers

import (
	"kpahmadyani/configs"
	"kpahmadyani/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []models.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(500, nil)
	}
	return c.JSON(200, models.BaseResponese{
		Message: "Success",
		Data:    users,
	})
}

func GetDetailUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(200, models.BaseResponese{
		Message: "Success",
		Data:    user,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	result := configs.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}

func LoginRequest(c echo.Context) error {
	var userInput models.User
	c.Bind(&userInput)
	result := configs.DB.Create(&userInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResponese{
		Message: "Success",
		Data:    userInput,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	updatedUser := new(models.User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	user.User = updatedUser.User
	user.Password = updatedUser.Password

	configs.DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated",
		"data":    user,
	})
}
