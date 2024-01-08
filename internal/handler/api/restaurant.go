package api

type CreateRestaurantRequest struct {
	Name   string `json:"name"`
	City   int64  `json:"city"`
	Street int64  `json:"street"`
}

type UpdateRestaurantRequest struct {
	ID
	Name   string `json:"name" binding:"required,min=5,max=100"`
	City   string `json:"city" binding:"required,min=5,max=100"`
	Street string `json:"street" binding:"required,min=5,max=100"`
}
