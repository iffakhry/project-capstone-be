package controllers

import (
	"final-project/lib/databases"
	"final-project/middlewares"
	"final-project/models"
	response "final-project/responses"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	photo_name := strings.ReplaceAll(new_product.Name_Product, " ", "-")
	uploaded_file.Filename = fmt.Sprintf("%s-%s.%s", photo_name, formatted, extension)
	new_product.Photo = uploaded_file.Filename
	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	new_product.Url = fmt.Sprintf("%v", u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BadRequestResponse("Failed to Upload File"))
	}

	v := validator.New()
	validasi_homestay := ValidatorProduct{
		Name_Product:   new_product.Name_Product,
		Price:          new_product.Price,
		Detail_Product: new_product.Detail_Product,
		Photo:          new_product.Photo,
		Url:            new_product.Url,
	}
	err = v.Struct(validasi_homestay)
	if err == nil {
		id_token, role := middlewares.ExtractTokenId(c)
		if new_product.UsersID == uint(id_token) && role == "admin" {
			_, err = databases.CreateProduct(&new_product)
		} else {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
		}
	}
	if err != nil {
		log.Println("error", err)
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
