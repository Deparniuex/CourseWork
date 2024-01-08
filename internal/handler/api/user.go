package api

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required,min=5,max=50" example:"username"`
	Email     string `json:"email" binding:"required,email" example:"example@gmail.com"`
	Password  string `json:"password" binding:"required,min=6,max=32" example:"password"`
	FirstName string `json:"first_name" binding:"required,min=2,max=50" example:"John"`
	LastName  string `json:"last_name" binding:"required,min=2,max=50" example:"Snow"`
	Phone     string `json:"phone" binding:"required"`
}

type UpdateUserRequest struct {
	Password  *string `json:"password" binding:"omitempty,min=6,max=32" example:"password"`
	FirstName *string `json:"first_name" binding:"omitempty,min=2,max=50" example:"John"`
	LastName  *string `json:"last_name" binding:"omitempty,min=2,max=50" example:"Snow"`
}

type LoginRequest struct {
	Credentials string `json:"credentials" binding:"required" example:"username"`
	Password    string `json:"password" binding:"required" example:"password"`
}

type GrantRoleToUser struct {
	Role string `json:"role" binding:"required"  example:"ADMIN"`
}
