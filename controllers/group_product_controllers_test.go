package controllers

import (
	"final-project/config"
	"final-project/models"

	"github.com/labstack/echo/v4"
)

var (
	mock_data_user1 = models.Users{
		Name:     "user1",
		Email:    "user1@mail.com",
		Password: "user123",
		Phone:    "+628257237462",
	}
	mock_data_product = models.Products{
		Name_Product:   "Netflix",
		Detail_Product: "lorem",
		Price:          200000,
		Limit:          5,
		Photo:          "netflix.jpg",
	}
	mock_data_group = models.GroupProduct{
		ProductsID: 1,
	}
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertMockToDb() {
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
}
