package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateDish(ctx context.Context, dish *entity.Dish) error {
	return m.Repository.CreateDish(ctx, dish)
}

func (m *Manager) DeleteDish(ctx context.Context, dishID int64) error {
	return m.Repository.DeleteDish(ctx, dishID)
}

func (m *Manager) UpdateDish(ctx context.Context, dish *entity.Dish) error {
	return m.Repository.UpdateDish(ctx, dish)
}
