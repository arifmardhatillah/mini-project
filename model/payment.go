package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID     uint      `json:"order_id"`
	TotalPrice  uint      `json:"total_price"`
	PaymentDate time.Time `json:"payment_date"`
	Order       Order     `json:"order"`
}
