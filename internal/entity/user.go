package entity

import "time"

type User struct {
	ID           int64     `json:"id" db:"user_id"`
	Username     string    `json:"username" db:"username" binding:"required"`
	Email        string    `json:"email" db:"email" binding:"required"`
	Password     string    `json:"password" db:"password" binding:"required"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	Gender       string    `json:"gender" db:"gender"`
	ProfileImage string    `json:"profileImage" db:"profile_image"`
}
