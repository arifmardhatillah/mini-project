package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Id_user string `json:"id_user" form:"id_user"`
	Judul   string `json:"judul" form:"judul"`
	Konten  string `json:"konten" form:"konten"`
}
