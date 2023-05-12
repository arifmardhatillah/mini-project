package usecase

import (
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
	"time"
)

type PaymentUsecase interface {
	CreatePayment(payment *model.Payment) error
	GetPayment(id uint) (payment model.Payment, err error)
}

func CreatePayment(payment *model.Payment) error {
	payment.PaymentDate = time.Now()
	err := database.CreatePayment(payment)
	if err != nil {
		return err
	}

	return nil
}

func GetPayment(id uint) (payment model.Payment, err error) {
	payment, err = database.GetPayment(id)
	if err != nil {
		fmt.Println("GetPayment: Error getting payment from database")
		return
	}
	return
}
