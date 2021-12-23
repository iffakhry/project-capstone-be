package controllers

import (
	"final-project/lib/databases"
	"final-project/middlewares"
	"final-project/models"
	response "final-project/responses"
	"net/http"
	"regexp"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateOrderControllers(c echo.Context) error {
	new_order := models.Order{}
	new_payment := models.ResPayment{}
	id_group, er := strconv.Atoi(c.Param("id_group"))

	c.Bind(&new_payment)
	v := validator.New()
	var len_phone = len(new_payment.Phone)

	id_user, role := middlewares.ExtractTokenId(c)
	t_price, _, _, n_product, status, errr := databases.GetDataGroupProductById(id_group)

	new_order.UsersID = uint(id_user)
	new_order.GroupProductID = uint(id_group)
	new_order.PriceOrder = t_price
	new_order.NameProduct = n_product
	new_order.DetailCredential = ""

	// mengecek apakah user sudah tergabung di group
	cek, e := databases.CekUserInGroup(uint(id_group), uint(id_user))
	if cek != 0 || role == "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	erro := v.Var(new_payment.Phone, "required")
	if erro != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	}
	if len_phone < 11 || len_phone > 13 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	}
	if !regexp.MustCompile(`^08[1-9][0-9].*$`).MatchString(new_payment.Phone) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	} else {

		data, err := databases.CreateOrder(&new_payment, &new_order, id_group)

		if err != nil || errr != nil || e != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
		}
		if data == nil || t_price == 0 {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Id Group Product Not Found"))
		}
		if status != "Available" {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Group Product Full"))
		}
		return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
	}
}

func GetOrderByIdOrderControllers(c echo.Context) error {
	id_order, err := strconv.Atoi(c.Param("id_order"))
	token, role := middlewares.ExtractTokenId(c)
	data, e, id_user := databases.GetOrderByIdOrder(id_order)

	if id_user != uint(token) && role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetOrderByIdGroupControllers(c echo.Context) error {
	id_group, err := strconv.Atoi(c.Param("id_group"))
	_, role := middlewares.ExtractTokenId(c)
	data, e, _ := databases.GetOrderByIdGroup(id_group)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetOrderByIdUsersControllers(c echo.Context) error {
	id_user, err := strconv.Atoi(c.Param("id_user"))
	token, role := middlewares.ExtractTokenId(c)
	data, e := databases.GetOrderByIdUser(id_user)

	if token != id_user && role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func UpdateOrderControllers(c echo.Context) error {
	detail := models.Detail{}
	id_order, err := strconv.Atoi(c.Param("id_order"))
	c.Bind(&detail)
	_, role := middlewares.ExtractTokenId(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}

	v := validator.New()
	erro := v.Var(detail.DetailCredential, "required")
	if erro != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Details Can't Be Empty"))
	}

	cek, _, _ := databases.GetOrderByIdOrder(id_order)
	if cek == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}

	data, e := databases.UpdateOrderDetail(id_order, detail.DetailCredential)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func DeleteOrderControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id_order"))
	logged, role := middlewares.ExtractTokenId(c) // check token
	if logged != id && role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	data, _, _ := databases.GetOrderByIdOrder(id)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	databases.DeleteOrder(id)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
