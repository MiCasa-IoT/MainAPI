package db

import (
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/array"
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
	connectionStr := os.Getenv("MONGODB_CONNECTION_STR")

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
	db := c.Database(os.Getenv("MONGODB_DB"))
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
		os.Getenv("MONGODB_COLLECTION")).
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

func FilterByEdgeID(edgeId int) ([]string, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	var result []string
	ctx := context.Background()
	cur, err := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION")).Find(ctx,
		bson.M{"edge_id": edgeId})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var c models.Connection
		err := cur.Decode(&c)
		if err != nil {
			return nil, err
		}
		if array.NotStrContains(result, c.UUID) {
			result = append(result, c.UUID)
		}
	}

	return result, nil
}

func GetLatestUUIDToFilterByEdgeID(uuid []string, edgeId int) ([]string, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	var result []string
	ctx := context.Background()
	collection := client.DB.Collection(os.Getenv("MONGODB_COLLECTION"))
	option := options.FindOne().SetSort(bson.M{"created_at": -1})

	for _, u := range uuid {
		r := collection.FindOne(ctx, bson.M{"uuid": u}, option)

		var c models.Connection
		err := r.Decode(&c)
		if err != nil {
			return nil, err
		}
		print(c.UUID, c.EdgeID, c.CreatedAt)
		if c.EdgeID == edgeId {
			result = append(result, c.UUID)
		}
	}

	return result, nil
}

func UpdateByID(params models.Connection) (*mongo.UpdateResult, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	filter := bson.M{"uuid": params.UUID}
	update := bson.M{"$set": bson.M{"created_at": params.CreatedAt, "deleted_at": params.DeletedAt}}

	collection := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION"))

	updateResult, err := collection.UpdateOne(
		context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func InsertRecord(params models.Connection) (*mongo.InsertOneResult, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	collection := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION"))

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

	filter := bson.M{"uuid": id}
	deleteResult, err := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION")).DeleteOne(
			context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
