package databases

import (
	"final-project/config"
	"final-project/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func CreateOrder(Payment *models.ResPayment, Order *models.Order, id_group int) (interface{}, error) {
	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}

	UpdatePlusGroupProductCapacity(id_group)
	Create_Res, _ := PaymentXendit(Order.ID, Payment.Phone, Order.PriceOrder)

	return Create_Res, nil
}

func GetOrderByIdOrder(id int) (i interface{}, e error, id_user uint) {
	order := models.GetOrder{}
	query := config.DB.Table("orders").Select("*").Where("orders.deleted_at IS NULL AND orders.id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error, 0
	}
	res_idorder, _ := CekUserInGroup(order.GroupProductID, order.UsersID)
	order.OrderID = res_idorder
	return order, nil, order.UsersID
}

func GetOrderByIdGroup(id int) (interface{}, error, int) {
	order := []models.GetOrder{}
	query := config.DB.Table("orders").Select("*").Where("orders.deleted_at IS NULL AND orders.group_product_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error, 0
	}
	for i, _ := range order {
		res_idorder, _ := CekUserInGroup(order[i].GroupProductID, order[i].UsersID)
		order[i].OrderID = res_idorder
	}
	capacity := len(order)
	return order, nil, capacity
}
func GetOrderByIdUser(id int) (i interface{}, e error) {
	order := []models.GetOrder{}
	query := config.DB.Table("orders").Select("*").Where("orders.deleted_at IS NULL AND orders.users_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	for i, _ := range order {
		res_idorder, _ := CekUserInGroup(order[i].GroupProductID, order[i].UsersID)
		order[i].OrderID = res_idorder
	}

	return order, nil
}

func GetUserOrderByIdGroup(id int) (interface{}, error) {
	order := []models.GetUserOrder{}
	query := config.DB.Table("orders").Select("users.name,orders.users_id,orders.group_product_id").Joins("join users on orders.users_id = users.id").Where("orders.deleted_at IS NULL AND orders.group_product_id = ? ", id).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	for i, _ := range order {
		res_idorder, _ := CekUserInGroup(order[i].GroupProductID, order[i].UsersID)
		order[i].OrderID = res_idorder
	}
	return order, nil
}
func CekUserInGroup(id_group, id_user uint) (id_order uint, er error) {
	order := models.Order{}
	query := config.DB.Where("deleted_at IS NULL AND group_product_id = ? && users_id = ?", id_group, id_user).Find(&order)
	if query.Error != nil || query.RowsAffected < 1 {
		return 0, query.Error
	}
	return order.ID, nil
}

func UpdateOrderDetail(id_order int, email, password string) (interface{}, error) {
	order := models.Order{}
	query1 := config.DB.Model(&order).Where("orders.deleted_at IS NULL AND orders.id = ?", id_order).Updates(map[string]interface{}{"email": email, "password": password}).Find(&order)
	if query1.Error != nil {
		return nil, query1.Error
	}
	res, _, _ := GetOrderByIdOrder(id_order)
	return res, nil
}

func PaymentXendit(id_order uint, phone string, amount int) (interface{}, error) {

	xendit.Opt.SecretKey = os.Getenv("KEY_XENDIT")

	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	data := ewallet.CreatePaymentParams{
		ExternalID:  "OVO-ewallet-" + formatted,
		Amount:      float64(amount),
		Phone:       phone,
		EWalletType: xendit.EWalletTypeOVO,
		CallbackURL: "mystore.com/callback",
		RedirectURL: "mystore.com/redirect",
	}

	resp, err := ewallet.CreatePayment(&data)
	if err != nil {
		log.Fatal(err)
	}

	res_pay := models.Payment{
		OrderID:     id_order,
		EwalletType: string(resp.EWalletType),
		ExternalId:  resp.ExternalID,
		Amount:      resp.Amount,
		BusinessId:  resp.BusinessID,
		Created:     resp.Created.Format("02-01-2006"),
	}

	if err := config.DB.Create(&res_pay).Error; err != nil {
		return nil, err
	}
	return res_pay, nil
}

func DeleteOrder(id_order int) (interface{}, error) {
	order := models.Order{}
	config.DB.Find(&order, id_order)
	UpdateMinusGroupProductCapacity(int(order.GroupProductID))
	if err := config.DB.Where("id = ?", id_order).Delete(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
