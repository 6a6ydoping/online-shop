package service

import (
	"context"
	"github.com/6a6ydoping/online-shop/internal/entity"
)

type UserService interface {
	// Auth

	GetAllUsers(context.Context) (*[]entity.User, error)
	CheckPasswordHash(password, hash string) bool
}
