package models

type Connection struct {
	UUID      string     `json:"uuid" bson:"uuid" `
	EdgeID	  int `json:"edge_id" bson:"edge_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}
