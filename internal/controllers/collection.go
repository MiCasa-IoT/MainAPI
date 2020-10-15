package controllers

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"time"
)

func ReadAllDocumentHandler(ctx *gin.Context) {
	client, err := db.Connect()
	if err != nil {
		log.Println(err)
	}
	cur, err := client.DB.Collection(
		os.Getenv("MONGODB_COLLECTION_DEV")).
		Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(client.Context)

	var documents []bson.M
	if err = cur.All(context.Background(), &documents); err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": documents,
	})
}

func CreateHandler(ctx *gin.Context) {
	var params models.User
	if err := ctx.BindJSON(&params); err != nil {
		log.Println(err)
	}
	params.CreatedAt = time.Now().Format(time.RFC3339)

	ctx.JSON(http.StatusOK, gin.H{
		"id": params.UserID,
		"name": params.Name,
		"email": params.Email,
		"created_at": params.CreatedAt,
	})
}

func UpdateHandler(ctx *gin.Context) {
	var params models.User
	if err := ctx.BindJSON(&params); err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func ReadHandler(ctx *gin.Context) {
	var params models.User
	if err := ctx.BindJSON(&params); err != nil {
		log.Println(err)
	}

	findResult, err := db.FindById(params.UserID)
	if err != nil {
		log.Println(err)
	}
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
