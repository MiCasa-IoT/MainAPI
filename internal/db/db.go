package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func Connect() {
	connectionStr := os.Getenv("MONOGODB_CONNECTION_STR")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI(connectionStr)

	c, err := mongo.NewClient(clientOptions)
	if err = c.Connect(ctx); err != nil {
		fmt.Println("Could not connect to MongoDB:", err)
	}

	if err := c.Ping(ctx, nil); err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
	} else {
		fmt.Println("Connection Successful")
	}
}
