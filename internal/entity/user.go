package entity

type User struct {
	ID       int64  `json:"id" db:"user_id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
