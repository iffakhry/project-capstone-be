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
	id := c.Param("id_products")
	id_product, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Param"))
	}
	c.Bind(&new_group)

	id_user, role := middlewares.ExtractTokenId(c)
	if role == "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	new_group.UsersID = uint(id_user)
	new_group.ProductsID = uint(id_product)

	d, er := databases.CreateGroupProduct(&new_group, new_group.ProductsID)
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if d == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Id Product Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

func GetByIdGroupProductControllers(c echo.Context) error {
	id := c.Param("id_group")
	id_group_product, err := strconv.Atoi(id)
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
func GetByIdProductsGroupProductControllers(c echo.Context) error {
	id := c.Param("id_products")
	id_products, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	data, e := databases.GetGroupProductByIdProducts(id_products)
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
