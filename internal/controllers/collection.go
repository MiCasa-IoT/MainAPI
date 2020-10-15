package controllers

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ReadAllDocumentHandler(ctx *gin.Context) {
	documents, err := db.FindAll()
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

	result, err := db.InsertRecord(params)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"result": result.InsertedID,
	})
}

func ReadHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintEror(err)

	findResult, err := db.FindByID(params.UserID)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"user_id":    findResult.UserID,
		"name":       findResult.Name,
		"email":      findResult.Email,
		"created_at": findResult.CreatedAt,
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

func DeleteHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintEror(err)

	deleteResult, err := db.DeleteByID(params.UserID)
	logging.PrintEror(err)

	ctx.JSON(http.StatusOK, gin.H{
		"result": deleteResult,
	})
}
