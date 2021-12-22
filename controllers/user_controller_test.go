package controllers

import (
	"bytes"
	"encoding/json"
	"final-project/config"
	"final-project/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type UserResponse struct {
	Message string
	Data    models.Users
}

type Login struct {
	Email    string
	Password string
}

// data dummy
var (
	mock_data_user = models.Users{
		Name:     "sahril",
		Email:    "sahril@gmail.com",
		Password: "bismillah",
		Phone:    "+628123456789",
	}
)

// inisialisasi echo
func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

// menambahkan user
func InsertUser() error {
	if err := config.DB.Save(&mock_data_user).Error; err != nil {
		return err
	}
	return nil
}

// Create User Controller : Test Case 1 (All Input is valid)
func TestCreateUserController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/users",
		expectCode: http.StatusOK,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 2 (Input invalid Name)
func TestCreateUserControllerFailed_1(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Invalid Name",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "",
		Email:    "sahril@gmail.com",
		Password: "bismillah",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 3 (Name contains non-alphanumeric)
func TestCreateUserControllerFailed_2(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Name can only contains alphanumeric",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "sahril mahendra",
		Email:    "sahril@gmail.com",
		Password: "bismillah",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 4 (Invalid email)
func TestCreateUserControllerFailed_3(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Invalid Email",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "@gmail.com",
		Password: "bismillah",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 5 (Invalid Password)
func TestCreateUserControllerFailed_4(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Invalid Password",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "sahril2@gmail.com",
		Password: "",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 6 (Password contain less than 6 characters)
func TestCreateUserControllerFailed_5(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Password must consist of 6 characters or more",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "sahril2@gmail.com",
		Password: "12345",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 7 (Invalid telephone number)
func TestCreateUserControllerFailed_6(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Invalid Telephone Number",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "sahril2@gmail.com",
		Password: "123456",
		Phone:    "0281",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 8 (Create user as admin at first)
func TestCreateUserControllerFailed_7(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Success Operation",
		path:       "/users",
		expectCode: http.StatusOK,
	}

	var newDataUser = models.Users{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "qwerty",
		Phone:    "+62811222112",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// Create User Controller : Test Case 9 (Email or Telephone Number Already Exist)
func TestCreateUserControllerFailed_8(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Email or Telephone Number Already Exist",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	var newDataUser = models.Users{
		Name:     "sahrilbaru",
		Email:    "sahril@gmail.com",
		Password: "qwerty",
		Phone:    "+628111222111",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}
