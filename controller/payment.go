package controller

import (
	"net/http"
	"project_structure/middleware"
	"project_structure/model"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetPaymentController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access",
			"error":   err.Error(),
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	payment, err := usecase.GetPayment(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get payment",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"payments": payment,
	})
}

func CreatePaymentController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access",
			"error":   err.Error(),
		})
	}

	payment := model.Payment{}
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create payment data",
			"errorDescription": err,
		})
	}

	if err := usecase.CreatePayment(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create payment",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new payment",
		"payment": payment,
	})
}
