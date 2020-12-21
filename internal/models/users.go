package models

type User struct {
	UserID    string `json:"user_id" bson:"user_id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
	Email     string `json:"email" bson:"email"`
	UUID      string `json:"uuid" bson:"uuid"`
}