package models

import "gorm.io/gorm"

// struktur data products
type Products struct {
	gorm.Model
	Name_Product   string `json:"name_product" form:"name_product"`
	Detail_Product string `json:"detail_product" form:"detail_product"`
	Price          int    `json:"price" form:"price"`
	Photo          string `json:"photo" form:"photo"`
	Url            string
}
