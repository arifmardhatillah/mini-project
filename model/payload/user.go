package payload

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Address  string `json:"address" validate:"required"`
	Role     string `json:"role"`
}

type CreateUserResponse struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type LoginUserResponse struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}
