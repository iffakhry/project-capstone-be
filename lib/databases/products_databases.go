package databases

import (
	"final-project/config"
	"final-project/models"
)

// function database untuk membuat data product baru
func CreateProduct(product *models.Products) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
