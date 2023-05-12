package payload

type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
}
