package controller

import (
	"net/http"

	"project_structure/model"
	"project_structure/usecase"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := usecase.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}

func LoginAdminController(c echo.Context) error {
	admin := model.User{}
	c.Bind(&admin)

	err := usecase.LoginAdmin(&admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"admin":  admin,
	})
}
