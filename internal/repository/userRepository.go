package repository

import (
	"github.com/6a6ydoping/online-shop/internal/entity"
	"golang.org/x/net/context"
)

type UserRepository interface {
	//GetUser(ctx context.Context) error

	GetAllUsers(c context.Context) (*[]entity.User, error)
	CreateUser(c context.Context, user entity.User) error
}
