package controllers

import (
	"encoding/json"
	"final-project/config"
	"final-project/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GroupResponseSuccess struct {
	Message string
	Data    []models.GroupProduct
}

var (
	mock_data_user1 = models.Users{
		Name:     "user1",
		Email:    "user1@mail.com",
		Password: "user123",
		Phone:    "+628257237462",
	}
	mock_data_user2 = models.Users{
		Name:     "user2",
		Email:    "user2@mail.com",
		Password: "user123",
		Phone:    "+628257327462",
	}
	mock_data_product = models.Products{
		Name_Product:   "Netflix",
		Detail_Product: "lorem",
		Price:          200000,
		Limit:          5,
		Photo:          "netflix.jpg",
	}
	mock_data_group = models.GroupProduct{
		UsersID:              1,
		ProductsID:           1,
		NameGroupProduct:     "netflix-1",
		CapacityGroupProduct: 1,
		AdminFee:             5000,
		TotalPrice:           250000,
		Status:               "Available",
	}
	mock_data_group2 = models.GroupProduct{
		UsersID:              2,
		ProductsID:           1,
		NameGroupProduct:     "netflix-2",
		CapacityGroupProduct: 1,
		AdminFee:             5000,
		TotalPrice:           250000,
		Status:               "Available",
	}
)

func InsertMockToDb() {
	config.DB.Save(&mock_data_user1)
	config.DB.Save(&mock_data_user2)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_group)
	config.DB.Save(&mock_data_group2)
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
	req := httptest.NewRequest(http.MethodGet, "/products/group", nil)
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

func TestGetAllGroupControllerFailed2(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Bad Request",
		path:       "/products/group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	config.DB.Migrator().DropTable(models.GroupProduct{})
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)

	// for index, testCase := range testCases {
	// context.SetPath(testCase.path)

	if assert.NoError(t, GetAllGroupProductControllers(context)) {
		body := res.Body.String()
		var responses GroupResponseSuccess
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
		// assert.Equal(t, "netflix-1", responses.Data[index].NameGroupProduct)
		// }
	}
}
func TestGetAllGroupControllerFailed1(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Data Not Found",
		path:       "/products/group",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// InsertMockToDb()
	// config.DB.Migrator().DropTable(models.GroupProduct{})
	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)

	// for index, testCase := range testCases {
	// context.SetPath(testCase.path)

	if assert.NoError(t, GetAllGroupProductControllers(context)) {
		body := res.Body.String()
		var responses GroupResponseSuccess
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
		// assert.Equal(t, "netflix-1", responses.Data[index].NameGroupProduct)
		// }
	}
}
