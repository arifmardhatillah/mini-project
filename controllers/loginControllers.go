package controllers

import (
	"net/http"
	"prak/lib/database"
	"prak/models"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
