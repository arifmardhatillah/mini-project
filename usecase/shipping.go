package usecase

import (
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type ShippingUsecase interface {
	CreateShipping(shipping *model.Shipping) error
	GetShipping(id uint) (shipping model.Shipping, err error)
}

func CreateShipping(userId uint, address string) (*model.Shipping, error) {
	_, err := database.GetUser(userId)
	if err != nil {
		return nil, err
	}

	shipping := &model.Shipping{
		UserID:  userId,
		Address: address,
	}

	err = database.CreateShipping(shipping)
	if err != nil {
		fmt.Println("error creating shipping")
		return nil, err
	}
	return shipping, nil
}

func GetShipping(id uint) (shipping model.Shipping, err error) {
	shipping, err = database.GetShipping(id)
	if err != nil {
		fmt.Println("GetShipping: Error getting shipping from database")
		return
	}
	return
}
