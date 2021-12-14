package databases

import (
	"final-project/config"
	"final-project/models"
)

func CreateOrder(Order *models.OrderRequest, id_group int) (interface{}, error) {
	t_price, limit, _, n_product, er := GetDataGroupProductById(id_group)
	if t_price == 0 {
		return nil, er
	}

	Order.Order.GroupProductID = uint(id_group)
	Order.Order.PriceOrder = t_price / limit
	Order.Order.NameProduct = n_product
	Order.Order.DetailCredential = "email passsword"

	if err := config.DB.Create(&Order.Order).Error; err != nil {
		return nil, err
	}
	req_credit := models.CreditCard{
		Typ:    Order.CreditCard.Typ,
		Name:   Order.CreditCard.Name,
		Number: Order.CreditCard.Number,
		Cvv:    Order.CreditCard.Cvv,
		Month:  Order.CreditCard.Month,
		Year:   Order.CreditCard.Year,
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
	return Create_Res, nil
}

func GetOrderByIdGroup(id int) (interface{}, error) {
	order := []models.GetOrder{}
	// UpdateGroupProductCapacity(id)
	query := config.DB.Table("orders").Select("users.name,orders.users_id,orders.group_product_id").Joins("join users on orders.users_id = users.id").Where("orders.deleted_at IS NULL AND orders.group_product_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return order, nil
}
