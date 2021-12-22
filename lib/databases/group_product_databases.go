package databases

import (
	"final-project/config"
	"final-project/models"
)

var query_join string = "products.limit,products.name_product,products.url,products.price,group_products.id,group_products.products_id,group_products.name_group_product,group_products.capacity_group_product,group_products.admin_fee,group_products.total_price,group_products.duration_group,group_products.status"

// function database untuk menambahkan user baru (registrasi)
func CreateGroupProduct(group *models.GroupProduct, id_product uint) (interface{}, error) {

	if err := config.DB.Create(&group).Error; err != nil {
		return nil, err
	}
	Create_Res := models.ResGroup{
		GroupProductID: group.ID,
	}

	return Create_Res, nil

}

func GetAllGroupProduct() (interface{}, int, error) {
	res_group := []models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select(query_join).Joins("join products on group_products.products_id = products.id").Where("group_products.deleted_at IS NULL").Find(&res_group)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, 0, query.Error
	}
	for i, _ := range res_group {
		user_order, _ := GetUserOrderByIdGroup(int(res_group[i].ID))
		res_group[i].GetOrder = user_order
	}
	return res_group, len(res_group), nil
}

func GetGroupProductById(id int) (interface{}, error) {
	group := models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select(query_join).Joins("join products on group_products.products_id = products.id").Where("group_products.deleted_at IS NULL AND group_products.id = ? ", id).Find(&group)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	user_order, _ := GetUserOrderByIdGroup(int(group.ID))
	group.GetOrder = user_order

	return group, nil
}

func GetGroupProductByAvailable(str string) (interface{}, error) {
	group := []models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select(query_join).Joins("join products on group_products.products_id = products.id").Where("group_products.deleted_at IS NULL AND group_products.status = ?", str).Find(&group)

	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	for i, _ := range group {
		user_order, _ := GetUserOrderByIdGroup(int(group[i].ID))
		group[i].GetOrder = user_order
	}
	return group, nil
}

func GetGroupProductByIdProducts(id_products int) (interface{}, error) {
	res_group := []models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select(query_join).Joins("join products on group_products.products_id = products.id").Where("group_products.deleted_at IS NULL AND group_products.products_id = ? ", id_products).Find(&res_group)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	for i, _ := range res_group {
		user_order, _ := GetUserOrderByIdGroup(int(res_group[i].ID))
		res_group[i].GetOrder = user_order
	}
	return res_group, nil
}

func UpdatePlusGroupProductCapacity(id_group_product int) (interface{}, error) {
	group := models.GroupProduct{}
	config.DB.Find(&group, id_group_product)
	_, _, limit, _ := GetDataProduct(int(group.ProductsID))
	if group.CapacityGroupProduct >= limit-1 {
		query := config.DB.Model(&group).Where("group_products.id = ?", id_group_product).Updates(map[string]interface{}{"status": "Full", "capacity_group_product": limit})
		if query.Error != nil {
			return nil, query.Error
		}
		return group, nil
	}
	capacity := group.CapacityGroupProduct + 1
	query1 := config.DB.Model(&group).Where("group_products.id = ?", id_group_product).Update("capacity_group_product", capacity)
	if query1.Error != nil {
		return nil, query1.Error
	}
	return group, nil
}

func UpdateMinusGroupProductCapacity(id_group_product int) (interface{}, error) {
	group := models.GroupProduct{}
	config.DB.Find(&group, id_group_product)
	capacity := group.CapacityGroupProduct - 1
	query1 := config.DB.Model(&group).Where("group_products.id = ?", id_group_product).Updates(map[string]interface{}{"status": "Available", "capacity_group_product": capacity})
	if query1.Error != nil {
		return nil, query1.Error
	}
	return group, nil
}

func GetDataProduct(id_product int) (name string, price, limit int, er error) {
	var get_product_by_id models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL AND products.id = ?", id_product).Find(&get_product_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return "", 0, 0, query.Error
	}
	return get_product_by_id.Name_Product, get_product_by_id.Price, get_product_by_id.Limit, nil
}

func GetDataGroupProductById(id int) (t_price, limit int, n_group, n_product, status string, er error) {
	res_group := models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select(query_join).Joins("join products on group_products.products_id = products.id").Where("group_products.deleted_at IS NULL AND group_products.id = ? ", id).Find(&res_group)
	if query.Error != nil || query.RowsAffected < 1 {
		return 0, 0, "", "", "", query.Error
	}
	return res_group.TotalPrice, res_group.Limit, res_group.NameGroupProduct, res_group.Name_Product, res_group.Status, nil
}

func DeleteGroupProduct(id_group int) (interface{}, error) {
	group := models.GroupProduct{}
	if err := config.DB.Where("id = ?", id_group).Delete(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
