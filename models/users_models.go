package models

import "gorm.io/gorm"

// struktur data users
type Users struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Phone        string `gorm:"unique" json:"phone" form:"phone"`
	Role         string
	Token        string
	GroupProduct []GroupProduct
	Order        []Order
}

// struktur get user *
type GetUser struct {
	ID    uint
	Name  string
	Email string
	Phone string
}

// struktur get login user
type GetLoginUser struct {
	ID    uint
	Name  string
	Token string
	Role  string
}
