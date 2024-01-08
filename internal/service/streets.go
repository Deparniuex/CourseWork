package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateStreet(ctx context.Context, street *entity.Street) error {
	return m.Repository.CreateStreet(ctx, street)
}

func (m *Manager) DeleteStreet(ctx context.Context, streetID int64) error {
	return m.Repository.DeleteStreet(ctx, streetID)
}

func (m *Manager) UpdateStreet(ctx context.Context, street *entity.Street) error {
	return m.Repository.UpdateStreet(ctx, street)
}
