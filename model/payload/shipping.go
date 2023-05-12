package payload

type ShippingRequest struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Address string `json:"address" validate:"required"`
}
