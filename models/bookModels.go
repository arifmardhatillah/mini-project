package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}
