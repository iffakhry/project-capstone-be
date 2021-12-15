package databases

import (
	"final-project/config"
	"final-project/models"
)

var get_product []models.GetProduct

// function database untuk membuat data product baru
func CreateProduct(product *models.Products) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// function database untuk menampilkan data semua product
func GetAllProduct() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Find(&get_product)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_product, nil
}

// function database untuk menampilkan data product by id
func GetProductById(id int) (interface{}, error) {
	var get_product_by_id models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL AND products.id = ?", id).Find(&get_product_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_product_by_id, nil
}

// function database untuk menampilkan data url dan photo product by id
func GetPhotoUrlProductById(id int) (string, string, error) {
	var get_product_by_id models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL AND products.id = ?", id).Find(&get_product_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return "", "", query.Error
	}
	return get_product_by_id.Photo, get_product_by_id.Url, nil
}

// function database untuk memperbarui data product by id
func UpdateProduct(id int, update_product *models.Products) (interface{}, error) {
	var product models.Products
	query_select := config.DB.Find(&product, id)
	if query_select.Error != nil || query_select.RowsAffected == 0 {
		return 0, query_select.Error
	}
	query_update := config.DB.Model(&product).Updates(update_product)
	if query_update.Error != nil {
		return nil, query_update.Error
	}
	return product, nil
}

// function database untuk menghapus product by id
func DeleteProduct(id int) (interface{}, error) {
	var product models.Products
	check_product := config.DB.Find(&product, id).RowsAffected

	err := config.DB.Delete(&product).Error
	if err != nil || check_product > 0 {
		return nil, err
	}
	return product, nil
}
