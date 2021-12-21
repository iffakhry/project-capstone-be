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
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Param"))
	}

	c.Bind(&new_payment)
	v := validator.New()
	var regx, _ = regexp.Compile(`^08[1-9][0-9].*$`)
	var len_phone = len(new_payment.Phone)

	id_user, role := middlewares.ExtractTokenId(c)
	t_price, _, _, n_product, status, er := databases.GetDataGroupProductById(id_group)

	new_order.UsersID = uint(id_user)
	new_order.GroupProductID = uint(id_group)
	new_order.PriceOrder = t_price
	new_order.NameProduct = n_product
	new_order.DetailCredential = "Email: , Password: "

	// mengecek apakah user sudah tergabung di group
	cek, e := databases.CekUserInGroup(uint(id_group), uint(id_user))
	if er != nil || e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if cek != 0 || role == "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	erro := v.Var(new_payment.Phone, "required")
	if erro != nil || len_phone < 11 || len_phone > 13 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	} else if !regx.MatchString(new_payment.Phone) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	} else {

		data, err := databases.CreateOrder(&new_payment, &new_order, id_group)

		if status != "Available" {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Group Product Full"))
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
		}
		if data == nil || t_price == 0 {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Id Group Product Not Found"))
		}
		return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
	}
}

func GetOrderByIdOrderControllers(c echo.Context) error {
	id_order, err := strconv.Atoi(c.Param("id_order"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	token, _ := middlewares.ExtractTokenId(c)

	data, e, id_user := databases.GetOrderByIdOrder(id_order)
	if id_user != uint(token) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	_, role := middlewares.ExtractTokenId(c)

	data, e := databases.GetOrderByIdGroup(id_group)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	token, _ := middlewares.ExtractTokenId(c)

	data, e := databases.GetOrderByIdUser(id_user)
	if token != id_user {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
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
	v := validator.New()
	erro := v.Var(detail.DetailCredential, "required")
	if erro != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Details Can't Be Empty"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	_, role := middlewares.ExtractTokenId(c)

	data, e := databases.UpdateOrderDetail(id_order, detail.DetailCredential)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}
