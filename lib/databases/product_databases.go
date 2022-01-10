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

// database function for get all product
func GetAllProducts() (interface{}, error) {
	var get_products []models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Find(&get_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_products, nil
}
