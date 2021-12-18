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
	Payment          Payment
}

type ResPayment struct {
	Phone  string `json:"phone" `
	Amount int    `json:"amount" `
}

type OrderRequest struct {
	Order      Order      `json:"order" `
	ResPayment ResPayment `json:"payment" `
}

type Detail struct {
	DetailCredential string `json:"detail" form:"detail"`
}

type GetOrder struct {
	UsersID          uint
	GroupProductID   uint
	NameProduct      string
	PriceOrder       int
	DetailCredential string
}

type GetUserOrder struct {
	UsersID        uint
	GroupProductID uint
	Name           string
}

type Payment struct {
	OrderID     uint
	EwalletType string
	ExternalId  string
	Amount      float64
	BusinessId  string
	Created     string
}
