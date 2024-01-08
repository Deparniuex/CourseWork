package service

import (
	"ripProject/internal/config"
	"ripProject/internal/repository"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
}

func New(repository repository.Repository, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Config:     config,
	}
}
