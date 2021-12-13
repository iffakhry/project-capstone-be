package routes

import (
	"final-project/constants"
	"final-project/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// function routes
func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// route users tanpa JWT
	e.POST("/signup", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUserControllers)
	e.GET("/users/:id", controllers.GetUserControllers)
	e.GET("/users", controllers.GetAllUsersControllers)

	//route group product tanpa JWT
	e.GET("/products/group", controllers.GetAllGroupProductControllers)
	e.GET("/products/group/:id", controllers.GetGroupProductControllers)
	e.GET("/products/group/status/:status", controllers.GetAvailableGroupProductControllers)

	// route product tanpa JWT
	e.GET("/products", controllers.GetAllProductControllers)
	e.GET("/products/:id", controllers.GetProductByIdControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	j.POST("/products/group", controllers.CreateGroupProductControllers)

	// route users dengan JWT
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	// route product dengan JWT
	j.POST("/products", controllers.CreateProductControllers)
	j.PUT("/products/:id", controllers.UpdateProductControllers)
	j.DELETE("/products/:id", controllers.DeleteProductControllers)
	return e
}
