package controller

import (
	"net/http"
	"project_structure/middleware"
	"project_structure/model/payload"
	"project_structure/usecase"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := usecase.GetListUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// create user
func CreateUserController(c echo.Context) error {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	// validasi request body
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create user data",
			"errorDescription": err,
		})
	}
	resp, err := usecase.CreateUser(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages":         "error create user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    resp,
	})
}

// creat admin
func CreateAdminController(c echo.Context) error {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	// validasi request body
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create admin data",
			"errorDescription": err,
		})
	}
	resp, err := usecase.CreateAdmin(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages":         "error create admin",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new admin",
		"data":    resp,
	})
}

func TopUpSaldoController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}
	payload := payload.TopUpRequest{}
	c.Bind(&payload)

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}

	user, err := usecase.TopUp(uint(userID), payload.Saldo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Top Up Balance successful",
		"user":    user,
	})
}
