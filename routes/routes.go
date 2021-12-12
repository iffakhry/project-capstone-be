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
	e.GET("/product/group/:id", controllers.GetGroupProductController)
	e.GET("/product/group", controllers.GetAllGroupProductControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	j.POST("/product/group", controllers.CreateGroupProductControllers)

	// route users dengan JWT
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	return e
}
