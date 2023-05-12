package payload

type TopUpRequest struct {
	Saldo uint `json:"saldo" validate:"required,min=50000000"`
}
