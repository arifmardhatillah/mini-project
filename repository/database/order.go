package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateOrder(order *model.Order) error {
	if err := config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrders() (orders []model.Order, err error) {
	if err := config.DB.Preload("Product").First(&orders).Error; err != nil {
		return []model.Order{}, err
	}
	return orders, nil
}

func GetOrder(id uint) (order model.Order, err error) {
	if err := config.DB.Preload("Product").First(&order, id).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func UpdateOrder(order *model.Order) error {
	if err := config.DB.Updates(order).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOrder(order *model.Order) error {
	if err := config.DB.Delete(order).Error; err != nil {
		return err
	}
	return nil
}
