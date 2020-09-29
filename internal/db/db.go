package db

import (
	"MiCasa-API/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func Connect() (*models.MgClient, error) {
	connectionStr := os.Getenv("MONGODB_CONNECTION_STR_DEV")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI(connectionStr)

	c, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	if err = c.Connect(ctx); err != nil {
		return nil, err
	}

	if err := c.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	} else {
		fmt.Println("Connection Successful")
	}
	db := c.Database(os.Getenv("MONGODB_DB_DEV"))
	return &models.MgClient{
		DB:      db,
		Client:  c,
		Context: ctx,
	}, nil
}
