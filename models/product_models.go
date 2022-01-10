package models

import "gorm.io/gorm"

// product structure
type Product struct {
	gorm.Model
	Product_Name        string `json:"product_name" form:"product_name"`
	Product_Price       int    `json:"product_price" form:"product_price"`
	Product_Description string `json:"product_description" form:"product_description"`
	Product_Qty         int    `json:"product_qty" form:"product_qty"`
}

// get product structure
type GetProduct struct {
	ID                  uint
	Product_Name        string
	Product_Price       int
	Product_Description string
	Product_Qty         int
}
