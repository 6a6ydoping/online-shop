package repository

import (
	"github.com/6a6ydoping/online-shop/internal/entity"
	"golang.org/x/net/context"
)

type UserRepository interface {
	//GetUser(ctx context.Context) error

	GetAllUsers(ctx context.Context) (*[]entity.User, error)
}
