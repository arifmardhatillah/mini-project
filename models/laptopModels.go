package models

import "github.com/jinzhu/gorm"

type Laptop struct {
	gorm.Model
	Product
	Brand     string `json:"brand"`
	Processor string `json:"processor"`
	RAM       string `json:"ram"`
	Storage   string `json:"storage"`
}
