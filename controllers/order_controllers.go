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
	id_user, _ := middlewares.ExtractTokenId(c)

	new_oder.Order.UsersID = uint(id_user)
	data, err := databases.CreateOrder(&new_oder, id_group)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Id Group Product Not Found"))
	}
	return c.JSON(http.StatusBadRequest, response.SuccessResponseData("Success Operation", data))

}
