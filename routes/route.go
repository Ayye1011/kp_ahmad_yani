package routes

import (
	"kpahmadyani/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	// controller user
	e.GET("/user", controllers.GetUserController)
	e.GET("/user/:id", controllers.GetDetailUserController)
	e.POST("/user", controllers.AddUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)

	//controller product
	e.GET("/products", controllers.GetProductController)
	e.GET("/products/:category", controllers.GetCategoryWithProductsController)
	e.POST("/products", controllers.AddProductController)
	e.DELETE("/products/:id", controllers.DeleteProductController)
	return e

}
