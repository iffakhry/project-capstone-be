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
	mock_data_group3 = models.GroupProduct{
		UsersID:              2,
		ProductsID:           1,
		NameGroupProduct:     "netflix-2",
		CapacityGroupProduct: 1,
		AdminFee:             5000,
		TotalPrice:           1,
		Status:               "Full",
	}
	mock_data_order2 = models.Order{
		UsersID:        1,
		GroupProductID: 1,
		PriceOrder:     45000,
	}
	mock_data_xendit2 = models.Payment{
		OrderID:     1,
		Amount:      0,
		EwalletType: "OVO",
		ExternalId:  "1982773",
	}
)

type OrderResponse struct {
	Message string
	Data    models.Order
}

func InsertMockToDbFailed() {
	config.DB.Save(&mock_data_admin)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product2)
	config.DB.Save(&mock_data_group3)
	config.DB.Save(&mock_data_xendit2)
	// config.DB.Save(&mock_data_order)
}

func TestCreateOrderControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/orders/:id_group",
		expectCode: http.StatusOK,
	}
	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTUser()
	if err != nil {
		panic(err)
	}
	config.InitMigrateTest()

	// InsertMockToDbFailed()
	// config.DB.Migrator().DropTable(models.Order{})
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_user2)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	// config.DB.Save(&mock_data_order)
	// config.DB.Save(&mock_data_xendit)

	body, error := json.Marshal(mock_data_payment)
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
	fmt.Println("body", body_res)
	var responses GroupResponseSuccess
	err = json.Unmarshal([]byte(body_res), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}
func TestCreateOrderControllerFailed3(t *testing.T) {
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
	// InsertMockToDbFailed()
	// config.DB.Migrator().DropTable(models.Order{})

	testCases.name = "Id Group Product Not Found"
	body, error := json.Marshal(mock_data_payment)
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
	context.SetParamValues("5")

	middleware.JWT([]byte(constants.SECRET_JWT))(CreateOrderControllersTesting())(context)

	body_res := res.Body.String()
	fmt.Println("body", body_res)
	var responses GroupResponseSuccess
	err = json.Unmarshal([]byte(body_res), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}
func TestCreateOrderControllerFailed2(t *testing.T) {
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
	InsertMockToDbFailed()
	// config.DB.Migrator().DropTable(models.Order{})

	t.Run("Id Group Product Not Found", func(t *testing.T) {
		testCases.name = "Group Product Full"
		body, error := json.Marshal(mock_data_payment)
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
		context.SetParamValues("3")

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
}

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
		context.SetParamValues("2")

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
		context.SetParamValues("2")

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
		context.SetParamValues("2")

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
}

func TestGetOrderByIdOrderSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/orders/id/:id_order",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTUser()
	if err != nil {
		panic(err)
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_order")
	context.SetParamValues("1")

	middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdOrderControllersTesting())(context)

	body := res.Body.String()
	fmt.Println("body", body)
	var responses OrderResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}
func TestGetOrderByIdOrderFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/orders/id/:id_order",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	t.Run("access forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_failed, err := UsingJWTUser2()
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_failed))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_Order")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Invalid Id", func(t *testing.T) {
		testCases.name = "Invalid Id"
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("hsa")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("5")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Bad request", func(t *testing.T) {
		testCases.name = "Bad Request"
		config.DB.Migrator().DropTable(models.Payment{})
		config.DB.Migrator().DropTable(models.Order{})
		config.DB.Migrator().DropTable(models.GetGroupProduct{})
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}
func TestGetOrderByIdUsersSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/orders/users/:id_user",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTUser()
	if err != nil {
		panic(err)
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_user")
	context.SetParamValues("2")

	middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdUsersControllersTesting())(context)

	body := res.Body.String()
	fmt.Println("body", body)
	var responses OrderResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}
func TestGetOrderByIdusersFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/orders/users/:id_user",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	t.Run("access forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_failed, err := UsingJWTUser2()
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_failed))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_user")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdUsersControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Invalid Id", func(t *testing.T) {
		testCases.name = "Invalid Id"
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_user")
		context.SetParamValues("hsa")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdUsersControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_user")
		context.SetParamValues("5")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdUsersControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Bad request", func(t *testing.T) {
		testCases.name = "Bad Request"

		config.DB.Migrator().DropTable(models.Order{})
		config.DB.Migrator().DropTable(models.GroupProduct{})
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_user")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdUsersControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}
func TestGetOrderByIdGroupSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/orders/group/:id_group",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_group")
	context.SetParamValues("1")

	middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdGroupControllersTesting())(context)

	body := res.Body.String()
	fmt.Println("body", body)
	var responses OrderResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}
func TestGetOrderByIdGroupFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/orders/group/:id_group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	t.Run("access forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_failed, err := UsingJWTUser2()
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_failed))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdGroupControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Invalid Id", func(t *testing.T) {
		testCases.name = "Invalid Id"
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("hsa")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdGroupControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("5")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdGroupControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Bad request", func(t *testing.T) {
		testCases.name = "Bad Request"

		config.DB.Migrator().DropTable(models.Order{})
		config.DB.Migrator().DropTable(models.GroupProduct{})
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(GetOrderByIdGroupControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}

func TestUpdateOrderControllersSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/orders/update/:id_order",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	var (
		mock_data_detail = models.Detail{
			Email:    "netflix@mail.com",
			Password: "1234qwer",
		}
	)
	body_detail, error := json.Marshal(mock_data_detail)
	if error != nil {
		t.Error(t, error, "error marshal")
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	req := httptest.NewRequest(http.MethodPut, testCases.path, bytes.NewBuffer(body_detail))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_order")
	context.SetParamValues("1")

	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

	body := res.Body.String()
	fmt.Println("body", body)
	var responses OrderResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)
}

func TestUpdateOrderControllersFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/orders/update/:id_order",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	var (
		mock_data_detail = models.Detail{
			Email:    "netflix@mail.com",
			Password: "1234qwer",
		}
	)
	body_detail, error := json.Marshal(mock_data_detail)
	if error != nil {
		t.Error(t, error, "error marshal")
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	t.Run("Access Forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_fail, err := UsingJWTUser()
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, testCases.path, bytes.NewBuffer(body_detail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_fail))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Invalid Id", func(t *testing.T) {
		testCases.name = "Invalid Id"
		req := httptest.NewRequest(http.MethodPut, testCases.path, bytes.NewBuffer(body_detail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("jd")

		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Invalid Email", func(t *testing.T) {

		mock_data_detail = models.Detail{
			Email:    "",
			Password: "1234qwer",
		}
		body_detail, error = json.Marshal(mock_data_detail)
		if error != nil {
			t.Error(t, error, "error marshal")
		}
		testCases.name = "Invalid Email"
		req := httptest.NewRequest(http.MethodPut, testCases.path, bytes.NewBuffer(body_detail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Invalid Password", func(t *testing.T) {

		mock_data_detail = models.Detail{
			Email:    "net@mail.com",
			Password: "",
		}
		body_detail, error = json.Marshal(mock_data_detail)
		if error != nil {
			t.Error(t, error, "error marshal")
		}
		testCases.name = "Invalid Password"
		req := httptest.NewRequest(http.MethodPut, testCases.path, bytes.NewBuffer(body_detail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		mock_data_detail = models.Detail{
			Email:    "net@mail.com",
			Password: "1234qwer",
		}
		body_detail, error = json.Marshal(mock_data_detail)
		if error != nil {
			t.Error(t, error, "error marshal")
		}
		testCases.name = "Data Not Found"
		req := httptest.NewRequest(http.MethodGet, testCases.path, bytes.NewBuffer(body_detail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("4")

		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Bad Request", func(t *testing.T) {
		mock_data_detail = models.Detail{
			Email:    "net@mail.com",
			Password: "1234qwer",
		}
		body_detail, error = json.Marshal(mock_data_detail)
		if error != nil {
			t.Error(t, error, "error marshal")
		}
		testCases.name = "Bad Request"
		config.DB.Migrator().DropTable(models.Order{})
		req := httptest.NewRequest(http.MethodGet, testCases.path, bytes.NewBuffer(body_detail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("4")

		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateOrderControllersTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}

func TestDeleteOrderByIdGroupSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/orders/delete/:id_order",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_order")
	context.SetParamValues("1")

	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteOrderControllerTesting())(context)

	body := res.Body.String()
	fmt.Println("body", body)
	var responses OrderResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}
func TestDeleteOrderByIdGroupFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/orders/delete/:id_order",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	fmt.Println("cek token", token)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_order)
	config.DB.Save(&mock_data_xendit)

	t.Run("Invalid Id", func(t *testing.T) {
		testCases.name = "Invalid Id"

		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("sd")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteOrderControllerTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Access Forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_fail, err := UsingJWTUser()
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_fail))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteOrderControllerTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_order")
		context.SetParamValues("4")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteOrderControllerTesting())(context)

		body := res.Body.String()
		fmt.Println("body", body)
		var responses OrderResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}
