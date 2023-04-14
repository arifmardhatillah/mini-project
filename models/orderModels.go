package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Product  Product
	Quantity string `json:"quantity"`
}
