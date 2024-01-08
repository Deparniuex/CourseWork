package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateOrderHistory(ctx context.Context, orderHistory *entity.OrderHistory) error {
	return m.Repository.CreateOrderHistory(ctx, orderHistory)
}

func (m *Manager) DeleteOrderHistory(ctx context.Context, orderHistoryID int64) error {
	return m.Repository.DeleteOrderHistory(ctx, orderHistoryID)
}

func (m *Manager) UpdateOrderHistory(ctx context.Context, orderHistory *entity.OrderHistory) error {
	return m.Repository.UpdateOrderHistory(ctx, orderHistory)
}
