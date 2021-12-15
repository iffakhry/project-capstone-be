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

type GetOrder struct {
	UsersID          uint
	GroupProductID   uint
	PriceOrder       int
	NameProduct      string
	DetailCredential string
}

type GetUserOrder struct {
	UsersID        uint
	GroupProductID uint
	Name           string
}
