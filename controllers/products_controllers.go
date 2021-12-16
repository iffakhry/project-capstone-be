package controllers

import (
	"final-project/lib/databases"
	"final-project/middlewares"
	"final-project/models"
	response "final-project/responses"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var storageClient *storage.Client

type ValidatorProduct struct {
	Name_Product   string `validate:"required"`
	Detail_Product string `validate:"required"`
	Price          int    `validate:"required,gt=0"`
	Limit          int    `validate:"required,gt=0"`
	Photo          string `validate:"required"`
	Url            string `validate:"required"`
}

// controller untuk menambahkan product baru
func CreateProductControllers(c echo.Context) error {
	new_product := models.Products{}
	c.Bind(&new_product)
	bucket := "barengin-bucket"

	var err error

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Can't Connect"))
	}

	f, uploaded_file, err := c.Request().FormFile("photo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
	}

	defer f.Close()

	ext := strings.Split(uploaded_file.Filename, ".")
	extension := ext[len(ext)-1]
	check_extension := strings.ToLower(extension)
	if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("File Extension Not Allowed"))
	}

	if uploaded_file.Size == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Illegal File"))
	} else if uploaded_file.Size > 1050000 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Size File Too Big"))
	}
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	photo_name := strings.ReplaceAll(new_product.Name_Product, " ", "-")
	uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", photo_name, formatted, extension)
	new_product.Photo = uploaded_file.Filename
	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	new_product.Url = fmt.Sprintf("%v", u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
	}

	v := validator.New()
	validasi_product := ValidatorProduct{
		Name_Product:   new_product.Name_Product,
		Price:          new_product.Price,
		Detail_Product: new_product.Detail_Product,
		Limit:          new_product.Limit,
		Photo:          new_product.Photo,
		Url:            new_product.Url,
	}
	err = v.Struct(validasi_product)
	if err == nil {
		id_user_token, role := middlewares.ExtractTokenId(c)
		new_product.UsersID = uint(id_user_token)
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
		}
		// log.Println("role", role)
		_, err = databases.CreateProduct(&new_product)
	}
	if err != nil {
		// log.Println("error", err)
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menampilkan seluruh data product
func GetAllProductControllers(c echo.Context) error {
	product, err := databases.GetAllProduct()
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}

// controller untuk menampilkan data product by id
func GetProductByIdControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	product, e := databases.GetProductById(id)
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}

// controller untuk memperbarui data product by id
func UpdateProductControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	product, _ := databases.GetProductById(id)
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	_, role := middlewares.ExtractTokenId(c) // check token
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	update_product := models.Products{}
	c.Bind(&update_product)

	bucket := "barengin-bucket"

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Can't Connect"))
	}

	f, uploaded_file, err := c.Request().FormFile("photo")
	if uploaded_file != nil {
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
		}

		defer f.Close()

		ext := strings.Split(uploaded_file.Filename, ".")
		extension := ext[len(ext)-1]
		check_extension := strings.ToLower(extension)
		if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("File Extension Not Allowed"))
		}

		if uploaded_file.Size == 0 {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Illegal File"))
		} else if uploaded_file.Size > 1050000 {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Size File Too Big"))
		}
		t := time.Now()
		formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		photo_name := strings.ReplaceAll(update_product.Name_Product, " ", "-")
		uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", photo_name, formatted, extension)
		update_product.Photo = uploaded_file.Filename
		sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)

		if _, err := io.Copy(sw, f); err != nil {
			return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
		}

		if err := sw.Close(); err != nil {
			return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
		}

		u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
		update_product.Url = fmt.Sprintf("%v", u)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse("Failed to Upload File"))
		}
	} else {
		photo, url, _ := databases.GetPhotoUrlProductById(id)
		update_product.Photo = photo
		update_product.Url = url
	}
	v := validator.New()
	validasi_product := ValidatorProduct{
		Name_Product:   update_product.Name_Product,
		Price:          update_product.Price,
		Detail_Product: update_product.Detail_Product,
		Limit:          update_product.Limit,
		Photo:          update_product.Photo,
		Url:            update_product.Url,
	}
	err = v.Struct(validasi_product)
	if err == nil {
		id_user_token, role := middlewares.ExtractTokenId(c)
		update_product.UsersID = uint(id_user_token)
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
		}
		// log.Println("role", role)
		_, err = databases.UpdateProduct(id, &update_product)
	}
	if err != nil {
		// log.Println("error", err)
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menghapus data product by id
func DeleteProductControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	product, _ := databases.GetProductById(id)
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	_, role := middlewares.ExtractTokenId(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	databases.DeleteProduct(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
