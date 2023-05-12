package controller

import (
	"net/http"
	"project_structure/middleware"
	"project_structure/model/payload"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetShippingController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only user can access",
			"error":   err.Error(),
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	shipping, err := usecase.GetShipping(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get shipping",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"shippings": shipping,
	})
}

func CreateShippingController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access",
			"error":   err.Error(),
		})
	}

	payload := payload.ShippingRequest{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	shipping, err := usecase.CreateShipping(payload.UserID, payload.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":  "success add shipping",
		"shipping": shipping,
	})
}
