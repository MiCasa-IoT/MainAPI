package controllers

import (
	"MiCasa-API/internal/db"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
)

func FindAllDocumentHandler(ctx *gin.Context) {
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
