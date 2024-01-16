package impl

import (
	"context"
	"github.com/6a6ydoping/online-shop/internal/config"
	"github.com/6a6ydoping/online-shop/internal/entity"
	"github.com/6a6ydoping/online-shop/internal/repository"
	"github.com/6a6ydoping/online-shop/internal/service"
)

type UserManager struct {
	Repository repository.UserRepository
	Config     *config.Config
}

func (m *UserManager) CheckPasswordHash(password, hash string) bool {
	return false
}

func (m *UserManager) GetAllUsers(c context.Context) (*[]entity.User, error) {
	var users *[]entity.User
	users, err := m.Repository.GetAllUsers(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserManager) CreateUser(c context.Context, user entity.User) error {
	// TODO: HERE CHECK EMAIL, PASSWORD, IF USER ALREADY EXISTS, ETC
	err := m.Repository.CreateUser(c, user)
	if err != nil {
		return err
	}

	return nil
}

func New(repository repository.UserRepository, config *config.Config) service.UserService {
	return &UserManager{
		Repository: repository,
		Config:     config,
	}
}
