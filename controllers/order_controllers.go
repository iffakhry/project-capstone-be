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

func CreateOrderControllers(c echo.Context) error {
	new_oder := models.OrderRequest{}
	id_group, er := strconv.Atoi(c.Param("id_group"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Param"))
	}

	c.Bind(&new_oder)
	id_user, role := middlewares.ExtractTokenId(c)

	new_oder.Order.UsersID = uint(id_user)

	// mengecek apakah user sudah tergabung di group
	cek, e := databases.CekUserInGroup(uint(id_group), uint(id_user))
	if cek != nil || role == "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	} else {

		status, data, err := databases.CreateOrder(&new_oder, id_group)

		if status == "Full" {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Group Product Full"))
		}
		if err != nil || e != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
		}
		if data == nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Id Group Product Not Found"))
		}
		return c.JSON(http.StatusBadRequest, response.SuccessResponseNonData("Success Operation"))
	}
}

func GetOrderByIdOrderControllers(c echo.Context) error {
	id_order, err := strconv.Atoi(c.Param("id_order"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	token, _ := middlewares.ExtractTokenId(c)

	data, e, id_user := databases.GetOrderByIdOrder(id_order)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if id_user != uint(token) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetOrderByIdGroupControllers(c echo.Context) error {
	id_group, err := strconv.Atoi(c.Param("id_group"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	_, role := middlewares.ExtractTokenId(c)

	data, e := databases.GetOrderByIdGroup(id_group)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetOrderByIdUsersControllers(c echo.Context) error {
	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	token, _ := middlewares.ExtractTokenId(c)

	data, e := databases.GetOrderByIdUser(id_user)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if token != id_user {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

// func CekOrderControllers(c echo.Context) error {
// 	id := c.Param("id_group")
// 	id_group_product, err := strconv.Atoi(id)
// 	id_user, _ := middlewares.ExtractTokenId(c)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
// 	}
// 	data, e := databases.CekUserInGroup(uint(id_group_product), uint(id_user))
// 	if data == nil {
// 		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
// 	}
// 	if e != nil {
// 		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
// 	}
// 	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
// }
