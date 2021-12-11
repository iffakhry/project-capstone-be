package controllers

import (
	"final-project/middlewares"
	"final-project/models"
	response "final-project/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan user (registrasi)
func CreateGroupProductControllers(c echo.Context) error {
	new_group := models.GroupProduct{}
	c.Bind(&new_group)

	id_user, _ := middlewares.ExtractTokenId(c)

	new_group.AdminFee = 6500
	new_group.CapacityGroupProduct = 4
	new_group.Duration = "1-12-2021"
	new_group.TotalPrice = new_group.AdminFee + 45000
	new_group.NameGroupProduct = "Group Product " + strconv.Itoa(int(new_group.ID))

	if id_user == 1 {

	}

	// v := validator.New()
	// err := v.Var(new_group.Name, "required")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Name"))
	// }
	// err = v.Var(new_group.Email, "required,email")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Email"))
	// }
	// err = v.Var(new_group.Password, "required")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Password"))
	// }
	// if len(new_group.Password) < 6 {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Password must consist of 6 characters or more"))
	// }
	// err = v.Var(new_group.Phone, "required,e164")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	// }
	// if new_group.Email == "admin@admin.com" {
	// 	new_group.Role = "admin"
	// } else {
	// 	new_group.Role = "user"
	// }

	// if err == nil {
	// 	new_group.Password, _ = helper.HashPassword(new_group.Password) // generate plan password menjadi hash
	// 	_, err = databases.CreateUser(&new_group)
	// }
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	// }
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
