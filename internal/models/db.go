package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgClient struct {
	DB *mongo.Database
	Client *mongo.Client
	Context context.Context
}