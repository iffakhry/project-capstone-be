package controllers

import (
	"encoding/json"
	"final-project/config"
	"final-project/constants"
	"final-project/middlewares"
	"final-project/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type GroupResponseSuccess struct {
	Message string
	Data    []models.GroupProduct
}
type OrderResponseSuccess struct {
	Message string
	Data    models.Payment
}

var (
	mock_data_login_admin = models.Users{
		Email:    "admin@admin.com",
		Password: "admin123",
	}
	mock_data_login_user1 = models.Users{
		Email:    "user1@mail.com",
		Password: "user123",
	}
	mock_data_admin = models.Users{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "admin123",
		Phone:    "+111111111111",
		Role:     "admin",
	}
	mock_data_user1 = models.Users{
		Name:     "user1",
		Email:    "user1@mail.com",
		Password: "user123",
		Phone:    "+628257237412",
		Role:     "customer",
	}
	mock_data_user2 = models.Users{
		Name:     "user2",
		Email:    "user2@mail.com",
		Password: "user123",
		Phone:    "+628257327462",
		Role:     "customer",
	}
	mock_data_product = models.Products{
		Name_Product:   "Netflix",
		Detail_Product: "lorem",
		Price:          200000,
		Limit:          5,
		Photo:          "netflix.jpg",
	}
	mock_data_product2 = models.Products{
		Name_Product:   "",
		Detail_Product: "lorem",
		Price:          0,
		Limit:          5,
		Photo:          "netflix.jpg",
	}
	mock_data_group = models.GroupProduct{
		UsersID:              2,
		ProductsID:           1,
		NameGroupProduct:     "netflix-1",
		CapacityGroupProduct: 1,
		AdminFee:             5000,
		TotalPrice:           250000,
		Status:               "Available",
	}
	mock_data_group2 = models.GroupProduct{
		UsersID:              3,
		ProductsID:           1,
		NameGroupProduct:     "netflix-2",
		CapacityGroupProduct: 1,
		AdminFee:             5000,
		TotalPrice:           250000,
		Status:               "Available",
	}
	mock_data_order = models.Order{
		UsersID:        1,
		GroupProductID: 1,
		PriceOrder:     45000,
	}
	mock_data_order2 = models.Order{
		UsersID:        1,
		GroupProductID: 3,
		PriceOrder:     45000,
	}
	mock_data_xendit = models.Payment{
		OrderID:     1,
		Amount:      45000,
		EwalletType: "OVO",
		ExternalId:  "1982773",
	}
)

func InsertMockToDb() {
	config.DB.Save(&mock_data_admin)
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_user2)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_group2)
	config.DB.Save(&mock_data_xendit)
	config.DB.Save(&mock_data_order)
}

// Fungsi untuk melakukan login dan ekstraksi token JWT
func UsingJWTAdmin() (string, error) {
	// Melakukan login data user test
	InsertMockToDb()
	var user models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login_admin.Email, mock_data_login_admin.Password).First(&user)
	if tx.Error != nil {
		return "", tx.Error
	}
	// Mengektraksi token data user test
	token, err := middlewares.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
func UsingJWTUser() (string, error) {
	// Melakukan login data user test
	InsertMockToDb()
	var user models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login_user1.Email, mock_data_login_user1.Password).First(&user)
	if tx.Error != nil {
		return "", tx.Error
	}
	// Mengektraksi token data user test
	token, err := middlewares.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func TestCreateGroupControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/products/group/:id_products",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTUser()
	if err != nil {
		panic(err)
	}

	// InsertMockToDb()
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)

	req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_products")
	context.SetParamValues("1")

	middleware.JWT([]byte(constants.SECRET_JWT))(CreateGroupProductControllersTesting())(context)

	body := res.Body.String()
	var responses GroupResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)
}

func TestCreateGroupControllerFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/products/group/:id_products",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTUser()
	if err != nil {
		panic(err)
	}

	InsertMockToDb()

	t.Run("tescase_create_access_forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_user, err := UsingJWTAdmin()
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_user))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("tescase_create_invalid_id", func(t *testing.T) {
		testCases.name = "Invalid Id"

		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("e")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("tescase_create_data_not_found", func(t *testing.T) {
		testCases.name = "Id Product Not Found"
		config.DB.Save(&mock_data_product2)
		// config.DB.Save(models.GroupProduct{})
		// config.DB.Migrator().DropTable(models.Products{})
		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("2")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("tescase_create_bad_request", func(t *testing.T) {
		testCases.name = "Bad Request"

		config.DB.Migrator().DropTable(models.GroupProduct{})

		req := httptest.NewRequest(http.MethodPost, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(CreateGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}

func TestGetByIdGroupControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/products/group/:id_group",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	InsertMockToDb()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_group")
	context.SetParamValues("1")

	if assert.NoError(t, GetByIdGroupProductControllers(context)) {
		body := res.Body.String()
		var responses GroupResponseSuccess
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}
}
func TestGetByIdGroupControllerFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/products/group/:id_group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	t.Run("tescase_data_not_found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("3")

		if assert.NoError(t, GetByIdGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)

		}
	})
	t.Run("tescase_bad_request", func(t *testing.T) {
		testCases.name = "Bad Request"

		InsertMockToDb()
		config.DB.Migrator().DropTable(models.GroupProduct{})

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		if assert.NoError(t, GetByIdGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {

				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
	t.Run("tescase_Invalid_param", func(t *testing.T) {
		testCases.name = "Invalid Id"

		InsertMockToDb()
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("a")

		if assert.NoError(t, GetByIdGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
}

// get group by id product success
func TestGetByIdProductGroupControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/products/group/products/:id_products",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	InsertMockToDb()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_products")
	context.SetParamValues("1")

	if assert.NoError(t, GetByIdProductsGroupProductControllers(context)) {
		body := res.Body.String()
		var responses GroupResponseSuccess
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}
}

// get group by id product failed
func TestGetByIdProductGroupControllerFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/products/group/products/:id_products",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	t.Run("tescase_data_not_found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("3")

		if assert.NoError(t, GetByIdProductsGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)

		}
	})
	t.Run("tescase_bad_request", func(t *testing.T) {
		testCases.name = "Bad Request"

		InsertMockToDb()
		config.DB.Migrator().DropTable(models.GroupProduct{})

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("1")

		if assert.NoError(t, GetByIdProductsGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {

				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
	t.Run("tescase_Invalid_param", func(t *testing.T) {
		testCases.name = "Invalid Id"

		InsertMockToDb()
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_products")
		context.SetParamValues("a")

		if assert.NoError(t, GetByIdProductsGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
}

func TestGetAllGroupControllerSuccess(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "Success Operation",
			path:       "/products/group",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	InsertMockToDb()
	req := httptest.NewRequest(http.MethodGet, testCases[0].path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)

	for index, testCase := range testCases {
		context.SetPath(testCase.path)

		if assert.NoError(t, GetAllGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases[index].expectCode, res.Code)
			assert.Equal(t, testCases[index].name, responses.Message)
			assert.Equal(t, "netflix-1", responses.Data[index].NameGroupProduct)
		}
	}
}

//get all group failed
func TestGetAllGroupControllerFailed2(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/products/group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	t.Run("get_all_data_not_found", func(t *testing.T) {
		testCases.name = "Data Not Found"
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
	t.Run("get_all_group_bad_request", func(t *testing.T) {
		testCases.name = "Bad Request"
		config.DB.Migrator().DropTable(models.GroupProduct{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
}

// get group by status failed
func TestGetByStatusGroupControllerFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/products/group/status/:status",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	t.Run("tescase_data_not_found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("status")
		context.SetParamValues("available")

		if assert.NoError(t, GetAvailableGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)

		}
	})
	t.Run("tescase_bad_request", func(t *testing.T) {
		testCases.name = "Bad Request"

		InsertMockToDb()
		config.DB.Migrator().DropTable(models.GroupProduct{})

		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("status")
		context.SetParamValues("available")

		if assert.NoError(t, GetAvailableGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {

				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
	t.Run("tescase_Invalid_param", func(t *testing.T) {
		testCases.name = "Invalid Param"

		InsertMockToDb()
		req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("status")
		context.SetParamValues("1")

		if assert.NoError(t, GetAvailableGroupProductControllers(context)) {
			body := res.Body.String()
			var responses GroupResponseSuccess
			err := json.Unmarshal([]byte(body), &responses)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCases.expectCode, res.Code)
			assert.Equal(t, testCases.name, responses.Message)
		}
	})
}

// Get by status success
func TestGetByStatusGroupControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/products/group/status/:status",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	InsertMockToDb()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("status")
	context.SetParamValues("available")

	if assert.NoError(t, GetAvailableGroupProductControllers(context)) {
		body := res.Body.String()
		var responses GroupResponseSuccess
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}
}

// Delete group success
func TestDeleteGroupControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/products/group/delete/:id_group",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	InsertMockToDb()
	config.DB.Migrator().DropTable(models.Order{})

	req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_group")
	context.SetParamValues("1")

	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteGroupProductControllersTesting())(context)

	body := res.Body.String()
	var responses GroupResponseSuccess
	err = json.Unmarshal([]byte(body), &responses)
	if err != nil {
		assert.Error(t, err, "error")
	}
	assert.Equal(t, testCases.expectCode, res.Code)
	assert.Equal(t, testCases.name, responses.Message)

}

func TestDeleteGroupControllerFailed(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "",
		path:       "/products/group/delete/:id_group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTAdmin()
	if err != nil {
		panic(err)
	}

	InsertMockToDb()

	t.Run("tescase_delete_access_forbidden", func(t *testing.T) {
		testCases.name = "Access Forbidden"
		token_user, err := UsingJWTUser()
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token_user))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("tescase_delete_invalid_id", func(t *testing.T) {
		testCases.name = "Invalid Id"

		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("e")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("tescase_delete_access_is_denied", func(t *testing.T) {
		testCases.name = "Access is denied ID data is in the order"

		config.DB.Save(&mock_data_order)
		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})
	t.Run("tescase_delete_data_not_found", func(t *testing.T) {
		testCases.name = "Data Not Found"

		config.DB.Migrator().DropTable(models.Order{})
		config.DB.Migrator().DropTable(models.GroupProduct{})
		req := httptest.NewRequest(http.MethodDelete, testCases.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath(testCases.path)
		context.SetParamNames("id_group")
		context.SetParamValues("1")

		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteGroupProductControllersTesting())(context)

		body := res.Body.String()
		var responses GroupResponseSuccess
		err = json.Unmarshal([]byte(body), &responses)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	})

}
