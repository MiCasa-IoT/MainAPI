package models

type Ping struct {
	Message      string     `json:"message" bson:"message" `
}
