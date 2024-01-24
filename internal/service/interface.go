package service

import (
	"context"
	"ripProject/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID int64) error
	GetUserByToken(ctx context.Context, token string) (*entity.User, error)
	GetUserByCredentials(ctx context.Context, credentials string) (*entity.User, error)

	Login(ctx context.Context, credentials string, password string) (*entity.Token, error)

	CreateStreet(ctx context.Context, street *entity.Street) error
	UpdateStreet(ctx context.Context, street *entity.Street) error
	DeleteStreet(ctx context.Context, streetID int64) error

	CreateRestaurant(ctx context.Context, restaurant *entity.Restaurant) error
	UpdateRestaurant(ctx context.Context, restaurant *entity.Restaurant) error
	DeleteRestaurant(ctx context.Context, restaurantID int64) error

	CreateOrderHistory(ctx context.Context, orderHistory *entity.OrderHistory) error
	UpdateOrderHistory(ctx context.Context, orderHistory *entity.OrderHistory) error
	DeleteOrderHistory(ctx context.Context, orderHistoryID int64) error

	CreateOrder(ctx context.Context, order *entity.Order) error
	UpdateOrder(ctx context.Context, order *entity.Order) error
	DeleteOrder(ctx context.Context, orderID int64) error

	CreateMenu(ctx context.Context, menu *entity.Menu) error
	UpdateMenu(ctx context.Context, menu *entity.Menu) error
	DeleteMenu(ctx context.Context, menuID int64) error

	CreateDish(ctx context.Context, dish *entity.Dish) error
	UpdateDish(ctx context.Context, dish *entity.Dish) error
	DeleteDish(ctx context.Context, dishID int64) error

	CreateCity(ctx context.Context, menu *entity.City) error
	UpdateCity(ctx context.Context, city *entity.City) error
	DeleteCity(ctx context.Context, cityID int64) error
	GetCities(ctx context.Context) ([]*entity.City, error)
}
