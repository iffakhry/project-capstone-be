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
}

type Login struct {
	Email    string
	Password string
}

type ResponseAllData struct {
	Message string
	Data    []models.Products
}

type ResponseData struct {
	Message string
	Data    models.Products
}

type ResponseNonData struct {
	Message string
}

// data dummy
var (
	mockNewUser = models.Users{
		Name:     "sahril",
		Email:    "sahril@gmail.com",
		Password: "qwerty",
		Phone:    "+628123456789",
		Role:     "customer",
	}
	mockLoginUser = models.Users{
		Email:    "sahril@gmail.com",
		Password: "qwerty",
	}
)

// inisialisasi echo
func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

// menambahkan user
func InsertUser(users *models.Users) error {
	if err := config.DB.Create(&users).Error; err != nil {
		return err
	}
	return nil
}

// Create User Controller : Test Case 1 (All Input is valid)
func TestCreateUserController(t *testing.T) {
	e := InitEcho()

	body, err := json.Marshal(mockNewUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Success Operation", user.Message)
	}
	config.DB.Migrator().DropTable(&models.Users{})
}

// Create User Controller : Test Case 2 (Input invalid Name)
func TestCreateUserControllerFailed_1(t *testing.T) {
	var newDataUser = models.Users{
		Name:     "",
		Email:    "sahril@gmail.com",
		Password: "qwerty",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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
		assert.Equal(t, "Invalid Name", user.Message)
	}
}

// Create User Controller : Test Case 3 (Name contains non-alphanumeric)
func TestCreateUserControllerFailed_2(t *testing.T) {
	var newDataUser = models.Users{
		Name:     "sahril mahendra",
		Email:    "sahril@gmail.com",
		Password: "qwerty",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Name can only contains alphanumeric", user.Message)
	}
}

// Create User Controller : Test Case 4 (Invalid email)
func TestCreateUserControllerFailed_3(t *testing.T) {
	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "@gmail.com",
		Password: "qwerty",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Invalid Email", user.Message)
	}
}

// Create User Controller : Test Case 5 (Invalid Password)
func TestCreateUserControllerFailed_4(t *testing.T) {
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
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Invalid Password", user.Message)
	}
}

// Create User Controller : Test Case 6 (Password contain less than 6 characters)
func TestCreateUserControllerFailed_5(t *testing.T) {
	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "sahril2@gmail.com",
		Password: "123",
		Phone:    "+628123456789",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Password must consist of 6 characters or more", user.Message)
	}
}

// Create User Controller : Test Case 7 (Invalid telephone number)
func TestCreateUserControllerFailed_6(t *testing.T) {
	var newDataUser = models.Users{
		Name:     "sahrilmahendra",
		Email:    "sahril2@gmail.com",
		Password: "qwerty",
		Phone:    "0281",
	}

	e := InitEcho()

	body, err := json.Marshal(newDataUser)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Invalid Telephone Number", user.Message)
	}
}

// Create User Controller : Test Case 8 (Create user as admin at first)
func TestCreateUserControllerFailed_7(t *testing.T) {
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
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
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

		assert.Equal(t, "Success Operation", user.Message)
	}
}

// // Create User Controller : Test Case 9 (Email or Telephone Number Already Exist)
// func TestCreateUserControllerFailed_8(t *testing.T) {
// 	var newDataUser = models.Users{
// 		Name:     "sahriltwo",
// 		Email:    "sahril@gmail.com",
// 		Password: "123456",
// 		Phone:    "+628123456789",
// 	}
// 	e := InitEcho()

// 	InsertUser(&mockNewUser)

// 	body, err := json.Marshal(newDataUser)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}
// 	// send data using request body with HTTP Method POST
// 	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, CreateUserControllers(c)) {
// 		bodyrecponses := rec.Body.String()
// 		var user UserResponse

// 		err := json.Unmarshal([]byte(bodyrecponses), &user)
// 		if err != nil {
// 			assert.Error(t, err, "error")
// 		}

// 		assert.Equal(t, "Email or Telephone Number Already Exist", user.Message)
// 	}
// 	config.DB.Migrator().DropTable(&models.Users{})
// }

