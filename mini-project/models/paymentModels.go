package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	Amount  float64 `json:"amount"`
	Methode string  `json:"methode"`
}
