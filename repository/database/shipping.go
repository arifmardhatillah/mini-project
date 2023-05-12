package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateShipping(shipping *model.Shipping) error {
	if err := config.DB.Create(shipping).Error; err != nil {
		return err
	}
	return nil
}

func GetShipping(id uint) (shipping model.Shipping, err error) {
	if err := config.DB.Preload("User").First(&shipping, id).Error; err != nil {
		return model.Shipping{}, err
	}
	return shipping, nil
}
