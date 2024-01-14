package service

type UserService interface {
	// Auth

	CheckPasswordHash(password, hash string) bool
}
