package controllers

import (
	"final-project/helper"
	"final-project/lib/databases"
	"final-project/middlewares"
	"final-project/models"
	response "final-project/responses"
	"log"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// controller untuk menampilkan seluruh data users
func GetAllUsersControllers(c echo.Context) error {
	users, err := databases.GetAllUsers()
	if users == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", users))
}

// controller untuk menampilkan data user by id
func GetUserControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	log.Println("id", conv_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	user, e := databases.GetUserById(conv_id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", user))
}

// controller untuk menambahkan user (registrasi)
func CreateUserControllers(c echo.Context) error {
	new_user := models.Users{}
	c.Bind(&new_user)

	v := validator.New()
	err := v.Var(new_user.Name, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Name"))
	}
	err = v.Var(new_user.Email, "required, email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Email"))
	}
	err = v.Var(new_user.Password, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Password"))
	}
	if len(new_user.Password) < 6 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Password must consist of 6 characters or more"))
	}
	err = v.Var(new_user.Phone, "required,e164")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	}
	if err == nil {
		new_user.Password, _ = helper.HashPassword(new_user.Password) // generate plan password menjadi hash
		_, err = databases.CreateUser(&new_user)
	}
	// check, _ := databases.GetUserByEmail(new_user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menghapus user by id
func DeleteUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}

	user, _ := databases.GetUserById(id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}

	logged := middlewares.ExtractTokenId(c) // check token
	if logged != id {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	databases.DeleteUser(id)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk memperbarui data user by id (update)
func UpdateUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}

	user, _ := databases.GetUserById(id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}

	logged := middlewares.ExtractTokenId(c) // check token
	if logged != id {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	users := models.Users{}
	c.Bind(&users)

	v := validator.New()
	er := v.Var(users.Name, "required")
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Name"))
	}
	er = v.Var(users.Email, "required, email")
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Email"))
	}
	er = v.Var(users.Password, "required")
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Password"))
	}
	if len(users.Password) < 6 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Password must consist of 6 characters or more"))
	}
	er = v.Var(users.Phone, "required,e164")
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Telephone Number"))
	}
	if er == nil {
		users.Password, _ = helper.HashPassword(users.Password) // generate plan password menjadi hash
		_, er = databases.UpdateUser(id, &users)
	}
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk login dan generate token (by email dan password)
func LoginUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	plan_pass := user.Password
	log.Println(plan_pass)
	token, e := databases.LoginUser(plan_pass, &user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Email or Password Incorrect"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Login Success", token))
}

// controller untuk kebutuhan testing get user
func GetUserControllersTesting() echo.HandlerFunc {
	return GetUserControllers
}

// controller untuk kebutuhan testing update user
func UpdateUserControllersTesting() echo.HandlerFunc {
	return UpdateUserControllers
}

// controller untuk kebutuhan testing delete user
func DeleteUserControllersTesting() echo.HandlerFunc {
	return DeleteUserControllers
}
