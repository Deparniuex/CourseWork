package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateMenu(ctx context.Context, menu *entity.Menu) error {
	return m.Repository.CreateMenu(ctx, menu)
}

func (m *Manager) DeleteMenu(ctx context.Context, menuID int64) error {
	return m.Repository.DeleteDish(ctx, menuID)
}

func (m *Manager) UpdateMenu(ctx context.Context, menu *entity.Menu) error {
	return m.Repository.UpdateMenu(ctx, menu)
}
