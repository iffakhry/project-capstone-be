package models

import "gorm.io/gorm"

type GrupProduct struct {
	gorm.Model
	UsersID              uint
	ProductID            uint `json:"id_product" form:"id_product"`
	NameGroupProduct     string
	CapacityGroupProduct int
	AdminFee             int
	TotalPrice           int
	Duration             string
}