// // Login User Controller : Test Case 1 (Correct Email & Password)
// func TestLoginGetUserControllers(t *testing.T) {
// 	e := InitEcho()

// 	var newUser = mockNewUser
// 	newUser.Password, _ = helper.HashPassword(newUser.Password)
// 	InsertUser(&newUser)

// 	// hash, _ := helper.HashPassword(newUser.Password)
// 	body, error := json.Marshal(Login{Email: newUser.Email, Password: "qwerty"})
// 	if error != nil {
// 		t.Error(t, error, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/login")

// 	if assert.NoError(t, LoginUserControllers(c)) {
// 		bodyrecponses := rec.Body.String()
// 		var user UserResponse

// 		err := json.Unmarshal([]byte(bodyrecponses), &user)
// 		if err != nil {
// 			assert.Error(t, err, "error")
// 		}

// 		assert.Equal(t, "Login Success", user.Message)
// 	}
// }

// // Login User Controller : Test Case 2 (Correct Email & Incorrect Password)
// func TestLoginUserControllersFailed(t *testing.T) {
// 	e := InitEcho()

// 	body, error := json.Marshal(Login{Email: "sahril@gmail.com", Password: "qwert"})
// 	if error != nil {
// 		t.Error(t, error, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/login")

// 	if assert.NoError(t, LoginUserControllers(c)) {
// 		bodyrecponses := rec.Body.String()
// 		var user UserResponse

// 		err := json.Unmarshal([]byte(bodyrecponses), &user)
// 		if err != nil {
// 			assert.Error(t, err, "error")
// 		}

// 		assert.Equal(t, "Email or Password Incorrect", user.Message)
// 	}
// }

// // Get User Controller : Test Case 1
// func TestGetUserControllersSuccess(t *testing.T) {
// 	e := InitEcho()

// 	var newUser = mockNewUser

// 	fmt.Println("Password asli", newUser.Password)

// 	newUser.Password, _ = helper.HashPassword(newUser.Password)
// 	InsertUser(&newUser)

// 	var userDB models.Users
// 	tx := config.DB.Where("email = ?", newUser.Email).First(&userDB)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}

// 	fmt.Println("Password login", mockNewUser.Password)

// 	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(mockNewUser.Password))
// 	if err != nil {
// 		panic(err)
// 	}

// 	token, err := middlewares.CreateToken(int(userDB.ID), userDB.Role)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	context := e.NewContext(req, rec)
// 	context.SetPath("/jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues(fmt.Sprint(int(userDB.ID)))
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserControllersTesting())(context)

// 	var user UserResponse
// 	rec_body := rec.Body.String()
// 	json.Unmarshal([]byte(rec_body), &user)
// 	if err != nil {
// 		assert.Error(t, err, "error")
// 	}

// 	t.Run("GET /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Success Operation", user.Message)
// 	})
// }

// // Get All Users Controller : Test Case 1
// func TestGetAllUsersControllersSuccess(t *testing.T) {
// 	e := InitEcho()

// 	var mockNewUser = mockNewUser
// 	mockNewUser.Password, _ = helper.HashPassword(mockNewUser.Password)
// 	InsertUser(&mockNewUser)

// 	var mock_data_admin = models.Users{
// 		Name:     "admin",
// 		Email:    "admin@admin.com",
// 		Password: "qwerty",
// 		Phone:    "+6281111111111",
// 	}
// 	mock_data_admin.Role = "admin"
// 	mock_data_admin.Password, _ = helper.HashPassword(mock_data_admin.Password)
// 	InsertUser(&mock_data_admin)

// 	var userDB models.Users
// 	tx := config.DB.Where("email = ?", "admin@admin.com").First(&userDB)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte("qwerty"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	token, err := middlewares.CreateToken(int(userDB.ID), userDB.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("token ada")
// 	req := httptest.NewRequest(http.MethodGet, "/jwt/users", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	context := e.NewContext(req, rec)
// 	context.SetPath("/jwt/users")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllUsersControllersTesting())(context)

