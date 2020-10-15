package controllers

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"github.com/gin-gonic/gin"
	"time"

	_ "MiCasa-API/docs"
)

// ...
// @Summary 全てのドキュメントを取得し返す
// @Tags Document
// @Produce  json
// @Success 200 {object} responses.SuccessResponse{data=[]models.User}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} controller.HTTPError
// @Router /api/v1/db/document/readall [get]
func ReadAllDocumentHandler(ctx *gin.Context) {
	documents, err := db.FindAll()
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, documents)
}

func CreateHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)
	params.CreatedAt = time.Now().Format(time.RFC3339)

	createResult, err := db.InsertRecord(params)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, createResult.InsertedID)
}

func ReadHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)

	findResult, err := db.FindByID(params.UserID)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, findResult)
}

func UpdateHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)
	updateResult, err := db.UpdateByID(params)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, updateResult)
}

func DeleteHandler(ctx *gin.Context) {
	var params models.User
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)

	deleteResult, err := db.DeleteByID(params.UserID)
	logging.PrintErorWithGinContext(err, ctx)
	if deleteResult.DeletedCount > 0 {
		logging.StatusOK(err, ctx, deleteResult)
	} else {
		logging.StatusBadRequest(err, ctx, "User Not found")
	}
}
