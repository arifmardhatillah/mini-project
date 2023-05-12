package config

import (
	"fmt"
	"project_structure/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "alta",
		"DB_Password": "root",
		"DB_Port":     "3306",
		"DB_Host":     "192.168.1.16",
		"DB_Name":     "new_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.Payment{}, &model.Shipping{})
}
