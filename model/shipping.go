package model

import "gorm.io/gorm"

type Shipping struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Address string `json:"address"`
	User    User   `json:"user"`
}
