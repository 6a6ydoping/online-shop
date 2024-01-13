package service

import (
	"github.com/6a6ydoping/online-shop/internal/config"
	"github.com/6a6ydoping/online-shop/internal/repository"
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
