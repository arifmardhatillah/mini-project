package controllers

import (
	"net/http"
	"prak/config"
	"prak/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get Blog by id
func GetBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var blogs models.Blog
	if err := config.DB.First(&blogs, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "blogs not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "berhasil",
		"data_blogs": blogs,
	})

}

// create new Blog
func CreateBlogController(c echo.Context) error {
	blogs := models.Blog{}
	c.Bind(&blogs)

	if err := config.DB.Save(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blogs",
		"blogs":   blogs,
	})
}

// delete Blog by id
func DeleteBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var blogs models.Blog
	if err := config.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "blogs not available",
		})
	}

	if err := config.DB.Delete(&blogs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil menghapus data",
		"blogs":  blogs,
	})
}

// update Blog by id
func UpdateBlogController(c echo.Context) error {
	body := new(models.Blog)

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

	var blogs models.Blog
	if err := config.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "blogs not available",
		})
	}

	blogs.Id_user = body.Id_user
	blogs.Judul = body.Judul
	blogs.Konten = body.Konten

	if err := config.DB.Save(&blogs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil update data",
		"blogs":  blogs,
	})
}
