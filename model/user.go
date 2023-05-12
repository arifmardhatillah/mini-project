package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Saldo    uint   `json:"saldo"`
	Role     string `json:"role"`
	Token    string `gorm:"-"`
}
