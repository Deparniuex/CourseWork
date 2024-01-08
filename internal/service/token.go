package service

import (
	"context"
	"ripProject/internal/entity"
	"ripProject/pkg/util"
)

func (m *Manager) Login(ctx context.Context, credentials string, password string) (*entity.Token, error) {
	user, err := m.Repository.GetUserByCredentials(ctx, credentials)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(password, *user.Password.Hash)
	if err != nil {
		return nil, err
	}

	token, err := util.GenerateToken()
	if err != nil {
		return nil, err
	}

	tokenEntity := &entity.Token{
		Plaintext: token.Plaintext,
		Hash:      token.Hash,
		UserID:    user.ID,
	}

	err = m.Repository.CreateToken(ctx, tokenEntity)
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
}
