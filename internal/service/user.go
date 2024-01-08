package service

import (
	"context"
	"ripProject/internal/entity"
	"ripProject/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, user *entity.User) error {
	hash, err := util.HashPassword(*user.Password.Plaintext)
	if err != nil {
		return err
	}

	user.Password.Hash = &hash
	return m.Repository.CreateUser(ctx, user)
}

func (m *Manager) GetUserByToken(ctx context.Context, token string) (*entity.User, error) {
	return m.Repository.GetUserByToken(ctx, token)
}

func (m *Manager) GetUserByCredentials(ctx context.Context, credentials string) (*entity.User, error) {
	return m.Repository.GetUserByCredentials(ctx, credentials)
}

func (m *Manager) DeleteUser(ctx context.Context, userID int64) error {
	return m.Repository.DeleteUser(ctx, userID)
}

func (m *Manager) UpdateUser(ctx context.Context, user *entity.User) error {
	return m.Repository.UpdateUser(ctx, user)
}
