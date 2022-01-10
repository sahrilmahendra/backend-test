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
	e.GET("/products/newest", controllers.GetAllProductsController)
	e.GET("/products/name/asc", controllers.GetAllProductsByNameAscController)

	return e
}
