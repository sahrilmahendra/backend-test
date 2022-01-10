package databases

import (
	"erajaya/config"
	"erajaya/models"
)

// database function for add product
func AddProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
