package payload

type OrderRequest struct {
	UserID    uint   `json:"user_id" validate:"required"`
	ProductID uint   `json:"product_id" validate:"required"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}
