package models

type User struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Email     string `json:"email"`
}
