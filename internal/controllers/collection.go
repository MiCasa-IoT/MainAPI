package controllers

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
	"time"
)

func ReadAllDocumentHandler(ctx *gin.Context) {
	client, err := db.Connect()
	logging.PrintEror(err)

	cur, err := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV")).
		Find(context.Background(), bson.M{})
	logging.PrintEror(err)
	defer cur.Close(client.Context)

	var documents []bson.M
	err = cur.All(context.Background(), &documents)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"results": documents,
	})
}

func CreateHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintEror(err)

	params.CreatedAt = time.Now().Format(time.RFC3339)

	client, err := db.Connect()
	logging.PrintEror(err)
	collection := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV"))

	insertResult, err := collection.InsertOne(context.TODO(), params)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"result": insertResult.InsertedID,
	})
}

func UpdateHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func ReadHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintEror(err)

	findResult, err := db.FindById(params.UserID)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": findResult.UserID,
		"name": findResult.Name,
		"email": findResult.Email,
		"created_at": findResult.CreatedAt,
	})
}

func DeleteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
