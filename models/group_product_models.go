package models

import "gorm.io/gorm"

type GroupProduct struct {
	gorm.Model
	UsersID              uint
	ProductsID           uint `json:"products_id" form:"products_id"`
	NameGroupProduct     string
	CapacityGroupProduct int
	AdminFee             int
	TotalPrice           int
	DurationGroup        string
	Status               string
	Order                []Order
}

type GetGroupProduct struct {
	ID                   uint
	ProductsID           uint
	NameGroupProduct     string
	Limit                int
	CapacityGroupProduct int
	Price                int
	AdminFee             int
	TotalPrice           int
	DurationGroup        string
	Name_Product         string
	Status               string
	Url                  string
	GetOrder             interface{}
}

type ResGroup struct {
	GroupProductID uint
}
