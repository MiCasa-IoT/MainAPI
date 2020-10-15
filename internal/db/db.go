package db

import (
	"MiCasa-API/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
	}

	fmt.Println("Connection Successful")
	db := c.Database(os.Getenv("MONGODB_DB_DEV"))
	return &models.MgClient{
		DB:      db,
		Client:  c,
		Context: ctx,
	}, nil
}

func FindAll() ([]bson.M, error){
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	cur, err := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV")).
		Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(client.Context)

	var documents []bson.M
	err = cur.All(context.Background(), &documents)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func FindByID(id string) (models.User, error) {
	client, err := Connect()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV")).FindOne(context.Background(),
		bson.M{"user_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateByID(params models.User) (*mongo.UpdateResult, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	filter := bson.M{"user_id": params.UserID}
	update := bson.M{"$set": bson.M{"email": params.Email, "name": params.Name}}

	collection := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV"))

	updateResult, err := collection.UpdateOne(
		context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func InsertRecord(params models.User) (*mongo.InsertOneResult, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	collection := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV"))

	insertResult, err := collection.InsertOne(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func DeleteByID(id string) (*mongo.DeleteResult, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": id}
	deleteResult, err := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV")).DeleteOne(
			context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
