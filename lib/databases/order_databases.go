package databases

import (
	"final-project/config"
	"final-project/models"
)

func CreateReservation(Order *models.OrderRequest) (interface{}, error) {

	Order.Order.UsersID = 1
	Order.Order.GroupProductID = 1
	Order.Order.PriceOrder = 1000
	Order.Order.NameProduct = "netfix"
	Order.Order.DetailCredential = "email passsword"

	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}
	req_order := models.Order{}

	req_credit := models.CreditCard{
		Typ:    Order.Credit.Typ,
		Name:   Order.Credit.Name,
		Number: Order.Credit.Number,
		Cvv:    Order.Credit.Cvv,
		Month:  Order.Credit.Month,
		Year:   Order.Credit.Year,
	}

	Create_Res := models.OrderRequest{
		Order:  req_order,
		Credit: req_credit,
	}
	// Order.Credit.ReservationID = Order.Reservation.ID
	config.DB.Create(&Order.Credit)
	return Create_Res, nil
}
