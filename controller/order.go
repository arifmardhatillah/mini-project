package controller

import (
	"net/http"
	"project_structure/middleware"
	"project_structure/model/payload"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetOrdersController(c echo.Context) error {
	orders, e := usecase.GetListOrders()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"orders": orders,
	})
}

func GetOrderController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access",
			"error":   err.Error(),
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	order, err := usecase.GetOrder(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get order",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"orders": order,
	})
}

// create new order
func CreateOrderController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only user can access",
			"error":   err.Error(),
		})
	}

	payload := payload.OrderRequest{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	order, err := usecase.CreateOrder(payload.UserID, payload.ProductID, payload.Name, payload.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add order",
		"order":   order,
	})
}
