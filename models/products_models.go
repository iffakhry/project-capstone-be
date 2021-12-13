package models

import "gorm.io/gorm"

// struktur data products
type Products struct {
	gorm.Model
	Name_Product   string `json:"name_product" form:"name_product"`
	Detail_Product string `json:"detail_product" form:"detail_product"`
	Price          int    `json:"price" form:"price"`
	Limit          int    `json:"limit" form:"limit"`
	Photo          string `json:"photo" form:"photo"`
	Url            string
	UsersID        uint
	GroupProduct   []GroupProduct
}

type GetProduct struct {
	Name_Product   string
	Detail_Product string
	Price          int
	Limit          int
	Photo          string
	Url            string
}
