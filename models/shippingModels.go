package models

import "github.com/jinzhu/gorm"

type Shipping struct {
	gorm.Model
	Address string `json:"address"`
}
