package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type OrderUsecase interface {
	CreateOrder(order *model.Order) error
	GetOrder(id uint) (order model.Order, err error)
	GetListOrders() (orders []model.Order, err error)
	UpdateOrder(order *model.Order) (err error)
	DeleteOrder(id uint) (err error)
}

func CreateOrder(userId, productId uint, name, address string) (*model.Order, error) {
	_, err := database.GetProduct(productId)
	if err != nil {
		return nil, err
	}

	orders, err := database.GetOrders()
	if err != nil {
		return nil, err
	}
	for _, v := range orders {
		if v.Name == name {
			return nil, errors.New("order already exist")
		}
	}
	order := &model.Order{
		UserID:    userId,
		ProductID: productId,
		Name:      name,
		Address:   address,
	}

	err = database.CreateOrder(order)
	if err != nil {
		fmt.Println("error creating product")
		return nil, err
	}
	return order, nil
}

func GetOrder(id uint) (order model.Order, err error) {
	order, err = database.GetOrder(id)
	if err != nil {
		fmt.Println("GetOrder: Error getting order from database")
		return
	}
	return
}

func GetListOrders() (orders []model.Order, err error) {
	orders, err = database.GetOrders()
	if err != nil {
		fmt.Println("GetListOrders: Error getting orders from database")
		return
	}
	return
}

func DeleteOrder(id uint) (err error) {
	order := model.Order{}
	order.ID = id
	err = database.DeleteOrder(&order)
	if err != nil {
		fmt.Println("DeleteOrder : error deleting order, err: ", err)
		return
	}

	return
}
