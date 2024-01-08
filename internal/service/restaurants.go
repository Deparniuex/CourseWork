package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateRestaurant(ctx context.Context, restaurant *entity.Restaurant) error {
	return m.Repository.CreateRestaurant(ctx, restaurant)
}

func (m *Manager) DeleteRestaurant(ctx context.Context, restaurantID int64) error {
	return m.Repository.DeleteRestaurant(ctx, restaurantID)
}

func (m *Manager) UpdateRestaurant(ctx context.Context, restaurant *entity.Restaurant) error {
	return m.Repository.UpdateRestaurant(ctx, restaurant)
}
