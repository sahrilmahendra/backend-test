package routes

import (
	"erajaya/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// function routes
func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// route products without auth
	e.POST("/products", controllers.AddProductController)
	e.GET("/products", controllers.GetAllProductsController)
	e.GET("/products/newest", controllers.GetAllProductsByCreatedDescController)
	e.GET("/products/name/asc", controllers.GetAllProductsByNameAscController)
	e.GET("/products/name/desc", controllers.GetAllProductsByNameDescController)
	e.GET("/products/price/lowest", controllers.GetAllProductsByPriceAscController)
	e.GET("/products/price/highest", controllers.GetAllProductsByPriceDescController)

	return e
}
