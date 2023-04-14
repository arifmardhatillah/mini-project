package controllers

import (
	"net/http"
	"prak/config"
	"prak/models"

	"github.com/labstack/echo/v4"
)

func GetLaptopController(c echo.Context) error {
	kucings := []models.Laptop{}

	err := config.DB.Find(&kucings).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "gagal mendapatkan data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Mendapatkan data",
		"data":    kucings,
	})

}

func PostLaptopController(c echo.Context) error {
	kucings := models.Laptop{}

	c.Bind(&kucings)

	err := config.DB.Create(&kucings).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Gagal melakukan decode",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil",
		"data":   kucings,
	})
}
