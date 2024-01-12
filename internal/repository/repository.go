package repository

import "context"

type Repository interface {
	GetUser(ctx context.Context) error
}
