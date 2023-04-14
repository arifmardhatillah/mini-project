package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `json: "name" form:"name"`
	Description string  `json: "Description" form:"description"`
	Price       float64 `json: "Price" form:"price"`
	Stock       int     `json: "Stock" form:"stock"`
}
