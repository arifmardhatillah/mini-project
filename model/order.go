package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Product   Product `json:"product"`
}
