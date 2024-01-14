package impl

import (
	"github.com/6a6ydoping/online-shop/internal/config"
	"github.com/6a6ydoping/online-shop/internal/repository"
	"github.com/6a6ydoping/online-shop/internal/service"
)

type UserManager struct {
	Repository repository.Repository
	Config     *config.Config
}

func (m *UserManager) CheckPasswordHash(password, hash string) bool {
	return false
}

func New(repository repository.Repository, config *config.Config) service.UserService {
	return &UserManager{
		Repository: repository,
		Config:     config,
	}
}
