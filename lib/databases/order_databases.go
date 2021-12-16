package databases

import (
	"final-project/config"
	"final-project/models"
	"fmt"
)

func CreateOrder(Order *models.OrderRequest, id_group int) (in string, i interface{}, er error) {
	t_price, limit, _, n_product, status, er := GetDataGroupProductById(id_group)
	if t_price == 0 {
		return "", nil, er
	}
	fmt.Println("cek status", status)
	if status != "Available" {
		return "Full", nil, er
	}

	Order.Order.GroupProductID = uint(id_group)
	Order.Order.PriceOrder = t_price / limit
	Order.Order.NameProduct = n_product
	Order.Order.DetailCredential = "email: subs.spotify@mail.com, password: spotify123"

	if err := config.DB.Create(&Order.Order).Error; err != nil {
		return "", nil, err
	}
	req_credit := models.CreditCard{
		OrderID: Order.Order.ID,
		Typ:     Order.CreditCard.Typ,
		Name:    Order.CreditCard.Name,
		Number:  Order.CreditCard.Number,
		Cvv:     Order.CreditCard.Cvv,
		Month:   Order.CreditCard.Month,
		Year:    Order.CreditCard.Year,
	}
	req_order := models.Order{
		CreditCard: req_credit,
	}

	Create_Res := models.OrderRequest{
		Order:      req_order,
		CreditCard: req_credit,
	}
	Order.CreditCard.OrderID = Order.Order.ID
	config.DB.Create(&Order.CreditCard)
	UpdateGroupProductCapacity(id_group)

	return "", Create_Res, nil
}

func GetOrderByIdOrder(id int) (i interface{}, e error, id_user uint) {
	order := models.GetOrder{}
	// UpdateGroupProductCapacity(id)
	query := config.DB.Table("orders").Select("*").Where("orders.deleted_at IS NULL AND orders.id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error, 0
	}
	return order, nil, order.UsersID
}

func GetOrderByIdGroup(id int) (i interface{}, e error) {
	order := []models.GetOrder{}
	// UpdateGroupProductCapacity(id)
	query := config.DB.Table("orders").Select("*").Where("orders.deleted_at IS NULL AND orders.group_product_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return order, nil
}
func GetOrderByIdUser(id int) (i interface{}, e error) {
	order := []models.GetOrder{}
	// UpdateGroupProductCapacity(id)
	query := config.DB.Table("orders").Select("*").Where("orders.deleted_at IS NULL AND orders.users_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return order, nil
}

func GetUserOrderByIdGroup(id int) (interface{}, error) {
	order := []models.GetUserOrder{}
	// UpdateGroupProductCapacity(id)
	query := config.DB.Table("orders").Select("users.name,orders.users_id,orders.group_product_id").Joins("join users on orders.users_id = users.id").Where("orders.deleted_at IS NULL AND orders.group_product_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return order, nil
}

func CekUserInGroup(id_group, id_user uint) (interface{}, error) {
	order := models.Order{}
	query := config.DB.Where("group_product_id = ? && users_id = ?", id_group, id_user).Find(&order)
	fmt.Println("cek roww", query.RowsAffected)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return order, nil
}

func UpdateOrderDetail(id_order int, detail string) (interface{}, error) {
	order := models.Order{}
	res := models.GetOrder{}
	query1 := config.DB.Model(&order).Where("orders.id = ?", id_order).Update("detail_credential", detail).Find(&res, id_order)
	if query1.Error != nil {
		return nil, query1.Error
	}
	return res, nil
}
