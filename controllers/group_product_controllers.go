package controllers

import (
	"final-project/lib/databases"
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
	new_group.UsersID = uint(id_user)

	_, er := databases.CreateGroupProduct(&new_group)
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

func GetGroupProductControllers(c echo.Context) error {
	id := c.Param("id")
	id_group_product, err := strconv.Atoi(id)
	// log.Println("id", id_group_product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	data, e := databases.GetGroupProductById(id_group_product)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

// controller untuk menampilkan seluruh data users
func GetAllGroupProductControllers(c echo.Context) error {
	data, _, err := databases.GetAllGroupProduct()
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetAvailableGroupProductControllers(c echo.Context) error {
	status := c.Param("status")
	// id_group_product, err := strconv.Atoi(id)
	// log.Println("id", id_group_product)
	if status != "available" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Param"))
	}
	data, e := databases.GetGroupProductByAvailable(status)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}
