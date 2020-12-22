package models

type Message struct {
	Title string `json:"title" bson:"title"`
	Body string `json:"body" bson:"body"`
	Tag string `json:"tag" bson:"tag"`
	Tokens []string `json:"tokens" bson:"tokens"`
}
