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

// database function for get all product order by name ascending
func GetAllProductsOrderByNameAsc() (interface{}, error) {
	var get_products []models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.product_name ASC").Find(&get_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_products, nil
}

// database function for get all product order by name descending
func GetAllProductsOrderByNameDesc() (interface{}, error) {
	var get_products []models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.product_name DESC").Find(&get_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_products, nil
}

// database function for get all product order by price ascending
func GetAllProductsOrderByPriceAsc() (interface{}, error) {
	var get_products []models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.product_price ASC").Find(&get_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_products, nil
}

// database function for get all product order by price descending
func GetAllProductsOrderByPriceDesc() (interface{}, error) {
	var get_products []models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.product_price DESC").Find(&get_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_products, nil
}
