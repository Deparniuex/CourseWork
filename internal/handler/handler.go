package handler

import (
	"ripProject/internal/config"
	"ripProject/internal/service"
)

type Handler struct {
	Services service.Service
	Config   *config.Config
}

func New(services service.Service, config *config.Config) *Handler {
	return &Handler{
		Services: services,
		Config:   config,
	}
}
