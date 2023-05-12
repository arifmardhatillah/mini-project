package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreatePayment(payment *model.Payment) error {
	if err := config.DB.Create(payment).Error; err != nil {
		return err
	}
	return nil
}

func GetPayment(id uint) (payment model.Payment, err error) {
	if err := config.DB.Preload("Order").First(&payment, id).Error; err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}