// 	var user UserResponse
// 	res_body := rec.Body.String()
// 	err = json.Unmarshal([]byte(res_body), &user)
// 	if err != nil {
// 		assert.Error(t, err, "error")
// 	}

// 	t.Run("GET /jwt/users", func(t *testing.T) {
// 		assert.Equal(t, "Success Operation", user.Message)
// 	})
// }

// // Create User Controller : Test Case 1 (Success Operation)
// func TestUpdateUserControllerSuccess(t *testing.T) {
// 	e := InitEcho()

// 	var newUser = mockNewUser
// 	newUser.Password, _ = helper.HashPassword(newUser.Password)
// 	InsertUser(&newUser)

// 	var userDB models.Users
// 	tx := config.DB.Where("email = ?", "sahril@gmail.com").First(&userDB)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte("qwerty"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	token, err := middlewares.CreateToken(int(userDB.ID), userDB.Role)
// 	if err != nil {
// 		panic(err)
// 	}

// 	mockUpdateData := models.Users{
// 		Name:     "sahrilUpdate",
// 		Email:    "sahrilUpdate@gmail.com",
// 		Password: "qwertyupdate",
// 		Phone:    "+6281234567892",
// 	}

// 	body, err := json.Marshal(mockUpdateData)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	// send data using request body with HTTP Method POST
// 	req := httptest.NewRequest(http.MethodPut, "/jwt/users/:id", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	context := e.NewContext(req, rec)
// 	context.SetPath("/jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues(fmt.Sprint(int(userDB.ID)))
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	var user UserResponse
// 	res_body := rec.Body.String()
// 	err = json.Unmarshal([]byte(res_body), &user)
// 	if err != nil {
// 		assert.Error(t, err, "error")
// 	}

// 	t.Run("PUT /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Success Operation", user.Message)
// 	})
// }

// // sahril
// var (
// 	mock_data_yuser = models.Users{
// 		Name:     "user",
// 		Email:    "user@example.com",
// 		Password: "user123",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust = models.Users{
// 		Name:     "userdua",
// 		Email:    "user@user.com",
// 		Password: "user123",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_copy = models.Users{
// 		Name:     "admin",
// 		Email:    "user@example.com",
// 		Password: "user123",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust_invalid_name = models.Users{
// 		Name:     "",
// 		Email:    "user@user.com",
// 		Password: "user123",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust_invalid_name_alphanum = models.Users{
// 		Name:     "sah 12",
// 		Email:    "user@user.com",
// 		Password: "user123",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust_invalid_email = models.Users{
// 		Name:     "userdua",
// 		Email:    "useruser",
// 		Password: "user123",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust_invalid_password = models.Users{
// 		Name:     "userdua",
// 		Email:    "user@user.com",
// 		Password: "",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust_invalid_password_less6 = models.Users{
// 		Name:     "userdua",
// 		Email:    "user@user.com",
// 		Password: "s",
// 		Phone:    "+6213456721",
// 		Role:     "customer",
// 	}
// 	mock_data_cust_invalid_phone = models.Users{
// 		Name:     "userdua",
// 		Email:    "user@user.com",
// 		Password: "user123",
// 		Phone:    "13456721",
// 		Role:     "customer",
// 	}
// 	mock_data_yuser_admin = models.Users{
// 		Name:     "adminbarengin",
// 		Email:    "admin@admin.com",
// 		Password: "admin123",
// 		Phone:    "+6213456721",
// 		Role:     "admin",
// 	}
// 	mock_data_product_yuser = models.Products{
// 		Name_Product:   "product yuser",
// 		Detail_Product: "lorem ipsum",
// 		Price:          10000,
// 		Limit:          5,
// 		Photo:          "dummy.jpg",
// 		Url:            "https://storage.googleapis.com/barengin-bucket/dummy.jpg",
// 		UsersID:        1,
// 	}
// 	mock_data_group_product_yuser = models.GroupProduct{
// 		UsersID:              1,
// 		ProductsID:           1,
// 		NameGroupProduct:     "group dummy",
// 		CapacityGroupProduct: 5,
// 		AdminFee:             5000,
// 		TotalPrice:           55000,
// 		DurationGroup:        "30-12-2021",
// 		Status:               "Available",
// 	}
// 	mock_data_order_yuser = models.Order{
// 		UsersID:        1,
// 		GroupProductID: 1,
// 		PriceOrder:     500,
// 		NameProduct:    "dummy",
// 	}
// )

