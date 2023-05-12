package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateProduct(product *model.Product) error {
	if err := config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProducts() (products []model.Product, err error) {
	if err = config.DB.Find(&products).Error; err != nil {
		return
	}
	return
}

func GetProduct(id uint) (product model.Product, err error) {
	product.ID = id
	if err = config.DB.First(&product).Error; err != nil {
		return
	}
	return
}

func UpdateProduct(product *model.Product) error {
	if err := config.DB.Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *model.Product) error {
	if err := config.DB.Delete(product).Error; err != nil {
		return err
	}
	return nil
}
