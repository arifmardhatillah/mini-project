package controllers

import (
	"net/http"
	"prak/config"
	"prak/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var books models.Book
	if err := config.DB.First(&books, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "books not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "berhasil",
		"data_books": books,
	})

}

// create new book
func CreateBookController(c echo.Context) error {
	books := models.Book{}
	c.Bind(&books)

	if err := config.DB.Save(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new books",
		"books":   books,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var books models.Book
	if err := config.DB.Where("id = ?", id).First(&books).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "books not available",
		})
	}

	if err := config.DB.Delete(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil menghapus data",
		"books":  books,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	body := new(models.Book)

	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}

	var books models.Book
	if err := config.DB.Where("id = ?", id).First(&books).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "books not available",
		})
	}

	books.Judul = body.Judul
	books.Penulis = body.Penulis
	books.Penerbit = body.Penerbit

	if err := config.DB.Save(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil update data",
		"books":  books,
	})
}
