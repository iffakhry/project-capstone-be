package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UsersID          uint
	GroupProductID   uint
	PriceOrder       int
	NameProduct      string
	DetailCredential string
	CreditCard       CreditCard
}

type CreditCard struct {
	OrderID uint
	Typ     string `json:"typ" `
	Name    string `json:"name" `
	Number  string `json:"number" `
	Cvv     int    `json:"cvv" `
	Month   int    `json:"month" `
	Year    int    `json:"year" `
}

type OrderRequest struct {
	Order      Order      `json:"order" `
	CreditCard CreditCard `json:"credit_card" `
}

// type GetReserv struct {
// 	UsersID       uint
// 	HomestayID    uint
// 	Name_Homestay string `json:"name_homestay" form:"name_homestay"`
// 	Start_date    string
// 	End_date      string
// 	Price         int `json:"price" form:"price"`
// 	Total_harga   int
// }

// type CekStatus struct {
// 	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
// 	Start_date string `json:"start_date" form:"start_date"`
// 	End_date   string `json:"end_date" form:"end_date"`
// }
