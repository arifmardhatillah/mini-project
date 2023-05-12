package payload

type UpdateStockRequest struct {
	Stock uint `json:"stock" validate:"required"`
}
