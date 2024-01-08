package api

import (
	"ripProject/internal/entity"
)

type LoginResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Body    *entity.Token `json:"body"`
}

type GetUserByUsernameResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Body    *entity.User `json:"body"`
}

type GetRestaurantByIDResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Body    *entity.Restaurant `json:"body"`
}

type GetRestaurantsResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Body    []*entity.Restaurant `json:"body"`
}

type GetDishesByIDResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Body    []*entity.Dish `json:"body"`
}
