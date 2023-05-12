package controller

import (
	"net/http"
	"project_structure/middleware"
	"project_structure/model"
	"project_structure/model/payload"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	products, e := usecase.GetListProducts()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}

func GetProductController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	product, err := usecase.GetProduct(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get product",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": product,
	})
}

// create new product
func CreateProductController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access",
			"error":   err.Error(),
		})
	}

	payload := payload.ProductRequest{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	product, err := usecase.CreateProduct(payload.Name, payload.Description, payload.Price, payload.Stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product",
		"product": product,
	})
}

// delete product by id
func DeleteProductController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete product",
			"errorDescription": err,
			"errorMessage":     "Sorry, the product cannot be deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete product",
	})
}

// update product by id
func UpdateProductController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	product := model.Product{}
	c.Bind(&product)
	product.ID = uint(id)

	updateStockRequest := payload.UpdateStockRequest{}
	if err := c.Bind(&updateStockRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update product data",
			"errorDescription": err,
			"errorMessage":     "Sorry, the product data cannot be changed",
		})
	}

	if err := usecase.UpdateProduct(&product, updateStockRequest.Stock); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update product",
			"errorDescription": err,
			"errorMessage":     "Sorry, the product cannot be changed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product",
	})
}
