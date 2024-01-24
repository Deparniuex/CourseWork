package service

import (
	"context"
	"ripProject/internal/entity"
)

func (m *Manager) CreateCity(ctx context.Context, city *entity.City) error {
	return m.Repository.CreateCity(ctx, city)
}

func (m *Manager) DeleteCity(ctx context.Context, cityID int64) error {
	return m.Repository.DeleteCity(ctx, cityID)
}

func (m *Manager) UpdateCity(ctx context.Context, city *entity.City) error {
	return m.Repository.UpdateCity(ctx, city)
}

func (m *Manager) GetCities(ctx context.Context) ([]*entity.City, error) {
	return m.Repository.GetCities(ctx)
}
