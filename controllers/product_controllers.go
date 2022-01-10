package controllers

import (
	"erajaya/lib/databases"
	"erajaya/models"
	"erajaya/responses"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// controller for add product
func AddProductController(c echo.Context) error {
	new_product := models.Product{}
	c.Bind(&new_product)

	var err error

	v := validator.New()
	err = v.Var(new_product.Product_Name, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Product's name can't be empty"))
	}

	err = v.Var(new_product.Product_Price, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Product's price can't be empty"))
	}
	err = v.Var(new_product.Product_Price, "gt=0")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Product's price must be greater than 0"))
	}

	err = v.Var(new_product.Product_Description, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Product's description can't be empty"))
	}

	err = v.Var(new_product.Product_Qty, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Product's quantity can't be empty"))
	}
	err = v.Var(new_product.Product_Qty, "gt=0")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Product's quantity must be greater than 0"))
	}

	if err == nil {
		_, err = databases.AddProduct(&new_product)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerError("Internal Server Error"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseNonData("Success Operation"))
}

// Controller for get all products
func GetAllProductsController(c echo.Context) error {
	products, err := databases.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Bad Request"))
	}
	if products == nil {
		return c.JSON(http.StatusNotFound, responses.StatusDataNotFound("Data Not Found"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseData("Success Operation", products))
}

// Controller for get all products by name ascending
func GetAllProductsByNameAscController(c echo.Context) error {
	products, err := databases.GetAllProductsOrderByNameAsc()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Bad Request"))
	}
	if products == nil {
		return c.JSON(http.StatusNotFound, responses.StatusDataNotFound("Data Not Found"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseData("Success Operation", products))
}

// Controller for get all products by name descending
func GetAllProductsByNameDescController(c echo.Context) error {
	products, err := databases.GetAllProductsOrderByNameAsc()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse("Bad Request"))
	}
	if products == nil {
		return c.JSON(http.StatusNotFound, responses.StatusDataNotFound("Data Not Found"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseData("Success Operation", products))
}
