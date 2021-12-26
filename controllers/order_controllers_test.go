package controllers

import (
	"bytes"
	"encoding/json"
	"final-project/config"
	"final-project/constants"
	"final-project/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

// type GroupResponseSuccess struct {
// 	Message string
// 	Data    []models.GroupProduct
// }

var (
	mock_data_payment = models.ResPayment{
		Phone: "08563729874",
	}
	mock_data_payment_kurang = models.ResPayment{
		Phone: "0856372",
	}
	mock_data_payment_salah = models.ResPayment{
		Phone: "12563722343",
	}
)

func TestCreateOrderControllerFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/orders/:id_group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTUser()
	if err != nil {
		panic(err)
	}

	InsertMockToDb()
	// config.DB.Migrator().DropTable(models.Order{})
	t.Run("access_forbidden", func(t *testing.T) {
		token_admin, err := UsingJWTAdmin()
		if err != nil {
			panic(err)
		}
		testCases.name = "Access Forbidden"
		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_admin))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)

	})
	t.Run("invalid_id", func(t *testing.T) {
		testCases.name = "Invalid Id"
		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("as")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)

	})
	t.Run("Invalid_Telephone_Number", func(t *testing.T) {
		testCases.name = "Invalid Telephone Number"
		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)

	})
	t.Run("Invalid_Telephone_Number", func(t *testing.T) {
		testCases.name = "Invalid Telephone Number"
		body, error := json.Marshal(mock_data_payment_kurang)
		if error != nil {
			t.Error(t, error, "error marshal")
		}

		req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

		body_res := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body_res), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)

	})
	t.Run("Invalid_Telephone_Number", func(t *testing.T) {
		testCases.name = "Invalid Telephone Number"
		body, error := json.Marshal(mock_data_payment_salah)
		if error != nil {
			t.Error(t, error, "error marshal")
		}

		req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

		body_res := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body_res), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)

	})
	t.Run("Bad_Request", func(t *testing.T) {
		testCases.name = "Bad Request"
		body, error := json.Marshal(mock_data_payment)
		if error != nil {
			t.Error(t, error, "error marshal")
		}
		config.DB.Migrator().DropTable(models.Order{})
		req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

		body_res := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body_res), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)

	})
	// t.Run("Id Group Product Not Found", func(t *testing.T) {
	// 	testCases.name = "Id Group Product Not Found"
	// 	config.DB.Save(&mock_data_product2)
	// 	config.DB.Save(&mock_data_group2)
	// 	body, error := json.Marshal(mock_data_payment)
	// 	if error != nil {
	// 		t.Error(t, error, "error marshal")
	// 	}
	// 	// config.DB.Migrator().DropTable(models.Order{})
	// 	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	// 	res := httptest.NewRecorder()
	// 	context := e.NewContext(req, res)
	// 	context.SetPath(testCases.path)
	// 	context.SetParamNames("id_group")
	// 	context.SetParamValues("2")

	// 	middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

	// 	body_res := res.Body.String()
	// 	var responses GroupResponseSuccess
	// 	err = json.Unmarshal([]byte(body_res), &responses)
	// 	if err != nil {
	// 		assert.Error(t, err, "error")
	// 	}
	// 	assert.Equal(t, testCases.expectCode, res.Code)
	// 	assert.Equal(t, testCases.name, responses.Message)

	// })

}
