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

	//route group product tanpa JWT
	e.GET("/products/group", controllers.GetAllGroupProductControllers)
	e.GET("/products/group/:id_group", controllers.GetByIdGroupProductControllers)
	e.GET("/products/group/status/:status", controllers.GetAvailableGroupProductControllers)
	e.GET("/products/group/products/:id_products", controllers.GetByIdProductsGroupProductControllers)

	// route product tanpa JWT
	e.GET("/products", controllers.GetAllProductControllers)
	e.GET("/products/:id", controllers.GetProductByIdControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// group product JWT
	j.POST("/products/group/:id_products", controllers.CreateGroupProductControllers)
	j.DELETE("/products/group/delete/:id_group", controllers.DeleteGroupProductControllers)

	// route users dengan JWT
	j.GET("/users", controllers.GetAllUsersControllers)       // admin
	j.GET("/users/:id", controllers.GetUserControllers)       // admin dan pemilik akun
	j.PUT("/users/:id", controllers.UpdateUserControllers)    // admin dan pemilik akun
	j.DELETE("/users/:id", controllers.DeleteUserControllers) // admin dan pemilik akun

	// route product dengan JWT
	j.POST("/products", controllers.CreateProductControllers)       // admin
	j.PUT("/products/:id", controllers.UpdateProductControllers)    // admin
	j.DELETE("/products/:id", controllers.DeleteProductControllers) // admin

	//route order
	j.POST("/orders/:id_group", controllers.CreateOrderControllers)
	j.GET("/orders/id/:id_order", controllers.GetOrderByIdOrderControllers)
	j.GET("/orders/group/:id_group", controllers.GetOrderByIdGroupControllers)
	j.GET("/orders/users/:id_user", controllers.GetOrderByIdUsersControllers)
	j.PUT("/orders/update/:id_order", controllers.UpdateOrderControllers)    //admin
	j.DELETE("/orders/delete/:id_order", controllers.DeleteOrderControllers) //admin
	return e
}
