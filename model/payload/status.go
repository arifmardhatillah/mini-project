package payload

type StatusRequest struct {
	Status string `json:"status" validate:"required,oneof=Delivery Complete"`
}
