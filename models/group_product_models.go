package models

import "gorm.io/gorm"

type GroupProduct struct {
	gorm.Model
	UsersID              uint
	ProductID            uint `json:"id_product" form:"id_product"`
	NameGroupProduct     string
	CapacityGroupProduct int
	AdminFee             int
	TotalPrice           int
	Duration             string
	Status               string
}

type GetGroupProduct struct {
	ID                   uint
	ProductID            uint
	NameGroupProduct     string
	CapacityGroupProduct int
	AdminFee             int
	TotalPrice           int
	Duration             string
	Status               string
}
