package payload

type PaymentRequest struct {
	OrderID uint `json:"order_id" validate:"required"`
}
