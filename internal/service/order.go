package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateOrder(ctx context.Context, order *entity.Order) error {
	return m.Repository.CreateOrder(ctx, order)
}

func (m *Manager) DeleteOrder(ctx context.Context, orderID int64) error {
	return m.Repository.DeleteOrder(ctx, orderID)
}

func (m *Manager) UpdateOrder(ctx context.Context, order *entity.Order) error {
	return m.Repository.UpdateOrder(ctx, order)
}