// func TestGetAllUserFailureAccessForbidden(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser.Email, mock_data_yuser.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllUsersControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseAllData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users", func(t *testing.T) {
// 		assert.Equal(t, "Access Forbidden", users.Message)
// 	})
// }

// func TestGetAllUserFailureBadRequest(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	config.DB.Migrator().DropTable(models.Users{})
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllUsersControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseAllData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users", func(t *testing.T) {
// 		assert.Equal(t, "Bad Request", users.Message)
// 	})
// }

// func TestGetAllUserFailureDataNotFound(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	databases.DeleteUser(1)
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetAllUsersControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseAllData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users", func(t *testing.T) {
// 		assert.Equal(t, "Data Not Found", users.Message)
// 	})
// }

// func TestGetUserFailureAccessForbidden(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser)
// 	config.DB.Save(&mock_data_cust)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser.Email, mock_data_yuser.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Access Forbidden", users.Message)
// 	})
// }
// func TestGetUserFailureInvalidID(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("$")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Id", users.Message)
// 	})
// }
// func TestGetUserFailureBadRequest(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	config.DB.Migrator().DropTable(models.Users{})
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Bad Request", users.Message)
// 	})
// }

// func TestGetUserFailureDataNotFound(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("GET /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Data Not Found", users.Message)
// 	})
// }

// func TestDeleteUserSuccess(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Success Operation", users.Message)
// 	})
// }
// func TestDeleteUserFailureAccessForbidden(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser)
// 	config.DB.Save(&mock_data_cust)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser.Email, mock_data_yuser.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Access Forbidden", users.Message)
// 	})
// }
// func TestDeleteUserFailureInvalidID(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("$")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Id", users.Message)
// 	})
// }

// func TestDeleteUserFailureProductUsed(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_product_yuser)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Access is denied ID data is in the product", users.Message)
// 	})
// }
// func TestDeleteUserFailureDataNotFound(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)
// 	body := res.Body.String()
// 	var users ResponseNonData
// 	json.Unmarshal([]byte(body), &users)
// 	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
// 		assert.Equal(t, "Data Not Found", users.Message)
// 	})
// }

// func TestUpdateUserControllersFailureAccessForbidden(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser)
// 	config.DB.Save(&mock_data_yuser_admin)

// 	body, err := json.Marshal(mock_data_cust)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser.Email, mock_data_yuser.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Access Forbidden", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidId(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("$")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Id", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureDataNotFound(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("5")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Data Not Found", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidName(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust_invalid_name)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Name", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidNameAlphaNum(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust_invalid_name_alphanum)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Name can only contains alphanumeric", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidEmail(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust_invalid_email)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Email", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidPassword(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust_invalid_password)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Password", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidPasswordLess6(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust_invalid_password_less6)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Password must consist of 6 characters or more", product.Message)
// 	})
// }
// func TestUpdateUserControllersFailureInvalidPhone(t *testing.T) {
// 	e := InitEcho()
// 	config.DB.Save(&mock_data_yuser_admin)
// 	config.DB.Save(&mock_data_yuser)

// 	body, err := json.Marshal(mock_data_cust_invalid_phone)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	var user_db models.Users
// 	tx := config.DB.Where("email = ? AND password = ?", mock_data_yuser_admin.Email, mock_data_yuser_admin.Password).First(&user_db)
// 	if tx.Error != nil {
// 		t.Error(tx.Error)
// 	}
// 	token, err := middlewares.CreateToken(int(user_db.ID), user_db.Role)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("jwt/users/:id")
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

// 	body_response := res.Body.String()
// 	var product ResponseNonData
// 	json.Unmarshal([]byte(body_response), &product)
// 	t.Run("update user", func(t *testing.T) {
// 		assert.Equal(t, "Invalid Telephone Number", product.Message)
// 	})
// }
