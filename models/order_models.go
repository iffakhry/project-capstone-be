package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UsersID        uint
	GroupProductID uint
	PriceOrder     int
	NameProduct    string
	Email          string
	Password       string
	Payment        Payment
}

type ResPayment struct {
	Phone string `json:"phone" form:"phone" `
}

type Detail struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Payment struct {
	OrderID     uint
	EwalletType string
	ExternalId  string
	Amount      float64
	BusinessId  string
	Created     string
}

type GetOrder struct {
	OrderID        uint
	UsersID        uint
	GroupProductID uint
	NameProduct    string
	PriceOrder     int
	EwalletType    string
	ExternalId     string
	Created        string
	Email          string
	Password       string
}

type GetUserOrder struct {
	OrderID        uint
	UsersID        uint
	GroupProductID uint
	Name           string
}

//down
