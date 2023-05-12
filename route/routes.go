package route

import (
	"net/http"
	"project_structure/constant"
	"project_structure/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/register/user", controller.CreateUserController)
	e.POST("/login/user", controller.LoginUserController)
	e.POST("/register/admin", controller.CreateAdminController)
	e.POST("/login/admin", controller.LoginAdminController)
	e.GET("/products", controller.GetProductsController)

	// user collection
	user := e.Group("/user", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.POST("/payments", controller.CreatePaymentController)
	user.GET("/products/:id", controller.GetProductController)
	user.POST("/orders", controller.CreateOrderController)
	user.POST("/topup", controller.TopUpSaldoController)
	user.GET("/orders/:id", controller.GetOrderController)
	user.GET("/shippings/:id", controller.GetShippingController)

	// admin collection
	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.GET("/users", controller.GetUsersController)
	user.GET("/payments/:id", controller.GetPaymentController)
	admin.POST("/products", controller.CreateProductController)
	admin.PUT("/products/:id", controller.UpdateProductController)
	admin.DELETE("/products/:id", controller.DeleteProductController)
	admin.GET("/orders", controller.GetOrdersController)
	admin.POST("/shippings", controller.CreateShippingController)
}
